#!/usr/bin/env bash

set -e # Stop on first error

# Set up contracts
cd contracts

bridge_c_deploy_result=$(forge create --private-key $user_private_key \
    src/CrossChainApplications/examples/ERC721Bridge/ERC721Bridge.sol:ERC721Bridge \
    --constructor-args $c_chain_teleporter_registry_address \
    --rpc-url $c_chain_rpc_url)
bridge_contract_address_c_chain=$(parseContractAddress "$bridge_c_deploy_result")

bridge_a_deploy_result=$(forge create --private-key $user_private_key \
    src/CrossChainApplications/examples/ERC721Bridge/ERC721Bridge.sol:ERC721Bridge \
    --constructor-args $subnet_a_teleporter_registry_address \
    --rpc-url $subnet_a_rpc_url)
bridge_contract_address_subnet_a=$(parseContractAddress "$bridge_a_deploy_result")

erc721_deploy_result=$(forge create --private-key $user_private_key \
src/Mocks/ExampleERC721.sol:ExampleERC721 \
--rpc-url $c_chain_rpc_url)
erc721_contract_address_c_chain=$(parseContractAddress "$erc721_deploy_result")

echo "ERC721Birdge contract deployed to $bridge_contract_address_c_chain on C-Chain"
echo "ERC721Birdge contract deployed to $bridge_contract_address_subnet_a on Subnet A"
echo "ExampleERC721 contract deployed to $erc721_contract_address_c_chain on the C-Chain"

# send request to create bridgeNFT contract on Subnet A
# function submitCreateBridgeNFT(
#     bytes32 destinationBlockchainID,
#     address destinationBridgeAddress,
#     ExampleERC721 nativeContract,
#     address messageFeeAsset,
#     uint256 messageFeeAmount
# )
echo "Sending request to create bridgeNFT contract on Subnet A"
cast send --private-key $user_private_key \
    $bridge_contract_address_c_chain \
    "submitCreateBridgeNFT(bytes32,address,address,address,uint256)" \
    $subnet_a_blockchain_id_hex \
    $bridge_contract_address_subnet_a \
    $erc721_contract_address_c_chain \
    "0x0000000000000000000000000000000000000000" \
    0 \
    --rpc-url $c_chain_rpc_url > /dev/null

sleep 5
# Check the bridgeNFT contract is created on Subnet A
# mapping(bytes32 nativeBlockchainID => mapping(address nativeBridgeAddress => mapping(address nativeTokenAddress => address bridgeNFTAddress))) public nativeToBridgedNFT;
bridge_nft_address=$(cast call $bridge_contract_address_subnet_a \
  "nativeToBridgedNFT(bytes32,address,address)(address)" \
  $c_chain_blockchain_id_hex \
  $bridge_contract_address_c_chain \
  $erc721_contract_address_c_chain \
  --rpc-url $subnet_a_rpc_url)

if [[ $bridge_nft_address == "0x0000000000000000000000000000000000000000" ]]; then
    echo "BridgeNFT contract is not created on Subnet A"
    exit 1
fi

echo "BridgeNFT contract created on Subnet A at $bridge_nft_address"

# Mint an ERC721 token on the C-Chain
token_id=1
echo "Minting ERC721 token with ID $token_id on the C-Chain"
cast send --private-key $user_private_key $erc721_contract_address_c_chain "mint(uint256)" $token_id --rpc-url $c_chain_rpc_url > /dev/null

# Check owner of the ERC721 token on C-Chain is the user
result=$(cast call $erc721_contract_address_c_chain "ownerOf(uint256)(address)" $token_id --rpc-url $c_chain_rpc_url)
if [[ $result != $user_address ]]; then
    echo "ERC721 token with ID $token_id is not owned by the user on the C-Chain"
    exit 1
fi
echo "ERC721 token with ID $token_id is owned by the user $user_address on the C-Chain"

echo "Approving the ERC721Bridge contract to transfer the ERC721 token on the C-Chain"
# Approve the ERC721Bridge contract as operator of the ERC721 token
cast send --private-key $user_private_key \
    $erc721_contract_address_c_chain \
    "approve(address,uint256)" \
    $bridge_contract_address_c_chain \
    $token_id \
    --rpc-url $c_chain_rpc_url > /dev/null

# Check the ERC721Bridge contract is approved to transfer the ERC721 token
echo "Checking the ERC721Bridge contract is approved to transfer the ERC721 token on the C-Chain"
approved=$(cast call $erc721_contract_address_c_chain "getApproved(uint256)(address)" $token_id --rpc-url $c_chain_rpc_url)
if [[ $approved != $bridge_contract_address_c_chain ]]; then
    echo "ERC721Bridge contract is not approved to transfer the ERC721 token on the C-Chain"
    exit 1
