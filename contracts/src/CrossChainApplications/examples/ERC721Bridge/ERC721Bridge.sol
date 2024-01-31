// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC721Bridge} from "./IERC721Bridge.sol";
import {BridgeNFT} from "./BridgeNFT.sol";
import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IWarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {IERC721, ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {IERC721Receiver} from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "forge-std/console.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Implementation of the {IERC721Bridge} interface.
 *
 * This implementation uses the {BridgeToken} contract to represent tokens on this chain, and uses
 * {ITeleporterMessenger} to send and receive messages to other chains.
 */
contract ERC721Bridge is
    IERC721Bridge,
    TeleporterOwnerUpgradeable,
    IERC721Receiver
{
    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;
    bytes32 public immutable currentBlockchainID;
    uint256 public constant CREATE_BRIDGE_TOKENS_REQUIRED_GAS = 2_000_000;
    uint256 public constant MINT_BRIDGE_TOKENS_REQUIRED_GAS = 200_000;
    uint256 public constant TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS = 300_000;

    // Mapping to keep track of bridged tokenIds
    // MAYBE UNNECESSARY???
    mapping(address nativeNFTContractAddress => mapping(uint256 tokenId => bool locked))
        public bridgedTokens;

    // Mapping to keep track of submitted create bridge token requests
    mapping(bytes32 destinationBlockchainID => mapping(address destinationBridgeAddress => mapping(address nftContract => bool submitted)))
        public submittedBridgeNFTCreations;

    // Set of bridge nfts created by this bridge instance.
    mapping(address bridgeNFT => bool bridgeTokenExists)
        public bridgedNFTContracts;

    // Tracks the wrapped bridge token contract address for each native token bridged to this bridge instance.
    // (nativeBlockchainID, nativeBridgeAddress, nativeTokenAddress) -> bridgeTokenAddress
    mapping(bytes32 nativeBlockchainID => mapping(address nativeBridgeAddress => mapping(address nativeTokenAddress => address bridgeNFTAddress)))
        public nativeToBridgedNFT;

    /**
     * @dev Initializes the Teleporter Messenger used for sending and receiving messages,
     * and initializes the current chain ID.
     */
    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {
        currentBlockchainID = IWarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes memory
    ) public virtual override returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /**
     * @dev See {IERC721Bridge-bridgeERC721}.
     *
     * This function is called by the bridge contract on the source chain to lock tokens in this contract
     * and mint bridge tokens on this chain.
     *
     * Requirements:
     *
     * - `destinationBlockchainID` cannot be the same as the current chain ID.
     * - `recipient` cannot be the zero address.
     * - `destinationBridgeAddress` cannot be the zero address.
     * - `nftContractAddress` must be a valid ERC721 contract.
     * - `tokenId` must be a valid token ID for the ERC721 contract.
     */
    function bridgeERC721(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nftContractAddress,
        address recipient,
        uint256 tokenId
    ) external nonReentrant {
        // Bridging tokens within a single chain is not allowed.
        require(
            destinationBlockchainID != currentBlockchainID,
            "ERC721Bridge: cannot bridge to same chain"
        );

        // Neither the recipient, nor the NFT contract, nor the destination bridge can be the zero address.
        require(
            nftContractAddress != address(0),
            "ERC721Bridge: zero NFT contract address"
        );
        require(
            recipient != address(0),
            "ERC721Bridge: zero recipient address"
        );
        require(
            destinationBridgeAddress != address(0),
            "ERC721Bridge: zero destination bridge address"
        );

        // If the token to be bridged is an existing bridged NFT of this bridge,
        // then handle it by burning the NFT on this chain, and sending a message
        // back to the native chain.
        // Otherwise, handle it by locking the NFT in this bridge instance,
        // and sending a message to the destination to mint the equivalent NFT on the destination chain.
        if (bridgedNFTContracts[nftContractAddress]) {
            return
                _processBridgedNftTransfer(
                    destinationBlockchainID,
                    destinationBridgeAddress,
                    nftContractAddress,
                    tokenId,
                    recipient
                );
        }

        // Check if requests to create a BridgeNFT contract on the destination chain have been submitted
        require(
            submittedBridgeNFTCreations[destinationBlockchainID][
                destinationBridgeAddress
            ][nftContractAddress],
            "ERC721Bridge: invalid bridge NFT contract"
        );

        // Check that the token ID is not already bridged
        // If the owner of the token is this contract, then the token is already bridged.
        address tokenOwner = IERC721(nftContractAddress).ownerOf(tokenId);
        require(
            tokenOwner != address(this),
            "ERC721Bridge: token already bridged"
        );

        // Check that the token ID is owned by the sender
        require(tokenOwner == msg.sender, "ERC721Bridge: invalid token ID");

        // Transfer and lock NFT into the bridge contract

        IERC721(nftContractAddress).safeTransferFrom(
            msg.sender,
            address(this),
            tokenId
        );

        // Mark token as locked
        // May be unnecessary, remove?
        bridgedTokens[nftContractAddress][tokenId] = true;

        _processNativeTokenTransfer(
            destinationBlockchainID,
            destinationBridgeAddress,
            nftContractAddress,
            recipient,
            tokenId
        );
    }

    /**
     * @dev See {IERC721Bridge-submitCreateBridgeToken}.
     *
     * We allow for `submitCreateBridgeToken` to be called multiple times with the same bridge and token
     * information because a previous message may have been dropped or otherwise selectively not delivered.
     * If the bridge token already exists on the destination, we are sending a message that will
     * simply have no effect on the destination.
     *
     * Emits a {SubmitCreateBridgeToken} event.
     */
    function submitCreateBridgeERC721(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        ERC721 nftContract
    ) external nonReentrant {
        require(
            destinationBridgeAddress != address(0),
            "ERC721Bridge: zero destination bridge address"
        );

        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();

        // Create the calldata to create the bridge token on the destination chain.
        bytes memory messageData = encodeCreateBridgeNFTData(
            address(nftContract),
            nftContract.name(),
            nftContract.symbol()
        );

        // Send Teleporter message.
        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: CREATE_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        submittedBridgeNFTCreations[destinationBlockchainID][
            destinationBridgeAddress
        ][address(nftContract)] = true;

        emit SubmitCreateBridgeNFT({
            destinationBlockchainID: destinationBlockchainID,
            destinationBridgeAddress: destinationBridgeAddress,
            nativeContractAddress: address(nftContract),
            teleporterMessageID: messageID
        });
    }

    function _processNativeTokenTransfer(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) private {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();

        // Send Teleporter message.
        bytes memory messageData = encodeMintBridgeNFTData(
            nativeContractAddress,
            recipient,
            tokenId
        );

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: MINT_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit BridgedNFT({
            tokenContractAddress: nativeContractAddress,
            destinationBlockchainID: destinationBlockchainID,
            teleporterMessageID: messageID,
            destinationBridgeAddress: destinationBridgeAddress,
            recipient: recipient,
            tokenId: tokenId
        });
    }

    /**
     * @dev Encodes the parameters for the Mint action to be decoded and executed on the destination.
     */
    function encodeMintBridgeNFTData(
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) public pure returns (bytes memory) {
        // ABI encode the Mint action and corresponding parameters for the mintBridgeTokens
        // call to to be decoded and executed on the destination.
        bytes memory paramsData = abi.encode(
            nativeContractAddress,
            recipient,
            tokenId
        );
        return abi.encode(BridgeAction.Mint, paramsData);
    }

    /**
     * @dev Encodes the parameters for the Create action to be decoded and executed on the destination.
     */
    function encodeCreateBridgeNFTData(
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol
    ) public pure returns (bytes memory) {
        // ABI encode the Create action and corresponding parameters for the createBridgeToken
        // call to to be decoded and executed on the destination.
        bytes memory paramsData = abi.encode(
            nativeContractAddress,
            nativeName,
            nativeSymbol
        );
        return abi.encode(BridgeAction.Create, paramsData);
    }

    /**
     * @dev See {TeleporterUpgradeable-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Decode the payload to recover the action and corresponding function parameters
        (BridgeAction action, bytes memory actionData) = abi.decode(
            message,
            (BridgeAction, bytes)
        );

        // Route to the appropriate function.
        if (action == BridgeAction.Create) {
            (
                address nativeContractAddress,
                string memory nativeName,
                string memory nativeSymbol
            ) = abi.decode(actionData, (address, string, string));
            _createBridgeToken({
                nativeBlockchainID: sourceBlockchainID,
                nativeBridgeAddress: originSenderAddress,
                nativeContractAddress: nativeContractAddress,
                nativeName: nativeName,
                nativeSymbol: nativeSymbol
            });
        } else if (action == BridgeAction.Mint) {
            (
                address nativeContractAddress,
                address recipient,
                uint256 tokenId
            ) = abi.decode(actionData, (address, address, uint256));
            _mintBridgeToken(
                sourceBlockchainID,
                originSenderAddress,
                nativeContractAddress,
                recipient,
                tokenId
            );
        } else if (action == BridgeAction.Transfer) {
            (
                bytes32 destinationBlockchainID,
                address destinationBridgeAddress,
                address nativeContractAddress,
                address recipient,
                uint256 tokenId
            ) = abi.decode(
                    actionData,
                    (bytes32, address, address, address, uint256)
                );
            _transferBridgeToken(
                destinationBlockchainID,
                destinationBridgeAddress,
                nativeContractAddress,
                recipient,
                tokenId
            );
        } else {
            revert("ERC721Bridge: invalid action");
        }
    }

    /**
     * @dev Teleporter message receiver for creating a new bridge token on this chain.
     *
     * Emits a {CreateBridgeToken} event.
     *
     * Note: This function is only called within `receiveTeleporterMessage`, which can only be
     * called by the Teleporter Messenger.
     */
    function _createBridgeToken(
        bytes32 nativeBlockchainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol
    ) private {
        // Check that the bridge token doesn't already exist.
        require(
            nativeToBridgedNFT[nativeBlockchainID][nativeBridgeAddress][
                nativeContractAddress
            ] == address(0),
            "ERC721Bridge: bridge token already exists"
        );

        address bridgeNFTAddress = address(
            new BridgeNFT({
                sourceBlockchainID: nativeBlockchainID,
                sourceBridge: nativeBridgeAddress,
                sourceAsset: nativeContractAddress,
                tokenName: nativeName,
                tokenSymbol: nativeSymbol
            })
        );

        bridgedNFTContracts[bridgeNFTAddress] = true;
        nativeToBridgedNFT[nativeBlockchainID][nativeBridgeAddress][
            nativeContractAddress
        ] = bridgeNFTAddress;

        emit CreateBridgeNFT(
            nativeBlockchainID,
            nativeBridgeAddress,
            nativeContractAddress,
            bridgeNFTAddress
        );
    }

    /**
     * @dev Teleporter message receiver for minting of an existing bridge token on this chain.
     *
     * Emits a {MintBridgeTokens} event.
     *
     * Note: This function is only called within `receiveTeleporterMessage`, which can only be
     * called by the Teleporter Messenger.
     */
    function _mintBridgeToken(
        bytes32 nativeBlockchainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) private {
        // The recipient cannot be the zero address.
        require(
            recipient != address(0),
            "ERC721Bridge: zero recipient address"
        );

        // Check that a bridge token exists for this native asset.
        // If not, one needs to be created by the delivery of a "createBridgeToken" message first
        // before this mint can be processed. Once the bridge token is create, this message
        // could then be retried to mint the tokens.
        address bridgeTokenAddress = nativeToBridgedNFT[nativeBlockchainID][
            nativeBridgeAddress
        ][nativeContractAddress];

        require(
            bridgeTokenAddress != address(0),
            "ERC721Bridge: bridge token does not exist"
        );

        // Mint the wrapped tokens.
        BridgeNFT(bridgeTokenAddress).mint(recipient, tokenId);
        emit MintBridgeNFT(bridgeTokenAddress, recipient, tokenId);
    }

    /**
     * @dev Teleporter message receiver for handling bridge tokens transfers back from another chain
     * and optionally routing them to a different third chain.
     *
     * Note: This function is only called within `receiveTeleporterMessage`, which can only be
     * called by the Teleporter Messenger.
     */
    function _transferBridgeToken(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) private {
        // Ensure that the destination blockchain ID is the current blockchain ID. No hops are supported at this time
        require(
            destinationBlockchainID == currentBlockchainID,
            "ERC721Bridge: invalid destination blockchain ID"
        );
        // Neither the recipient nor the destination bridge can be the zero address.
        require(
            recipient != address(0),
            "ERC721Bridge: zero recipient address"
        );
        require(
            destinationBridgeAddress != address(0),
            "ERC721Bridge: zero destination bridge address"
        );

        require(
            destinationBridgeAddress == address(this),
            "ERC721Bridge: invalid destination bridge address"
        );

        // Transfer tokens to the recipient.
        IERC721(nativeContractAddress).safeTransferFrom(
            address(this),
            recipient,
            tokenId
        );

        // Mark token as unlocked
        bridgedTokens[nativeContractAddress][tokenId] = false;
    }

    /**
     * @dev Processes a wrapped token transfer by burning the tokens and sending a Teleporter message
     * to the native chain and bridge of the wrapped asset that was burned.
     *
     * It is the caller's responsibility to ensure that the wrapped token contract is supported by this bridge instance.
     * Emits a {BridgeTokens} event.
     */
    function _processBridgedNftTransfer(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address bridgedNFTContractAddress,
        uint256 tokenId,
        address recipient
    ) private {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();
        // Burn the wrapped tokens to be bridged.
        // The bridge amount is the total amount minus the original fee amount. Even if the adjusted fee amount
        // is less than the original fee amount, the original amount is the portion that is spent out of the total
        // amount. We know that the burnFrom call will decrease the total supply by bridgeAmount because the
        // bridgeToken contract was deployed by this contract itself and does not implement "fee on burn" functionality.

        BridgeNFT bridgeNTF = BridgeNFT(bridgedNFTContractAddress);
        bridgeNTF.burn(tokenId);
        // If the destination chain ID is the native chain ID for the wrapped token, the bridge address must also match.
        // This is because you are not allowed to bridge a token within its native chain.
        bytes32 nativeBlockchainID = bridgeNTF.nativeBlockchainID();
        address nativeBridgeAddress = bridgeNTF.nativeBridge();

        // Curently, we don't support hopping to a destination chain that is not the native chain of the wrapped token
        // until we figure out a better way to handle the fee.
        require(
            destinationBlockchainID == nativeBlockchainID,
            "ERC721Bridge: invalid native destination blockchain ID"
        );

        require(
            destinationBridgeAddress == nativeBridgeAddress,
            "ERC721Bridge: invalid native destination bridge address"
        );

        // Send a message to the native chain and bridge of the wrapped asset that was burned.
        // The message includes the destination chain ID  and bridge contract, which will differ from the native
        // ones in the event that the tokens are being bridge from one non-native chain to another with two hops.
        bytes memory messageData = encodeTransferBridgeNFTData({
            destinationBlockchainID: destinationBlockchainID,
            destinationBridgeAddress: destinationBridgeAddress,
            nativeContractAddress: bridgeNTF.nativeAsset(),
            recipient: recipient,
            tokenId: tokenId
        });

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: nativeBlockchainID,
                destinationAddress: nativeBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit BridgedNFT({
            tokenContractAddress: bridgedNFTContractAddress,
            destinationBlockchainID: destinationBlockchainID,
            teleporterMessageID: messageID,
            destinationBridgeAddress: destinationBridgeAddress,
            recipient: recipient,
            tokenId: tokenId
        });
    }

    /**
     * @dev Encodes the parameters for the Transfer action to be decoded and executed on the destination.
     */
    function encodeTransferBridgeNFTData(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) public pure returns (bytes memory) {
        // ABI encode the Transfer action and corresponding parameters for the transferBridgeToken
        // call to to be decoded and executed on the destination.
        // solhint-disable-next-line func-named-parameters
        bytes memory paramsData = abi.encode(
            destinationBlockchainID,
            destinationBridgeAddress,
            nativeContractAddress,
            recipient,
            tokenId
        );
        return abi.encode(BridgeAction.Transfer, paramsData);
    }
}
