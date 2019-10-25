const Raffle = artifacts.require("Raffle");
const BCB = artifacts.require("BCB");

const MintAmount = "10000000000000000000000";
const cost = BCB.web3.utils.toWei("1","ether")

module.exports = function(callback) {
  async function mintBCB() {
    bcb = await BCB.deployed();
    accounts = await BCB.web3.eth.getAccounts();
    for (account of accounts) {
    receipt = await bcb.mint(account, MintAmount);
    console.log(
      `${MintAmount} BCB Minted to ${
        account
      } at a cost of ${BCB.web3.utils.fromWei(
        BCB.web3.utils.toWei(receipt.receipt.gasUsed.toString(), "Gwei"),
        "ether"
      )} ETH`
    );
      }
    totalSupply = await bcb.totalSupply();

    console.log(`${totalSupply} BCB in circulation`);
  }

  async function startRaffle() {
    raffle = await Raffle.deployed();
    bcb = await BCB.deployed();
    accounts = await Raffle.web3.eth.getAccounts();
    await raffle.setCharities(['0xEc78144387eA2e2c0ABbdc7cb6115D0Aa2DF06BC','0xEc78144387eA2e2c0ABbdc7cb6115D0Aa2DF06BC'])
    for (i = 1; i < 20; i++) {
      for (account in accounts) {
        if (account != 0){
          await bcb.transferWithData(raffle.address, cost, [], {from: accounts[account]});
        }
        // ticket = await raffle.createTicket(accounts[account]);
        // for (log of ticket.logs){
        //   switch (log.event) {
        //     case "TicketPurchased":
        //       console.log(`Ticket Purchaed by ${log.args['1']}`)
        //       break;
        //     case "DrawingCompleted":
        //       console.log(`Drawing completed Winner ${log.args.winner} Ticket ${log.args.ticketNumber.toString()} Amount ${log.args.amount.toString()}`);
        //       break;
        //     case "NewRound":
        //       console.log(`New Round ${log.args.round.toString()} Limit ${log.args.potLimit.toString()}`);
        //       break;
        //   }
        // }
      }
    }
    callback();
  }
  mintBCB()
  startRaffle();
};
