pragma solidity ^0.5.0;

import './BCB.sol';
import './node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol';
import './node_modules/openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract Raffle is Ownable {
  using SafeMath for uint256;

  string public name;
  uint public currentRound = 0;
  uint public currentPotLimit = 5;
  uint public potMax = 160;
  address public token;
  address[] public charities;
  mapping(uint => Round) public rounds;

  struct Round {
    address winner;
    uint ticketCount;
    uint pot;
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

  event NewRound (
    uint round,
    uint potLimit
  );

  // Lets you send ether to contract
  function() external payable {}

  function withdraw(uint256 value) public onlyOwner {
    msg.sender.transfer(value);
  }

  function createTicket(address _player) public onlyOwner returns (uint) {
    uint ticketCount = rounds[currentRound].ticketCount;
    rounds[currentRound].tickets[ticketCount].player = _player;
    rounds[currentRound].ticketCount ++;
    rounds[currentRound].pot ++;
    emit TicketPurchased(ticketCount, _player);
    if (rounds[currentRound].pot >= currentPotLimit){
      newDrawing();
    }
    return ticketCount;
  }

  function getTicket(uint _round, uint _ticket) public view returns (address _player) {
    return rounds[_round].tickets[_ticket].player;
  }

  function newDrawing() public onlyOwner {
    uint _winner = uint(blockhash(block.number-1))%rounds[currentRound].ticketCount;
    rounds[currentRound].winner = rounds[currentRound].tickets[_winner].player;

    BCB erc20token = BCB(token);
    uint balance = erc20token.balanceOf(address(this));
    erc20token.transferWithData(rounds[currentRound].tickets[_winner].player, balance/2, bytes("You Won the Raffle"));
    for (uint i = 0; i < charities.length; i++) {
      erc20token.transfer(charities[i], balance/2/charities.length);
    }

    emit DrawingCompleted(currentRound,rounds[currentRound].tickets[_winner].player,_winner,balance/2);
    if (currentPotLimit >= potMax ){
      currentPotLimit = 5;
    }else{
      currentPotLimit = currentPotLimit * 2;
    }
    currentRound++;
    emit NewRound(currentRound, currentPotLimit);
  }

  function setTokenContract(address _contract) public onlyOwner {
    token = _contract;
  }

  function setCharities(address[] memory _charities) public onlyOwner {
    charities = _charities;
  }

  function setCurrentPotLimit(uint _potLimit) public onlyOwner {
    currentPotLimit = _potLimit;
  }

  function setPotMax(uint _potMax) public onlyOwner {
    potMax = _potMax;
  }

  constructor() public {
    name = "Buttercup 50/50 Raffle";
    rounds[currentRound] = Round(address(0),0,5);
  }
}
