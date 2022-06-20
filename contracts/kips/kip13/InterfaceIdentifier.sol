// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../../sc/bridge/interface/IERC20BridgeReceiver.sol";

interface InterfaceIdentifier is IERC20BridgeReceiver {
    /// @notice Query if a contract implements an interface
    /// @param interfaceID The interface identifier, as defined in KIP-13.
    /// @dev Interface identifier is defined in KIP-13. This function
    ///  uses less than 30,000 gas.
    /// @return `true` if the contract implements `interfaceID` and
    ///  `interfaceID` is not 0xffffffff, `false` otherwise.
    function supportsInterface(bytes4 interfaceID) external view returns (bool);
}