fi
echo "ERC721Bridge contract $bridge_contract_address_c_chain is approved to transfer the ERC721 token on the C-Chain"

# Call bridgeToken on the ERC721Bridge contract to bridge the ERC721 token to Subnet A
# function bridgeToken(
#     bytes32 destinationBlockchainID,
#     address destinationBridgeAddress,
#     address nftContractAddress,
#     address recipient,
#     uint256 tokenId,
#     address messageFeeAsset,
#     uint256 messageFeeAmount
# )
echo "Bridging the ERC721 token to Subnet A"
cast send --private-key $user_private_key \
    $bridge_contract_address_c_chain \
    "bridgeToken(bytes32,address,address,address,uint256,address,uint256)" \
    $subnet_a_blockchain_id_hex \
    $bridge_contract_address_subnet_a \
    $erc721_contract_address_c_chain \
    $user_address \
    $token_id  \
    "0x0000000000000000000000000000000000000000" \
    0 \
    --rpc-url $c_chain_rpc_url > /dev/null

sleep 5
# Check owner of the ERC721 token on C-Chain is the ERC721Bridge contract
echo "Checking the ERC721Bridge contract is the owner of the ERC721 token on the C-Chain"
result=$(cast call $erc721_contract_address_c_chain "ownerOf(uint256)(address)" $token_id --rpc-url $c_chain_rpc_url)
if [[ $result != $bridge_contract_address_c_chain ]]; then
    echo "ERC721 token $token_id is not owned by the ERC721Bridge contract on the C-Chain"
    exit 1
fi
echo "ERC721 token $token_id is owned by the ERC721Bridge contract ($bridge_contract_address_c_chain) on the C-Chain"

# Check the ERC721 token is minted on Subnet A to the user
echo "Checking the ERC721 token is minted on Subnet A to the user"
result=$(cast call $bridge_nft_address "ownerOf(uint256)(address)" $token_id --rpc-url $subnet_a_rpc_url)
if [[ $result != $user_address ]]; then
    echo "ERC721 token with ID $token_id is not owned by the user on Subnet A"
    exit 1
fi

# Approve the ERC721Bridge contract as operator of the ERC721 token on Subnet A
echo "Approving the ERC721Bridge contract to transfer the ERC721 token on Subnet A"
cast send --private-key $user_private_key \
    $bridge_nft_address \
    "approve(address,uint256)" \
    $bridge_contract_address_subnet_a \
    $token_id \
    --rpc-url $subnet_a_rpc_url > /dev/null

echo "Checking the ERC721Bridge contract is approved to transfer the ERC721 token on Subnet A"
approved=$(cast call $bridge_nft_address "getApproved(uint256)(address)" $token_id --rpc-url $subnet_a_rpc_url)
if [[ $approved != $bridge_contract_address_subnet_a ]]; then
    echo "ERC721Bridge contract is not approved to transfer the ERC721 token on Subnet A"
    exit 1
fi
echo "ERC721Bridge contract $bridge_contract_address_subnet_a is approved to transfer the ERC721 token on Subnet A"

# Send the ERC721 token back to the C-Chain
echo "Sending the ERC721 token back to the C-Chain"
cast send --private-key $user_private_key \
    $bridge_contract_address_subnet_a \
    "bridgeToken(bytes32,address,address,address,uint256,address,uint256)" \
    $c_chain_blockchain_id_hex \
    $bridge_contract_address_c_chain \
    $bridge_nft_address \
    $user_address \
    $token_id \
    "0x0000000000000000000000000000000000000000" \
    0 \
    --rpc-url $subnet_a_rpc_url > /dev/null


sleep 5

# To check if the token has been burned on Subnet A
# cast call $bridge_nft_address "ownerOf(uint256)(address)" $token_id --rpc-url $subnet_a_rpc_url
# This call should revert with "invalid token ID" error

# Check owner of the ERC721 token on C-Chain is the ERC721Bridge contract
# Owner should be the user address
echo "Checking the ERC721Bridge contract is the owner of the ERC721 token on C chain"
result=$(cast call $erc721_contract_address_c_chain "ownerOf(uint256)(address)" $token_id --rpc-url $c_chain_rpc_url)
if [[ $result != $user_address ]]; then
    echo "ERC721 token $token_id is not owned by the user on the C-Chain"
    exit 1
fi
echo "Token with ID $token_id is transferred back to the user $user_address on the C-Chain"
