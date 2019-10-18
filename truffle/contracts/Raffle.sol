pragma solidity ^0.5.0;

import 'openzeppelin-solidity/contracts/math/SafeMath.sol';
import 'openzeppelin-solidity/contracts/token/ERC20/ERC20.sol';
import 'openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract Raffle is Ownable {
  using SafeMath for uint256;

  string public name;
  uint public currentRound = 0;
  address public token;
  address[] public charities;
  mapping(uint => Round) public rounds;

  struct Round {
    address winner;
    uint ticketCount;
    mapping(uint => Ticket) tickets;
  }

  struct Ticket {
    address player;
  }

  event TicketPurchased (
    uint id,
    address player
  );

  event DrawingCompleted (
    uint round,
    address winner,
    uint ticketNumber,
    uint amount
  );

  // Lets you send ether to contract
  function() external payable {}

  function withdraw(uint256 value) public onlyOwner {
    msg.sender.transfer(value);
  }

  function createTicket(address _player) public onlyOwner {
    uint ticketCount = rounds[currentRound].ticketCount;
    rounds[currentRound].tickets[ticketCount].player = _player;
    rounds[currentRound].ticketCount ++;
    emit TicketPurchased(ticketCount, _player);
  }

  function getTicket(uint _round, uint _ticket) public view returns (address _player) {
    return rounds[_round].tickets[_ticket].player;
  }

  function convert(bytes32 b) public pure returns(uint) {
        return uint(b);
  }

  function newDrawing() public onlyOwner {
    uint _winner = uint(blockhash(block.number-1))%rounds[currentRound].ticketCount;
    rounds[currentRound].winner = rounds[currentRound].tickets[_winner].player;

    ERC20 erc20token = ERC20(token);
    uint balance = erc20token.balanceOf(address(this));
    erc20token.transfer(rounds[currentRound].tickets[_winner].player, balance/2);
    for (uint i = 0; i < charities.length; i++) {
      erc20token.transfer(charities[i], balance/2/charities.length);
    }

    emit DrawingCompleted(_winner,rounds[currentRound].tickets[_winner].player,_winner,balance/2);
    currentRound++;
  }

  function setTokenContract(address _contract) public onlyOwner {
    token = _contract;
  }

  function setCharities(address[] memory _charities) public onlyOwner {
    charities = _charities;
  }

  constructor() public {
    name = "Buttercup 50/50 Raffle";
    rounds[currentRound] = Round(address(0),0);
  }
}
