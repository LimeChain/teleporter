package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	bridgenft "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC721Bridge/BridgeNFT"
	erc721bridge "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC721Bridge/ERC721Bridge"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC721NativeTokenBridge(network interfaces.Network) {
	sourceSubnet := network.GetPrimaryNetworkInfo()
	_, destSubnet := utils.GetTwoSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	sourceTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		sourceSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())

	destTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		destSubnet.RPCClient,
	)

	Expect(err).Should(BeNil())

	// Deploy an ERC20 to subnet A
	nativeERC721Address, nativeERC721 := utils.DeployExampleERC721(
		context.Background(),
		fundedKey,
		sourceSubnet,
	)

	// Deploy the ERC721 bridge to source subnet
	sourceERC721BridgeAddress, sourceERC721Bridge := utils.DeployERC721Bridge(ctx, fundedKey, sourceSubnet)
	// Deploy the ERC20 bridge to destination subnet
	destERC721BridgeAddress, destERC721Bridge := utils.DeployERC721Bridge(ctx, fundedKey, destSubnet)

	// Mint an ERC721 token to the funded address
	tokenId := big.NewInt(1)
	utils.ERC721Mint(
		ctx,
		nativeERC721,
		tokenId,
		sourceSubnet,
		fundedKey,
	)

	// Check owner of ERC721 token
	tokenOwner, err := nativeERC721.OwnerOf(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(tokenOwner).Should(Equal(fundedAddress))

	// Approve source ERC721 bridge as operator of the ERC721 token
	utils.ERC721Approve(
		ctx,
		nativeERC721,
		sourceERC721BridgeAddress,
		tokenId,
		sourceSubnet,
		fundedKey,
	)

	// Check bridge is approved as operator of the ERC721 token
	tokenOperator, err := nativeERC721.GetApproved(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(tokenOperator).Should(Equal(sourceERC721BridgeAddress))

	// Submit a CreateBridgeERC721 message to the source subnet
	receipt, messageID := SubmitCreateBridgeERC721(
		ctx,
		sourceSubnet,
		destSubnet.BlockchainID,
		destERC721BridgeAddress,
		nativeERC721Address,
		fundedAddress,
		fundedKey,
		sourceERC721Bridge,
		sourceTeleporterMessenger,
	)

	// Relay message
	network.RelayMessage(ctx, receipt, sourceSubnet, destSubnet, true)

	// Check Teleporter message received on the destination
	delivered, err := destTeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		messageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check bridged nft contract has been mapped to the teleporter
	bridgedERC721Address, err := destERC721Bridge.NativeToBridgedNFT(
		&bind.CallOpts{},
		sourceSubnet.BlockchainID,
		sourceERC721BridgeAddress,
		nativeERC721Address,
	)
	Expect(err).Should(BeNil())
	Expect(bridgedERC721Address).ShouldNot(Equal(common.Address{}))

	// Bridge the ERC721 token to the destination subnet
	receipt, messageID = bridgeERC721(
		ctx,
		sourceSubnet,
		destSubnet.BlockchainID,
		destERC721BridgeAddress,
		nativeERC721Address,
		tokenId,
		fundedAddress,
		fundedKey,
		sourceERC721Bridge,
		true,
		sourceSubnet.BlockchainID,
		sourceTeleporterMessenger,
	)

	// Relay message
	destChainReceipt := network.RelayMessage(ctx, receipt, sourceSubnet, destSubnet, true)
	_, err = utils.GetEventFromLogs(
		destChainReceipt.Logs,
		destSubnet.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination
	delivered, err = destTeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check token owner is the bridge on native subnet
	tokenOwner, err = nativeERC721.OwnerOf(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(tokenOwner).Should(Equal(sourceERC721BridgeAddress))

	// Check bridged token has been minted on the destination
	destBridgeNFT, err := bridgenft.NewBridgeNFT(bridgedERC721Address, destSubnet.RPCClient)
	Expect(err).Should(BeNil())

	// Check balance of funded address
	balance, err := destBridgeNFT.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(big.NewInt(1)))
	// Check owner of bridged token
	destTokenOwner, err := destBridgeNFT.OwnerOf(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(destTokenOwner).Should(Equal(fundedAddress))

	// Approve destination ERC721 bridge as operator of the bridged ERC721 token
	approveBridgeNFT(
		ctx,
		destBridgeNFT,
		destERC721BridgeAddress,
		tokenId,
		destSubnet,
		fundedKey,
	)

	// Check destination bridge is approved as operator of the bridged ERC721 token
	bridgedTokenOperator, err := destBridgeNFT.GetApproved(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(bridgedTokenOperator).Should(Equal(destERC721BridgeAddress))

	// Bridge the ERC721 token back to the source subnet
	destChainReceipt, destMessageID := bridgeERC721(
		ctx,
		destSubnet,
		sourceSubnet.BlockchainID,
		sourceERC721BridgeAddress,
		bridgedERC721Address,
		tokenId,
		fundedAddress,
		fundedKey,
		destERC721Bridge,
		false,
		sourceSubnet.BlockchainID,
		destTeleporterMessenger,
	)

	// Relay message
	sourceChainReceipt := network.RelayMessage(ctx, destChainReceipt, destSubnet, sourceSubnet, true)
	logs, err := utils.GetEventFromLogs(
		sourceChainReceipt.Logs,
		destSubnet.TeleporterMessenger.ParseReceiveCrossChainMessage)

	log.Info("logs", "logs", logs)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the source
	delivered, err = sourceTeleporterMessenger.MessageReceived(&bind.CallOpts{}, destMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check balance of funded address on the destination
	// For some reason the value of balance is <*big.Int | 0x1400115f460>: {neg: false, abs: []}
	// while value of big.NewInt(0) is <*big.Int | 0x1400115f480>: {neg: false, abs: nil}
	// so we need to compare the bytes
	balance, err = destBridgeNFT.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	Expect(balance.Bytes()).Should(Equal(big.NewInt(0).Bytes()))

	// Check if token has been burned on the destination
	// ERC721 contract should revert with "ERC721: invalid token ID"
	// if the token has been burned and "owner" is zero address
	_, err = destBridgeNFT.OwnerOf(&bind.CallOpts{}, tokenId)
	Expect(err.Error()).Should(ContainSubstring("ERC721: invalid token ID"))

	// Check token has been transferred to the funded address on the native subnet
	tokenOwner, err = nativeERC721.OwnerOf(&bind.CallOpts{}, tokenId)
	Expect(err).Should(BeNil())
	Expect(tokenOwner).Should(Equal(fundedAddress))
}

func SubmitCreateBridgeERC721(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	nativeToken common.Address,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *erc721bridge.ERC721Bridge,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.SubmitCreateBridgeNFT(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
		nativeToken,
		common.HexToAddress("0x0"),
		big.NewInt(0),
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, source, tx)

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))

	log.Info("Successfully SubmitCreateBridgeERC721",
		"txHash", tx.Hash().Hex(),
		"messageID", event.MessageID)

	return receipt, event.MessageID
}

func bridgeERC721(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	token common.Address,
	tokenId *big.Int,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *erc721bridge.ERC721Bridge,
	isNative bool,
	nativeTokenChainID ids.ID,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.BridgeToken(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
		token,
		fundedAddress,
		tokenId,
		common.HexToAddress("0x0"),
		big.NewInt(0),
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, source, tx)

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	if isNative {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))
	} else {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(nativeTokenChainID[:]))
	}

	return receipt, event.MessageID
}

func approveBridgeNFT(
	ctx context.Context,
	token *bridgenft.BridgeNFT,
	operator common.Address,
	tokenId *big.Int,
	source interfaces.SubnetTestInfo,
	senderKey *ecdsa.PrivateKey,
) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := token.Approve(opts, operator, tokenId)
	Expect(err).Should(BeNil())

	utils.WaitForTransactionSuccess(ctx, source, tx)
}
