pragma solidity ^0.4.11;


import "./StandardToken.sol";


/**
 * @title SimpleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `StandardToken` functions.
 */
contract HealToken is StandardToken {

  string public name = "Heal Token";
  string public symbol = "Heal";
  uint256 public decimals = 18;
  uint256 public INITIAL_SUPPLY = 100000000 * 1 ether;
  mapping(address => bool) islab;

  /**
   * @dev Contructor that gives msg.sender all of existing tokens.
   */
  function HealToken() {
    totalSupply = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
  }

  function is_lab() constant returns (bool res){
    return islab[msg.sender];
  }
  
  function set_lab(bool labstatus)    {
	islab[msg.sender] = labstatus;
  }
  
  
}