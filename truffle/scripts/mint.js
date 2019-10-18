const BCB = artifacts.require('BCB');

const MintAmount = '10000000000000000000000';

module.exports = function(callback) {
    async function mintBCB() {
        bcb = await BCB.deployed();
        accounts = await BCB.web3.eth.getAccounts();
        receipt = await bcb.mint(accounts[0], MintAmount);
        console.log(
            `${MintAmount} BCB Minted to ${accounts[0]} at a cost of ${BCB.web3.utils.fromWei(BCB.web3.utils.toWei(
                receipt.receipt.gasUsed.toString(), 'Gwei'),'ether'
            )} ETH`
        );
        totalSupply = await bcb.totalSupply();
        console.log(`${totalSupply} BCB in circulation`);
        callback();
    }
    mintBCB();
};