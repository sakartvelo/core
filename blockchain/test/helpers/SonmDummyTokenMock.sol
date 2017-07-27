pragma solidity ^0.4.11;


import '../../contracts/SDT.sol';


// mock class using StandardToken
contract SonmDummyTokenMock is SonmDummyToken {

  function SonmDummyTokenMock(address initialAccount, uint initialBalance) {
    balances[initialAccount] = initialBalance;
    totalSupply = initialBalance;
  }

}
