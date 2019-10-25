package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	raffle "github.com/dexa187/bcb-raffle/truffle/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var key = os.Getenv("PK")

type RaffleClient struct {
	client         ethclient.Client
	raffleContract raffle.Raffle
	bcbContract    raffle.BCB
	walletPtr      string
	contractPtr    string
}

func main() {

	contractPtr := flag.String("bcbContract", "0x4B78a47532D9e966574D30189B3dE734A232A78a", "Address of BCB Contract")
	walletPtr := flag.String("raffleContract", "0xCE23697A91Bd50aB4d3Fa49DfeECD3f5Ee44b4E7", "Address of the Wallet to Watch")
	wsURLPtr := flag.String("wsURL", "wss://dai-trace-ws.blockscout.com/ws", "Websocket URL of blockchain node")

	flag.Parse()

	client, err := ethclient.Dial(*wsURLPtr)
	if err != nil {
		log.Fatal(err)
	}

	raffleContract, err := raffle.NewRaffle(common.HexToAddress(*walletPtr), client)
	if err != nil {
		log.Fatal("Failed to Instantiate Raffle Contract Check the wallet address")
	}
	bcbContract, err := raffle.NewBCB(common.HexToAddress(*contractPtr), client)
	if err != nil {
		log.Fatal("Failed to Instantiate BCB Contract Check the contract address")
	}

	raffleClient := RaffleClient{*client, *raffleContract, *bcbContract, *walletPtr, *contractPtr}

	name, _ := raffleContract.Name(nil)

	fmt.Printf("Connected to %s\n", name)

	watchIncomingTokens(raffleClient)
	watchNewTickets(raffleClient)
	watchBlockHeaders(*client)
}

func watchIncomingTokens(raffleClient RaffleClient) {
	// Watch for incoming BCB to raffle contract
	bcbEvents := make(chan *raffle.BCBTransfer)
	options := &bind.WatchOpts{nil, nil}

	_, err := raffleClient.bcbContract.WatchTransfer(options, bcbEvents, nil, []common.Address{common.HexToAddress(raffleClient.walletPtr)})
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for {
			event := <-bcbEvents
			if event.Value.Uint64() >= 1000000000000000000 {
				// TODO Repeat buying ticket if the user sends more than 1 BCB
				//for i := 0; i < 1000000000000000000%event.Value.Uint64; i++ {
				fmt.Printf("Buying Ticket For %s\n", event.From.Hex())
				auth, err := getTransactionOpts(raffleClient.client)
				if err != nil {
					log.Fatal("Failed to Get Auth", err)
				}
				for {
					_, err = raffleClient.raffleContract.CreateTicket(auth, event.From)
					if err != nil {
						fmt.Printf("Failed to purchase ticket Retrying %s\n", err)
						auth, _ = getTransactionOpts(raffleClient.client)
					} else {
						break
					}
				}
				//}
			}
		}
	}()
}

func watchNewTickets(raffleClient RaffleClient) {
	// Watch for new ticket Events
	ticketEvents := make(chan *raffle.RaffleTicketPurchased)
	options := &bind.WatchOpts{nil, nil}

	_, err := raffleClient.raffleContract.WatchTicketPurchased(options, ticketEvents)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			event := <-ticketEvents

			fmt.Printf("New Ticket: %s For %s\n", event.Id, event.Player.Hex())
			auth, err := getTransactionOpts(raffleClient.client)
			if err != nil {
				log.Fatal("Failed to Get Auth", err)
			}
			for {
				_, err = raffleClient.bcbContract.TransferWithData(auth, event.Player, big.NewInt(10000000000000000), []byte(fmt.Sprintf("You Got Ticket %s", event.Id)))
				if err != nil {
					fmt.Printf("Failed To Send Message %s\n", err)
					auth, _ = getTransactionOpts(raffleClient.client)
				} else {
					break
				}
			}
		}
	}()
}

func watchBlockHeaders(client ethclient.Client) {
	//This keeps the websocket connection alive
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Printf("New Block Hash %s \n", header.Hash().Hex())
		}
	}
}

func getTransactionOpts(client ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return &bind.TransactOpts{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return &bind.TransactOpts{}, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return &bind.TransactOpts{}, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return &bind.TransactOpts{}, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
