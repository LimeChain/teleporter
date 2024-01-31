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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816000908162000063919062000120565b50600162000072828262000120565b505050620001ec565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000a657607f821691505b602082108103620000c757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200011b57600081815260208120601f850160051c81016020861015620000f65750805b601f850160051c820191505b81811015620001175782815560010162000102565b5050505b505050565b81516001600160401b038111156200013c576200013c6200007b565b62000154816200014d845462000091565b84620000cd565b602080601f8311600181146200018c5760008415620001735750858301515b600019600386901b1c1916600185901b17855562000117565b600085815260208120601f198616915b82811015620001bd578886015182559484019460019091019084016200019c565b5085821015620001dc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6114eb80620001fc6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636352211e11610097578063a22cb46511610066578063a22cb465146101ff578063b88d4fde14610212578063c87b56dd14610225578063e985e9c51461023857600080fd5b80636352211e146101b057806370a08231146101c357806395d89b41146101e4578063a0712d68146101ec57600080fd5b8063095ea7b3116100d3578063095ea7b31461016257806323b872dd1461017757806342842e0e1461018a57806342966c681461019d57600080fd5b806301ffc9a7146100fa57806306fdde0314610122578063081812fc14610137575b600080fd5b61010d61010836600461103b565b61024b565b60405190151581526020015b60405180910390f35b61012a61029d565b60405161011991906110a8565b61014a6101453660046110bb565b61032f565b6040516001600160a01b039091168152602001610119565b6101756101703660046110f0565b610356565b005b61017561018536600461111a565b610470565b61017561019836600461111a565b6104a2565b6101756101ab3660046110bb565b6104bd565b61014a6101be3660046110bb565b6104ee565b6101d66101d1366004611156565b61054e565b604051908152602001610119565b61012a6105d4565b6101756101fa3660046110bb565b6105e3565b61017561020d366004611171565b610649565b6101756102203660046111c3565b610658565b61012a6102333660046110bb565b610690565b61010d61024636600461129f565b610729565b60006001600160e01b031982166380ac58cd60e01b148061027c57506001600160e01b03198216635b5e139f60e01b145b8061029757506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546102ac906112d2565b80601f01602080910402602001604051908101604052809291908181526020018280546102d8906112d2565b80156103255780601f106102fa57610100808354040283529160200191610325565b820191906000526020600020905b81548152906001019060200180831161030857829003601f168201915b5050505050905090565b600061033a82610757565b506000908152600460205260409020546001600160a01b031690565b6000610361826104ee565b9050806001600160a01b0316836001600160a01b0316036103d35760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806103ef57506103ef8133610729565b6104615760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c00000060648201526084016103ca565b61046b83836107b6565b505050565b61047b335b82610824565b6104975760405162461bcd60e51b81526004016103ca9061130c565b61046b838383610883565b61046b83838360405180602001604052806000815250610658565b6104c633610475565b6104e25760405162461bcd60e51b81526004016103ca9061130c565b6104eb816109f4565b50565b6000818152600260205260408120546001600160a01b0316806102975760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016103ca565b60006001600160a01b0382166105b85760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016103ca565b506001600160a01b031660009081526003602052604090205490565b6060600180546102ac906112d2565b6000818152600260205260409020546001600160a01b03161561063f5760405162461bcd60e51b8152602060048201526014602482015273151bdad95b88185b1c9958591e481b5a5b9d195960621b60448201526064016103ca565b6104eb3382610a97565b610654338383610c30565b5050565b6106623383610824565b61067e5760405162461bcd60e51b81526004016103ca9061130c565b61068a84848484610cfe565b50505050565b606061069b82610757565b60006106d760408051808201909152601981527f68747470733a2f2f6578616d706c652e636f6d2f697066732f00000000000000602082015290565b905060008151116106f75760405180602001604052806000815250610722565b8061070184610d31565b604051602001610712929190611359565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000818152600260205260409020546001600160a01b03166104eb5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016103ca565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906107eb826104ee565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610830836104ee565b9050806001600160a01b0316846001600160a01b0316148061085757506108578185610729565b8061087b5750836001600160a01b03166108708461032f565b6001600160a01b0316145b949350505050565b826001600160a01b0316610896826104ee565b6001600160a01b0316146108bc5760405162461bcd60e51b81526004016103ca90611388565b6001600160a01b03821661091e5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016103ca565b61092b8383836001610dc4565b826001600160a01b031661093e826104ee565b6001600160a01b0316146109645760405162461bcd60e51b81526004016103ca90611388565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60006109ff826104ee565b9050610a0f816000846001610dc4565b610a18826104ee565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6001600160a01b038216610aed5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016103ca565b6000818152600260205260409020546001600160a01b031615610b525760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016103ca565b610b60600083836001610dc4565b6000818152600260205260409020546001600160a01b031615610bc55760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016103ca565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b031603610c915760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016103ca565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610d09848484610883565b610d1584848484610e4c565b61068a5760405162461bcd60e51b81526004016103ca906113cd565b60606000610d3e83610f4d565b600101905060008167ffffffffffffffff811115610d5e57610d5e6111ad565b6040519080825280601f01601f191660200182016040528015610d88576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610d9257509392505050565b600181111561068a576001600160a01b03841615610e0a576001600160a01b03841660009081526003602052604081208054839290610e04908490611435565b90915550505b6001600160a01b0383161561068a576001600160a01b03831660009081526003602052604081208054839290610e41908490611448565b909155505050505050565b60006001600160a01b0384163b15610f4257604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610e9090339089908890889060040161145b565b6020604051808303816000875af1925050508015610ecb575060408051601f3d908101601f19168201909252610ec891810190611498565b60015b610f28573d808015610ef9576040519150601f19603f3d011682016040523d82523d6000602084013e610efe565b606091505b508051600003610f205760405162461bcd60e51b81526004016103ca906113cd565b805181602001fd5b6001600160e01b031916630a85bd0160e11b14905061087b565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310610f8c5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310610fb8576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610fd657662386f26fc10000830492506010015b6305f5e1008310610fee576305f5e100830492506008015b612710831061100257612710830492506004015b60648310611014576064830492506002015b600a83106102975760010192915050565b6001600160e01b0319811681146104eb57600080fd5b60006020828403121561104d57600080fd5b813561072281611025565b60005b8381101561107357818101518382015260200161105b565b50506000910152565b60008151808452611094816020860160208601611058565b601f01601f19169290920160200192915050565b602081526000610722602083018461107c565b6000602082840312156110cd57600080fd5b5035919050565b80356001600160a01b03811681146110eb57600080fd5b919050565b6000806040838503121561110357600080fd5b61110c836110d4565b946020939093013593505050565b60008060006060848603121561112f57600080fd5b611138846110d4565b9250611146602085016110d4565b9150604084013590509250925092565b60006020828403121561116857600080fd5b610722826110d4565b6000806040838503121561118457600080fd5b61118d836110d4565b9150602083013580151581146111a257600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156111d957600080fd5b6111e2856110d4565b93506111f0602086016110d4565b925060408501359150606085013567ffffffffffffffff8082111561121457600080fd5b818701915087601f83011261122857600080fd5b81358181111561123a5761123a6111ad565b604051601f8201601f19908116603f01168101908382118183101715611262576112626111ad565b816040528281528a602084870101111561127b57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b600080604083850312156112b257600080fd5b6112bb836110d4565b91506112c9602084016110d4565b90509250929050565b600181811c908216806112e657607f821691505b60208210810361130657634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6000835161136b818460208801611058565b83519083019061137f818360208801611058565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b818103818111156102975761029761141f565b808201808211156102975761029761141f565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061148e9083018461107c565b9695505050505050565b6000602082840312156114aa57600080fd5b81516107228161102556fea2646970667358221220c00750aafd07068c48ca478ddcb7cd40cd77873ba5b292c87c83a581e1073ba964736f6c63430008120033",
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
