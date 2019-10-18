pragma solidity ^0.5.0;

import 'openzeppelin-solidity/contracts/token/ERC20/ERC20.sol';
import 'openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract BCB is ERC20, Ownable {
    string public name;
    string public symbol;
    uint8 public decimals = 18;

    constructor(string memory _name, string memory _symbol) public {
        name = _name;
        symbol = _symbol;
    }

    function mint(address to, uint256 amount) public onlyOwner returns (bool) {
        _mint(to, amount);
        return true;
    }

    function burn(address from, uint256 value) public onlyOwner returns (bool) {
        _burn(from, value);
        return true;
    }

    function transferWithData(address to, uint256 value, bytes memory data) public returns (bool) {
        emit TransferWithData(msg.sender, to, value, data);
        return transfer(to, value);
    }

    event TransferWithData(address indexed from, address indexed to, uint256 value, bytes data);
}