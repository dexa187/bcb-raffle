const Raffle = artifacts.require("Raffle");
const BCB = artifacts.require("BCB")

module.exports = async function(_deployer) {
  bcb = await _deployer.deploy(BCB, 'Buttercup Bucks', 'BCB');
  raffle = await _deployer.deploy(Raffle);
  raffle.setTokenContract(BCB.address)
};
