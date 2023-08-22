// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin/contracts/access/Ownable.sol";

contract TitanVPS is Ownable {
    mapping(string => bool) public orders;

    event OrderEvent (
        string indexed orderID,
        address indexed buyerAddr,
        uint256 price,
        uint256 expiration,
        uint256 timestamp
    );

    /**
     * @dev Create a new order with the provided details.
     *
     * @param _orderID The unique identifier for the order.
     * @param _buyerAddr The address of the buyer placing the order.
     * @param _price The price associated with the order.
     * @param _expiration The expiration timestamp of the order.
     * @param _timestamp The timestamp when the order is created.
     */
    function newOrder(
        string memory _orderID,
        address _buyerAddr,
        uint256 _price,
        uint256 _expiration,
        uint256 _timestamp
    ) public onlyOwner {
        // Mark the order as active
        orders[_orderID] = true;

        // Emit an event to notify about the new order
        emit OrderEvent(_orderID, _buyerAddr, _price, _expiration, _timestamp);
    }
}
