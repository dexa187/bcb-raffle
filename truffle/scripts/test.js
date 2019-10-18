const Raffle = artifacts.require("Raffle");
const BCB = artifacts.require("BCB");

const MintAmount = "10000000000000000000000";

module.exports = function(callback) {
  async function mintBCB() {
    bcb = await BCB.deployed();
    accounts = await BCB.web3.eth.getAccounts();
    receipt = await bcb.mint(accounts[0], MintAmount);
    console.log(
      `${MintAmount} BCB Minted to ${
        accounts[0]
      } at a cost of ${BCB.web3.utils.fromWei(
        BCB.web3.utils.toWei(receipt.receipt.gasUsed.toString(), "Gwei"),
        "ether"
      )} ETH`
    );
    totalSupply = await bcb.totalSupply();

    console.log(`${totalSupply} BCB in circulation`);
  }

  async function startRaffle() {
    raffle = await Raffle.deployed();
    bcb = await BCB.deployed();
    accounts = await Raffle.web3.eth.getAccounts();
    await raffle.setCharities(['0xEc78144387eA2e2c0ABbdc7cb6115D0Aa2DF06BC','0xEc78144387eA2e2c0ABbdc7cb6115D0Aa2DF06BC'])
    console.log(accounts);
    for (i = 0; i < 20; i++) {
      console.log(`Round ${i}`);
      for (account in accounts) {
        ticket = await raffle.createTicket(accounts[account]);
      }
      //console.log(raffle.address)
      await bcb.transfer(raffle.address,100)
      receipt = await raffle.newDrawing();
      console.log(receipt.logs[0].args.winner);
      console.log(receipt.logs[0].args.ticketNumber.toString())
      console.log(receipt.logs[0].args.amount.toString())
    }
    callback();
  }
  mintBCB()
  startRaffle();
};
