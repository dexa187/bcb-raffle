const Raffle = artifacts.require("Raffle");
const BCB = artifacts.require("BCB")

module.exports = async function(_deployer, network) {
  if (network == "xdai") {
    bcb = await BCB.at("0x4B78a47532D9e966574D30189B3dE734A232A78a")
  }else {
    bcb = await _deployer.deploy(BCB, 'Buttercup Bucks Test', 'TBCB');
  }
  raffle = await _deployer.deploy(Raffle);
  raffle.setTokenContract(BCB.address)
};
