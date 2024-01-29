// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {BridgeNFT, ERC721Bridge, ERC721, IERC721, TeleporterMessageInput, TeleporterFeeInfo, IWarpMessenger, ITeleporterMessenger} from "../ERC721Bridge.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {UnitTestMockERC721} from "@mocks/UnitTestMockERC721.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

contract ERC721BridgeTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(
            hex"1111111111111111111111111111111111111111111111111111111111111111"
        );
    bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 private constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(
            hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
        );
    address private constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    string private constant _DEFAULT_TOKEN_NAME = "Test Token";
    string private constant _DEFAULT_SYMBOL = "TSTTK";
    string private constant _DEFAULT_TOKEN_URI = "https://example.com/ipfs";
    address private constant _DEFAULT_RECIPIENT =
        0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    ERC721Bridge public erc721Bridge;
    UnitTestMockERC721 public mockERC721;

    event BridgedNFT(
        address indexed tokenContractAddress,
        bytes32 indexed destinationBlockchainID,
        bytes32 indexed teleporterMessageID,
        address destinationBridgeAddress,
        address recipient,
        uint256 tokenId
    );

    event SubmitCreateBridgeNFT(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationBridgeAddress,
        address indexed nativeContractAddress,
        bytes32 teleporterMessageID
    );

    event CreateBridgeNFT(
        bytes32 indexed nativeBlockchainID,
        address indexed nativeBridgeAddress,
        address indexed nativeContractAddress,
        address bridgeTokenAddress
    );

    event MintBridgeNFT(
        address indexed contractAddress,
        address recipient,
        uint256 tokenId
    );

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS)
                    .latestVersion
                    .selector
            )
        );

        erc721Bridge = new ERC721Bridge(MOCK_TELEPORTER_REGISTRY_ADDRESS);
        mockERC721 = new UnitTestMockERC721();
    }

    function testSameBlockchainID() public {
        vm.expectRevert(
            _formatERC721BridgeErrorMessage("cannot bridge to same chain")
        );
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _MOCK_BLOCKCHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: 1
        });
    }

    function testZeroDestinationBridgeAddress() public {
        vm.expectRevert(
            _formatERC721BridgeErrorMessage("zero destination bridge address")
        );
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: address(0),
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: 1
        });
    }

    function testZeroNFTContractAddress() public {
        vm.expectRevert(
            _formatERC721BridgeErrorMessage("zero NFT contract address")
        );
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(0),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: 1
        });
    }

    function testZeroRecipient() public {
        vm.expectRevert(
            _formatERC721BridgeErrorMessage("zero recipient address")
        );
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: address(0),
            tokenId: 1
        });
    }

    function testInvalidBridgeNFTContract() public {
        vm.expectRevert(
            _formatERC721BridgeErrorMessage("invalid bridge NFT contract")
        );
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: 1
        });
    }

    function testAlreadyBridgedTokenId() public {
        uint256 tokenId = 1;

        // For the test purposes we will mint the NFT to the bridge contract directly  to simulate the NFT being already bridged.
        // Otherwise we would need to mint it to the user, then allow the bridge contract to transfer it, bridge it, and then
        // try to  bridge it again.
        mockERC721.mint(address(erc721Bridge), tokenId);

        // Bridge the NFT.
        _setUpMockERC721ContractValues(address(mockERC721));
        _submitCreateBridgeERC721(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC721)
        );

        vm.expectRevert(
            _formatERC721BridgeErrorMessage("token already bridged")
        );
        vm.prank(_DEFAULT_RECIPIENT);
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: tokenId
        });
    }

    function testNotOwnerOfTokenID() public {
        uint256 tokenId = 1;

        // Mint the token to a address address(1)
        mockERC721.mint(address(1), tokenId);

        // Bridge the NFT.
        _setUpMockERC721ContractValues(address(mockERC721));
        _submitCreateBridgeERC721(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC721)
        );

        vm.expectRevert(_formatERC721BridgeErrorMessage("invalid token ID"));
        vm.prank(_DEFAULT_RECIPIENT);
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: tokenId
        });
    }

    function testTransferNativeNFT() public {
        uint256 tokenId = 1;

        // Mint the token to the default recipient.
        // Approve the bridge contract to manage the token.
        mockERC721.mint(_DEFAULT_RECIPIENT, tokenId);
        vm.prank(_DEFAULT_RECIPIENT);
        mockERC721.approve(address(erc721Bridge), tokenId);

        // Bridge the NFT.
        _setUpMockERC721ContractValues(address(mockERC721));
        _submitCreateBridgeERC721(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC721)
        );

       _submitBridgeERC721(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC721),
            _DEFAULT_RECIPIENT,
            tokenId
        );

        assertEq(address(erc721Bridge), mockERC721.ownerOf(tokenId));
    }

    function testNewBridgeTokenMint() public {
        uint256 tokenId = 1;
        // Mint NFT to the default recipient.
        mockERC721.mint(_DEFAULT_RECIPIENT, tokenId);

        assertEq(_DEFAULT_RECIPIENT, mockERC721.ownerOf(tokenId));

        address bridgeTokenAddress = _setUpBridgeERC721({
            nativeBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            nativeBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nativeContractAddress: address(mockERC721),
            nativeName: _DEFAULT_TOKEN_NAME,
            nativeSymbol: _DEFAULT_SYMBOL,
            contractNonce: 1
        });

        bytes memory message = erc721Bridge.encodeMintBridgeNFTData(
            address(mockERC721),
            _DEFAULT_RECIPIENT,
            tokenId
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc721Bridge));
        emit MintBridgeNFT(bridgeTokenAddress, _DEFAULT_RECIPIENT, tokenId);
        _setUpExpectedMintCall(bridgeTokenAddress, _DEFAULT_RECIPIENT, tokenId);
        erc721Bridge.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            message
        );

        // Check the values and balance of the newly created ERC20.
        BridgeNFT newBridgeNFT = BridgeNFT(bridgeTokenAddress);
        assertEq(_DEFAULT_RECIPIENT, newBridgeNFT.ownerOf(tokenId));
        assertEq(1, newBridgeNFT.balanceOf(_DEFAULT_RECIPIENT));
        assertEq(_DEFAULT_TOKEN_NAME, newBridgeNFT.name());
        assertEq(_DEFAULT_SYMBOL, newBridgeNFT.symbol());
        assertEq(_DEFAULT_OTHER_CHAIN_ID, newBridgeNFT.nativeBlockchainID());
        assertEq(_DEFAULT_OTHER_BRIDGE_ADDRESS, newBridgeNFT.nativeBridge());
        assertEq(address(mockERC721), newBridgeNFT.nativeAsset());
    }

    // function testMintExistingBridgeToken() public {
    //     address recipient1 = address(56);
    //     uint256 amount1 = 654321;
    //     address recipient2 = address(57);
    //     uint256 amount2 = 123456;
    //     address bridgeTokenAddress = _setUpBridgeToken({
    //         nativeBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
    //         nativeBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
    //         nativeContractAddress: address(mockERC20),
    //         nativeName: _DEFAULT_TOKEN_NAME,
    //         nativeSymbol: _DEFAULT_SYMBOL,
    //         nativeDecimals: _DEFAULT_DECIMALS,
    //         contractNonce: 1
    //     });

    //     bytes memory message1 =
    //         erc20Bridge.encodeMintBridgeTokensData(address(mockERC20), recipient1, amount1);

    //     // Call mintBridgeTokens the first time, which should create the new bridge token.
    //     vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
    //     vm.expectEmit(true, true, true, true, address(erc20Bridge));
    //     emit MintBridgeTokens(bridgeTokenAddress, recipient1, amount1);
    //     erc20Bridge.receiveTeleporterMessage(
    //         _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, message1
    //     );

    //     bytes memory message2 =
    //         erc20Bridge.encodeMintBridgeTokensData(address(mockERC20), recipient2, amount2);

    //     // Call mintBridgeTokens the second time, which should mint additional of the existing token.
    //     vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
    //     vm.expectEmit(true, true, true, true, address(erc20Bridge));
    //     emit MintBridgeTokens(bridgeTokenAddress, recipient2, amount2);
    //     erc20Bridge.receiveTeleporterMessage(
    //         _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, message2
    //     );

    //     // Check the values and balance of the newly created ERC20.
    //     BridgeToken newBridgeToken = BridgeToken(bridgeTokenAddress);
    //     assertEq(amount1 + amount2, newBridgeToken.totalSupply());
    //     assertEq(amount1, newBridgeToken.balanceOf(recipient1));
    //     assertEq(amount2, newBridgeToken.balanceOf(recipient2));
    //     assertEq(_DEFAULT_TOKEN_NAME, newBridgeToken.name());
    //     assertEq(_DEFAULT_SYMBOL, newBridgeToken.symbol());
    //     assertEq(_DEFAULT_DECIMALS, newBridgeToken.decimals());
    //     assertEq(_DEFAULT_OTHER_CHAIN_ID, newBridgeToken.nativeBlockchainID());
    //     assertEq(_DEFAULT_OTHER_BRIDGE_ADDRESS, newBridgeToken.nativeBridge());
    //     assertEq(address(mockERC20), newBridgeToken.nativeAsset());
    // }

    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert(
            "TeleporterUpgradeable: zero teleporter registry address"
        );
        new ERC721Bridge(address(0));
    }

    function _initMockTeleporterRegistry() internal {
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS)
                    .latestVersion
                    .selector
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry.getVersionFromAddress.selector,
                (MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry.getAddressFromVersion.selector,
                (1)
            ),
            abi.encode(MOCK_TELEPORTER_MESSENGER_ADDRESS)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry.getLatestTeleporter.selector
            ),
            abi.encode(ITeleporterMessenger(MOCK_TELEPORTER_MESSENGER_ADDRESS))
        );
    }

    function _setUpExpectedMintCall(
        address nftContract,
        address recipient,
        uint256 tokenId
    ) private {
        vm.expectCall(
            nftContract,
            abi.encodeCall(UnitTestMockERC721.mint, (recipient, tokenId))
        );
    }

    function _setUpExpectedTransferFromCall(
        address nftContract,
        address from,
        address to,
        uint256 tokenId
    ) private {
        vm.expectCall(
            nftContract,
            abi.encode("safeTransferFrom(address,address,uint256)", string.concat("(", Strings.toHexString(from), ",", Strings.toHexString(to), ",", Strings.toHexString(tokenId), ")"))
        );
    }

    // Calls submitCreateBridgeERC721 of the test's ERC721Bridge instance to add the specified
    // token to be able to be sent to the specified destination bridge. Checks that the expected
    // call to the Teleporter contract is made and that the expected event is emitted. This is
    // required before attempting to call bridgeTokens for the given token and bridge.
    function _submitCreateBridgeERC721(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress
    ) private {
        ERC721 nativeToken = ERC721(nativeContractAddress);

        TeleporterMessageInput
            memory expectedMessageInput = TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: erc721Bridge
                    .CREATE_BRIDGE_TOKENS_REQUIRED_GAS(),
                allowedRelayerAddresses: new address[](0),
                message: erc721Bridge.encodeCreateBridgeNFTData(
                    nativeContractAddress,
                    nativeToken.name(),
                    nativeToken.symbol()
                )
            });

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            ),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            )
        );

        vm.expectEmit(true, true, true, true, address(erc721Bridge));
        emit SubmitCreateBridgeNFT(
            destinationBlockchainID,
            destinationBridgeAddress,
            address(nativeToken),
            _MOCK_MESSAGE_ID
        );

        erc721Bridge.submitCreateBridgeERC721(
            destinationBlockchainID,
            destinationBridgeAddress,
            ERC721(nativeContractAddress)
        );
    }

    // Calls bridgeERC721 of the test's ERC721Bridge instance to bridge the specified token to the
    // specified destination bridge. Checks that the expected call to the Teleporter contract is
    // made and that the expected event is emitted.
    function _submitBridgeERC721(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 tokenId
    ) private {
        vm.expectCall(
            nativeContractAddress,
            abi.encodeWithSignature("safeTransferFrom(address,address,uint256)", recipient, address(erc721Bridge), tokenId)
        );

        TeleporterMessageInput
            memory expectedMessageInput = TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: erc721Bridge
                    .MINT_BRIDGE_TOKENS_REQUIRED_GAS(),
                allowedRelayerAddresses: new address[](0),
                message: erc721Bridge.encodeMintBridgeNFTData(
                    nativeContractAddress,
                    recipient,
                    tokenId
                )
            });

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            ),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            )
        );

        vm.expectEmit(true, true, true, true, address(erc721Bridge));
        emit BridgedNFT(
            nativeContractAddress,
            destinationBlockchainID,
            _MOCK_MESSAGE_ID,
            destinationBridgeAddress,
            recipient,
            tokenId
        );

        vm.prank(_DEFAULT_RECIPIENT);
        erc721Bridge.bridgeERC721({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nftContractAddress: address(mockERC721),
            recipient: _DEFAULT_RECIPIENT,
            tokenId: tokenId
        });
    }

    function _setUpBridgeERC721(
        bytes32 nativeBlockchainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol,
        uint8 contractNonce
    ) private returns (address) {
        address expectedBridgeTokenAddress = _deriveExpectedContractAddress(
            address(erc721Bridge),
            contractNonce
        );
        bytes memory message = erc721Bridge.encodeCreateBridgeNFTData(
            nativeContractAddress,
            nativeName,
            nativeSymbol
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc721Bridge));
        emit CreateBridgeNFT(
            nativeBlockchainID,
            nativeBridgeAddress,
            nativeContractAddress,
            expectedBridgeTokenAddress
        );
        erc721Bridge.receiveTeleporterMessage(
            nativeBlockchainID,
            nativeBridgeAddress,
            message
        );
        return expectedBridgeTokenAddress;
    }

    function _setUpMockERC721ContractValues(address nftContract) private {
        ERC721 token = ERC721(nftContract);
        vm.mockCall(
            nftContract,
            abi.encodeCall(token.name, ()),
            abi.encode(_DEFAULT_TOKEN_NAME)
        );
        vm.expectCall(nftContract, abi.encodeCall(token.name, ()));
        vm.mockCall(
            nftContract,
            abi.encodeCall(token.symbol, ()),
            abi.encode(_DEFAULT_SYMBOL)
        );
        vm.expectCall(nftContract, abi.encodeCall(token.symbol, ()));
    }

    function _deriveExpectedContractAddress(
        address creator,
        uint8 nonce
    ) private pure returns (address) {
        // Taken from https://ethereum.stackexchange.com/questions/760/how-is-the-address-of-an-ethereum-contract-computed
        // We must use encodePacked rather than encode so that the parameters are not padded with extra zeros.
        return
            address(
                uint160(
                    uint256(
                        keccak256(
                            abi.encodePacked(
                                bytes1(0xd6),
                                bytes1(0x94),
                                creator,
                                bytes1(nonce)
                            )
                        )
                    )
                )
            );
    }

    function _formatERC721BridgeErrorMessage(
        string memory errorMessage
    ) private pure returns (bytes memory) {
        return bytes(string.concat("ERC721Bridge: ", errorMessage));
    }
}
