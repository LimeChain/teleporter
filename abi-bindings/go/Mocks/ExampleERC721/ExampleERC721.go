// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleerc721

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExampleERC721MetaData contains all meta data concerning the ExampleERC721 contract.
var ExampleERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816000908162000063919062000120565b50600162000072828262000120565b505050620001ec565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000a657607f821691505b602082108103620000c757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200011b57600081815260208120601f850160051c81016020861015620000f65750805b601f850160051c820191505b81811015620001175782815560010162000102565b5050505b505050565b81516001600160401b038111156200013c576200013c6200007b565b62000154816200014d845462000091565b84620000cd565b602080601f8311600181146200018c5760008415620001735750858301515b600019600386901b1c1916600185901b17855562000117565b600085815260208120601f198616915b82811015620001bd578886015182559484019460019091019084016200019c565b5085821015620001dc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61152780620001fc6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806370a0823111610097578063a22cb46511610066578063a22cb4651461023f578063b88d4fde14610252578063c87b56dd14610265578063e985e9c51461027857600080fd5b806370a08231146101ce57806395d89b41146101ef5780639abc8320146101f7578063a0712d681461022c57600080fd5b806323b872dd116100d357806323b872dd1461018257806342842e0e1461019557806342966c68146101a85780636352211e146101bb57600080fd5b806301ffc9a71461010557806306fdde031461012d578063081812fc14610142578063095ea7b31461016d575b600080fd5b610118610113366004611077565b61028b565b60405190151581526020015b60405180910390f35b6101356102dd565b60405161012491906110e4565b6101556101503660046110f7565b61036f565b6040516001600160a01b039091168152602001610124565b61018061017b36600461112c565b610396565b005b610180610190366004611156565b6104b0565b6101806101a3366004611156565b6104e2565b6101806101b63660046110f7565b6104fd565b6101556101c93660046110f7565b61052e565b6101e16101dc366004611192565b61058e565b604051908152602001610124565b610135610614565b60408051808201909152601981527868747470733a2f2f6578616d706c652e636f6d2f697066732f60381b6020820152610135565b61018061023a3660046110f7565b610623565b61018061024d3660046111ad565b610689565b6101806102603660046111ff565b610698565b6101356102733660046110f7565b6106d0565b6101186102863660046112db565b610765565b60006001600160e01b031982166380ac58cd60e01b14806102bc57506001600160e01b03198216635b5e139f60e01b145b806102d757506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546102ec9061130e565b80601f01602080910402602001604051908101604052809291908181526020018280546103189061130e565b80156103655780601f1061033a57610100808354040283529160200191610365565b820191906000526020600020905b81548152906001019060200180831161034857829003601f168201915b5050505050905090565b600061037a82610793565b506000908152600460205260409020546001600160a01b031690565b60006103a18261052e565b9050806001600160a01b0316836001600160a01b0316036104135760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061042f575061042f8133610765565b6104a15760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161040a565b6104ab83836107f2565b505050565b6104bb335b82610860565b6104d75760405162461bcd60e51b815260040161040a90611348565b6104ab8383836108bf565b6104ab83838360405180602001604052806000815250610698565b610506336104b5565b6105225760405162461bcd60e51b815260040161040a90611348565b61052b81610a30565b50565b6000818152600260205260408120546001600160a01b0316806102d75760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161040a565b60006001600160a01b0382166105f85760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161040a565b506001600160a01b031660009081526003602052604090205490565b6060600180546102ec9061130e565b6000818152600260205260409020546001600160a01b03161561067f5760405162461bcd60e51b8152602060048201526014602482015273151bdad95b88185b1c9958591e481b5a5b9d195960621b604482015260640161040a565b61052b3382610ad3565b610694338383610c6c565b5050565b6106a23383610860565b6106be5760405162461bcd60e51b815260040161040a90611348565b6106ca84848484610d3a565b50505050565b60606106db82610793565b600061071360408051808201909152601981527868747470733a2f2f6578616d706c652e636f6d2f697066732f60381b602082015290565b90506000815111610733576040518060200160405280600081525061075e565b8061073d84610d6d565b60405160200161074e929190611395565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000818152600260205260409020546001600160a01b031661052b5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161040a565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108278261052e565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60008061086c8361052e565b9050806001600160a01b0316846001600160a01b0316148061089357506108938185610765565b806108b75750836001600160a01b03166108ac8461036f565b6001600160a01b0316145b949350505050565b826001600160a01b03166108d28261052e565b6001600160a01b0316146108f85760405162461bcd60e51b815260040161040a906113c4565b6001600160a01b03821661095a5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161040a565b6109678383836001610e00565b826001600160a01b031661097a8261052e565b6001600160a01b0316146109a05760405162461bcd60e51b815260040161040a906113c4565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000610a3b8261052e565b9050610a4b816000846001610e00565b610a548261052e565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6001600160a01b038216610b295760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161040a565b6000818152600260205260409020546001600160a01b031615610b8e5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161040a565b610b9c600083836001610e00565b6000818152600260205260409020546001600160a01b031615610c015760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161040a565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b031603610ccd5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161040a565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610d458484846108bf565b610d5184848484610e88565b6106ca5760405162461bcd60e51b815260040161040a90611409565b60606000610d7a83610f89565b600101905060008167ffffffffffffffff811115610d9a57610d9a6111e9565b6040519080825280601f01601f191660200182016040528015610dc4576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610dce57509392505050565b60018111156106ca576001600160a01b03841615610e46576001600160a01b03841660009081526003602052604081208054839290610e40908490611471565b90915550505b6001600160a01b038316156106ca576001600160a01b03831660009081526003602052604081208054839290610e7d908490611484565b909155505050505050565b60006001600160a01b0384163b15610f7e57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610ecc903390899088908890600401611497565b6020604051808303816000875af1925050508015610f07575060408051601f3d908101601f19168201909252610f04918101906114d4565b60015b610f64573d808015610f35576040519150601f19603f3d011682016040523d82523d6000602084013e610f3a565b606091505b508051600003610f5c5760405162461bcd60e51b815260040161040a90611409565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506108b7565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310610fc85772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310610ff4576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061101257662386f26fc10000830492506010015b6305f5e100831061102a576305f5e100830492506008015b612710831061103e57612710830492506004015b60648310611050576064830492506002015b600a83106102d75760010192915050565b6001600160e01b03198116811461052b57600080fd5b60006020828403121561108957600080fd5b813561075e81611061565b60005b838110156110af578181015183820152602001611097565b50506000910152565b600081518084526110d0816020860160208601611094565b601f01601f19169290920160200192915050565b60208152600061075e60208301846110b8565b60006020828403121561110957600080fd5b5035919050565b80356001600160a01b038116811461112757600080fd5b919050565b6000806040838503121561113f57600080fd5b61114883611110565b946020939093013593505050565b60008060006060848603121561116b57600080fd5b61117484611110565b925061118260208501611110565b9150604084013590509250925092565b6000602082840312156111a457600080fd5b61075e82611110565b600080604083850312156111c057600080fd5b6111c983611110565b9150602083013580151581146111de57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561121557600080fd5b61121e85611110565b935061122c60208601611110565b925060408501359150606085013567ffffffffffffffff8082111561125057600080fd5b818701915087601f83011261126457600080fd5b813581811115611276576112766111e9565b604051601f8201601f19908116603f0116810190838211818310171561129e5761129e6111e9565b816040528281528a60208487010111156112b757600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b600080604083850312156112ee57600080fd5b6112f783611110565b915061130560208401611110565b90509250929050565b600181811c9082168061132257607f821691505b60208210810361134257634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b600083516113a7818460208801611094565b8351908301906113bb818360208801611094565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b818103818111156102d7576102d761145b565b808201808211156102d7576102d761145b565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906114ca908301846110b8565b9695505050505050565b6000602082840312156114e657600080fd5b815161075e8161106156fea26469706673582212207618093cd9190546702532e2e804ed8bf10126d95dd5eabad9cdae84677b615064736f6c63430008120033",
}

// ExampleERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleERC721MetaData.ABI instead.
var ExampleERC721ABI = ExampleERC721MetaData.ABI

// ExampleERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleERC721MetaData.Bin instead.
var ExampleERC721Bin = ExampleERC721MetaData.Bin

// DeployExampleERC721 deploys a new Ethereum contract, binding an instance of ExampleERC721 to it.
func DeployExampleERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExampleERC721, error) {
	parsed, err := ExampleERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleERC721{ExampleERC721Caller: ExampleERC721Caller{contract: contract}, ExampleERC721Transactor: ExampleERC721Transactor{contract: contract}, ExampleERC721Filterer: ExampleERC721Filterer{contract: contract}}, nil
}

// ExampleERC721 is an auto generated Go binding around an Ethereum contract.
type ExampleERC721 struct {
	ExampleERC721Caller     // Read-only binding to the contract
	ExampleERC721Transactor // Write-only binding to the contract
	ExampleERC721Filterer   // Log filterer for contract events
}

// ExampleERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleERC721Session struct {
	Contract     *ExampleERC721    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleERC721CallerSession struct {
	Contract *ExampleERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ExampleERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleERC721TransactorSession struct {
	Contract     *ExampleERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExampleERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleERC721Raw struct {
	Contract *ExampleERC721 // Generic contract binding to access the raw methods on
}

// ExampleERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleERC721CallerRaw struct {
	Contract *ExampleERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ExampleERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleERC721TransactorRaw struct {
	Contract *ExampleERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleERC721 creates a new instance of ExampleERC721, bound to a specific deployed contract.
func NewExampleERC721(address common.Address, backend bind.ContractBackend) (*ExampleERC721, error) {
	contract, err := bindExampleERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721{ExampleERC721Caller: ExampleERC721Caller{contract: contract}, ExampleERC721Transactor: ExampleERC721Transactor{contract: contract}, ExampleERC721Filterer: ExampleERC721Filterer{contract: contract}}, nil
}

// NewExampleERC721Caller creates a new read-only instance of ExampleERC721, bound to a specific deployed contract.
func NewExampleERC721Caller(address common.Address, caller bind.ContractCaller) (*ExampleERC721Caller, error) {
	contract, err := bindExampleERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721Caller{contract: contract}, nil
}

// NewExampleERC721Transactor creates a new write-only instance of ExampleERC721, bound to a specific deployed contract.
func NewExampleERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ExampleERC721Transactor, error) {
	contract, err := bindExampleERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721Transactor{contract: contract}, nil
}

// NewExampleERC721Filterer creates a new log filterer instance of ExampleERC721, bound to a specific deployed contract.
func NewExampleERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ExampleERC721Filterer, error) {
	contract, err := bindExampleERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721Filterer{contract: contract}, nil
}

// bindExampleERC721 binds a generic wrapper to an already deployed contract.
func bindExampleERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC721 *ExampleERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC721.Contract.ExampleERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC721 *ExampleERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC721.Contract.ExampleERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC721 *ExampleERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC721.Contract.ExampleERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC721 *ExampleERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC721 *ExampleERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC721 *ExampleERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ExampleERC721 *ExampleERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ExampleERC721 *ExampleERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ExampleERC721.Contract.BalanceOf(&_ExampleERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ExampleERC721 *ExampleERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ExampleERC721.Contract.BalanceOf(&_ExampleERC721.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() pure returns(string)
func (_ExampleERC721 *ExampleERC721Caller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() pure returns(string)
func (_ExampleERC721 *ExampleERC721Session) BaseUri() (string, error) {
	return _ExampleERC721.Contract.BaseUri(&_ExampleERC721.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() pure returns(string)
func (_ExampleERC721 *ExampleERC721CallerSession) BaseUri() (string, error) {
	return _ExampleERC721.Contract.BaseUri(&_ExampleERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ExampleERC721.Contract.GetApproved(&_ExampleERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ExampleERC721.Contract.GetApproved(&_ExampleERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ExampleERC721 *ExampleERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ExampleERC721 *ExampleERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ExampleERC721.Contract.IsApprovedForAll(&_ExampleERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ExampleERC721 *ExampleERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ExampleERC721.Contract.IsApprovedForAll(&_ExampleERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC721 *ExampleERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC721 *ExampleERC721Session) Name() (string, error) {
	return _ExampleERC721.Contract.Name(&_ExampleERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC721 *ExampleERC721CallerSession) Name() (string, error) {
	return _ExampleERC721.Contract.Name(&_ExampleERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ExampleERC721.Contract.OwnerOf(&_ExampleERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ExampleERC721 *ExampleERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ExampleERC721.Contract.OwnerOf(&_ExampleERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleERC721 *ExampleERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleERC721 *ExampleERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExampleERC721.Contract.SupportsInterface(&_ExampleERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleERC721 *ExampleERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExampleERC721.Contract.SupportsInterface(&_ExampleERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC721 *ExampleERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC721 *ExampleERC721Session) Symbol() (string, error) {
	return _ExampleERC721.Contract.Symbol(&_ExampleERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC721 *ExampleERC721CallerSession) Symbol() (string, error) {
	return _ExampleERC721.Contract.Symbol(&_ExampleERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ExampleERC721 *ExampleERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ExampleERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ExampleERC721 *ExampleERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ExampleERC721.Contract.TokenURI(&_ExampleERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ExampleERC721 *ExampleERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ExampleERC721.Contract.TokenURI(&_ExampleERC721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Approve(&_ExampleERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Approve(&_ExampleERC721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Burn(&_ExampleERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Burn(&_ExampleERC721.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Transactor) Mint(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "mint", tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Session) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Mint(&_ExampleERC721.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.Mint(&_ExampleERC721.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SafeTransferFrom(&_ExampleERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SafeTransferFrom(&_ExampleERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ExampleERC721 *ExampleERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ExampleERC721 *ExampleERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SafeTransferFrom0(&_ExampleERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SafeTransferFrom0(&_ExampleERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ExampleERC721 *ExampleERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ExampleERC721 *ExampleERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SetApprovalForAll(&_ExampleERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ExampleERC721.Contract.SetApprovalForAll(&_ExampleERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.TransferFrom(&_ExampleERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ExampleERC721 *ExampleERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleERC721.Contract.TransferFrom(&_ExampleERC721.TransactOpts, from, to, tokenId)
}

// ExampleERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExampleERC721 contract.
type ExampleERC721ApprovalIterator struct {
	Event *ExampleERC721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleERC721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC721Approval represents a Approval event raised by the ExampleERC721 contract.
type ExampleERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ExampleERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExampleERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721ApprovalIterator{contract: _ExampleERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExampleERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExampleERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC721Approval)
				if err := _ExampleERC721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) ParseApproval(log types.Log) (*ExampleERC721Approval, error) {
	event := new(ExampleERC721Approval)
	if err := _ExampleERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ExampleERC721 contract.
type ExampleERC721ApprovalForAllIterator struct {
	Event *ExampleERC721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleERC721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC721ApprovalForAll represents a ApprovalForAll event raised by the ExampleERC721 contract.
type ExampleERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ExampleERC721 *ExampleERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ExampleERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ExampleERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721ApprovalForAllIterator{contract: _ExampleERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ExampleERC721 *ExampleERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ExampleERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ExampleERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC721ApprovalForAll)
				if err := _ExampleERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ExampleERC721 *ExampleERC721Filterer) ParseApprovalForAll(log types.Log) (*ExampleERC721ApprovalForAll, error) {
	event := new(ExampleERC721ApprovalForAll)
	if err := _ExampleERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleERC721 contract.
type ExampleERC721TransferIterator struct {
	Event *ExampleERC721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleERC721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC721Transfer represents a Transfer event raised by the ExampleERC721 contract.
type ExampleERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ExampleERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExampleERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC721TransferIterator{contract: _ExampleERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExampleERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC721Transfer)
				if err := _ExampleERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ExampleERC721 *ExampleERC721Filterer) ParseTransfer(log types.Log) (*ExampleERC721Transfer, error) {
	event := new(ExampleERC721Transfer)
	if err := _ExampleERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
