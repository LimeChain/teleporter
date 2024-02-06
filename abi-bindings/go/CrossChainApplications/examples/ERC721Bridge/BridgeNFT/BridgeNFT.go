// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgenft

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

// BridgeNFTMetaData contains all meta data concerning the BridgeNFT contract.
var BridgeNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162001a6638038062001a66833981016040819052620000359162000170565b82826000620000458382620002bf565b506001620000548282620002bf565b5050336080525060a08690526001600160a01b0385811660c052841660e0526006620000818282620002bf565b505050505050506200038b565b80516001600160a01b0381168114620000a657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000d357600080fd5b81516001600160401b0380821115620000f057620000f0620000ab565b604051601f8301601f19908116603f011681019082821181831017156200011b576200011b620000ab565b816040528381526020925086838588010111156200013857600080fd5b600091505b838210156200015c57858201830151818301840152908201906200013d565b600093810190920192909252949350505050565b60008060008060008060c087890312156200018a57600080fd5b865195506200019c602088016200008e565b9450620001ac604088016200008e565b60608801519094506001600160401b0380821115620001ca57600080fd5b620001d88a838b01620000c1565b94506080890151915080821115620001ef57600080fd5b620001fd8a838b01620000c1565b935060a08901519150808211156200021457600080fd5b506200022389828a01620000c1565b9150509295509295509295565b600181811c908216806200024557607f821691505b6020821081036200026657634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002ba57600081815260208120601f850160051c81016020861015620002955750805b601f850160051c820191505b81811015620002b657828155600101620002a1565b5050505b505050565b81516001600160401b03811115620002db57620002db620000ab565b620002f381620002ec845462000230565b846200026c565b602080601f8311600181146200032b5760008415620003125750858301515b600019600386901b1c1916600185901b178555620002b6565b600085815260208120601f198616915b828110156200035c578886015182559484019460019091019084016200033b565b50858210156200037b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051611693620003d36000396000610262015260006101b3015260006103040152600081816102ca0152818161061501526106b001526116936000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80636352211e116100ad578063b88d4fde11610071578063b88d4fde1461029f578063c87b56dd146102b2578063cd596583146102c5578063e985e9c5146102ec578063f7253968146102ff57600080fd5b80636352211e1461022957806370a082311461023c57806374d32ad41461025d57806395d89b4114610284578063a22cb4651461028c57600080fd5b806323b872dd116100f457806323b872dd146101d55780632d5c0446146101e857806340c10f19146101f057806342842e0e1461020357806342966c681461021657600080fd5b806301ffc9a71461013157806306fdde0314610159578063081812fc1461016e578063095ea7b3146101995780631a0b79bf146101ae575b600080fd5b61014461013f3660046111e3565b610326565b60405190151581526020015b60405180910390f35b610161610378565b6040516101509190611250565b61018161017c366004611263565b61040a565b6040516001600160a01b039091168152602001610150565b6101ac6101a7366004611298565b610431565b005b6101817f000000000000000000000000000000000000000000000000000000000000000081565b6101ac6101e33660046112c2565b61054b565b61016161057c565b6101ac6101fe366004611298565b61060a565b6101ac6102113660046112c2565b61068a565b6101ac610224366004611263565b6106a5565b610181610237366004611263565b610723565b61024f61024a3660046112fe565b610783565b604051908152602001610150565b6101817f000000000000000000000000000000000000000000000000000000000000000081565b610161610809565b6101ac61029a366004611319565b610818565b6101ac6102ad36600461136b565b610823565b6101616102c0366004611263565b61085b565b6101817f000000000000000000000000000000000000000000000000000000000000000081565b6101446102fa366004611447565b6108c2565b61024f7f000000000000000000000000000000000000000000000000000000000000000081565b60006001600160e01b031982166380ac58cd60e01b148061035757506001600160e01b03198216635b5e139f60e01b145b8061037257506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546103879061147a565b80601f01602080910402602001604051908101604052809291908181526020018280546103b39061147a565b80156104005780601f106103d557610100808354040283529160200191610400565b820191906000526020600020905b8154815290600101906020018083116103e357829003601f168201915b5050505050905090565b6000610415826108f0565b506000908152600460205260409020546001600160a01b031690565b600061043c82610723565b9050806001600160a01b0316836001600160a01b0316036104ae5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104ca57506104ca81336108c2565b61053c5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c00000060648201526084016104a5565b610546838361094f565b505050565b61055533826109bd565b6105715760405162461bcd60e51b81526004016104a5906114b4565b610546838383610a1c565b600680546105899061147a565b80601f01602080910402602001604051908101604052809291908181526020018280546105b59061147a565b80156106025780601f106105d757610100808354040283529160200191610602565b820191906000526020600020905b8154815290600101906020018083116105e557829003601f168201915b505050505081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461067c5760405162461bcd60e51b8152602060048201526017602482015276109c9a5919d95391950e881d5b985d5d1a1bdc9a5e9959604a1b60448201526064016104a5565b6106868282610b8d565b5050565b61054683838360405180602001604052806000815250610823565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107175760405162461bcd60e51b8152602060048201526017602482015276109c9a5919d95391950e881d5b985d5d1a1bdc9a5e9959604a1b60448201526064016104a5565b61072081610d26565b50565b6000818152600260205260408120546001600160a01b0316806103725760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104a5565b60006001600160a01b0382166107ed5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016104a5565b506001600160a01b031660009081526003602052604090205490565b6060600180546103879061147a565b610686338383610dc9565b61082d33836109bd565b6108495760405162461bcd60e51b81526004016104a5906114b4565b61085584848484610e97565b50505050565b6060610866826108f0565b6000610870610eca565b9050600081511161089057604051806020016040528060008152506108bb565b8061089a84610ed9565b6040516020016108ab929190611501565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000818152600260205260409020546001600160a01b03166107205760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016104a5565b600081815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061098482610723565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806109c983610723565b9050806001600160a01b0316846001600160a01b031614806109f057506109f081856108c2565b80610a145750836001600160a01b0316610a098461040a565b6001600160a01b0316145b949350505050565b826001600160a01b0316610a2f82610723565b6001600160a01b031614610a555760405162461bcd60e51b81526004016104a590611530565b6001600160a01b038216610ab75760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016104a5565b610ac48383836001610f6c565b826001600160a01b0316610ad782610723565b6001600160a01b031614610afd5760405162461bcd60e51b81526004016104a590611530565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216610be35760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016104a5565b6000818152600260205260409020546001600160a01b031615610c485760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104a5565b610c56600083836001610f6c565b6000818152600260205260409020546001600160a01b031615610cbb5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016104a5565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000610d3182610723565b9050610d41816000846001610f6c565b610d4a82610723565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b816001600160a01b0316836001600160a01b031603610e2a5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016104a5565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610ea2848484610a1c565b610eae84848484610ff4565b6108555760405162461bcd60e51b81526004016104a590611575565b6060600680546103879061147a565b60606000610ee6836110f5565b600101905060008167ffffffffffffffff811115610f0657610f06611355565b6040519080825280601f01601f191660200182016040528015610f30576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610f3a57509392505050565b6001811115610855576001600160a01b03841615610fb2576001600160a01b03841660009081526003602052604081208054839290610fac9084906115dd565b90915550505b6001600160a01b03831615610855576001600160a01b03831660009081526003602052604081208054839290610fe99084906115f0565b909155505050505050565b60006001600160a01b0384163b156110ea57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611038903390899088908890600401611603565b6020604051808303816000875af1925050508015611073575060408051601f3d908101601f1916820190925261107091810190611640565b60015b6110d0573d8080156110a1576040519150601f19603f3d011682016040523d82523d6000602084013e6110a6565b606091505b5080516000036110c85760405162461bcd60e51b81526004016104a590611575565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610a14565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106111345772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611160576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061117e57662386f26fc10000830492506010015b6305f5e1008310611196576305f5e100830492506008015b61271083106111aa57612710830492506004015b606483106111bc576064830492506002015b600a83106103725760010192915050565b6001600160e01b03198116811461072057600080fd5b6000602082840312156111f557600080fd5b81356108bb816111cd565b60005b8381101561121b578181015183820152602001611203565b50506000910152565b6000815180845261123c816020860160208601611200565b601f01601f19169290920160200192915050565b6020815260006108bb6020830184611224565b60006020828403121561127557600080fd5b5035919050565b80356001600160a01b038116811461129357600080fd5b919050565b600080604083850312156112ab57600080fd5b6112b48361127c565b946020939093013593505050565b6000806000606084860312156112d757600080fd5b6112e08461127c565b92506112ee6020850161127c565b9150604084013590509250925092565b60006020828403121561131057600080fd5b6108bb8261127c565b6000806040838503121561132c57600080fd5b6113358361127c565b91506020830135801515811461134a57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561138157600080fd5b61138a8561127c565b93506113986020860161127c565b925060408501359150606085013567ffffffffffffffff808211156113bc57600080fd5b818701915087601f8301126113d057600080fd5b8135818111156113e2576113e2611355565b604051601f8201601f19908116603f0116810190838211818310171561140a5761140a611355565b816040528281528a602084870101111561142357600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561145a57600080fd5b6114638361127c565b91506114716020840161127c565b90509250929050565b600181811c9082168061148e57607f821691505b6020821081036114ae57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b60008351611513818460208801611200565b835190830190611527818360208801611200565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610372576103726115c7565b80820180821115610372576103726115c7565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061163690830184611224565b9695505050505050565b60006020828403121561165257600080fd5b81516108bb816111cd56fea2646970667358221220d4fe8a8b4909e01b6c8f8dc82eb5cc4463048e54e0ee713e30cdd346288abdd764736f6c63430008120033",
}

// BridgeNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeNFTMetaData.ABI instead.
var BridgeNFTABI = BridgeNFTMetaData.ABI

// BridgeNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeNFTMetaData.Bin instead.
var BridgeNFTBin = BridgeNFTMetaData.Bin

// DeployBridgeNFT deploys a new Ethereum contract, binding an instance of BridgeNFT to it.
func DeployBridgeNFT(auth *bind.TransactOpts, backend bind.ContractBackend, sourceBlockchainID [32]byte, sourceBridge common.Address, sourceAsset common.Address, tokenName string, tokenSymbol string, tokenURI string) (common.Address, *types.Transaction, *BridgeNFT, error) {
	parsed, err := BridgeNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeNFTBin), backend, sourceBlockchainID, sourceBridge, sourceAsset, tokenName, tokenSymbol, tokenURI)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeNFT{BridgeNFTCaller: BridgeNFTCaller{contract: contract}, BridgeNFTTransactor: BridgeNFTTransactor{contract: contract}, BridgeNFTFilterer: BridgeNFTFilterer{contract: contract}}, nil
}

// BridgeNFT is an auto generated Go binding around an Ethereum contract.
type BridgeNFT struct {
	BridgeNFTCaller     // Read-only binding to the contract
	BridgeNFTTransactor // Write-only binding to the contract
	BridgeNFTFilterer   // Log filterer for contract events
}

// BridgeNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeNFTSession struct {
	Contract     *BridgeNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeNFTCallerSession struct {
	Contract *BridgeNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeNFTTransactorSession struct {
	Contract     *BridgeNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeNFTRaw struct {
	Contract *BridgeNFT // Generic contract binding to access the raw methods on
}

// BridgeNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeNFTCallerRaw struct {
	Contract *BridgeNFTCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeNFTTransactorRaw struct {
	Contract *BridgeNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeNFT creates a new instance of BridgeNFT, bound to a specific deployed contract.
func NewBridgeNFT(address common.Address, backend bind.ContractBackend) (*BridgeNFT, error) {
	contract, err := bindBridgeNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeNFT{BridgeNFTCaller: BridgeNFTCaller{contract: contract}, BridgeNFTTransactor: BridgeNFTTransactor{contract: contract}, BridgeNFTFilterer: BridgeNFTFilterer{contract: contract}}, nil
}

// NewBridgeNFTCaller creates a new read-only instance of BridgeNFT, bound to a specific deployed contract.
func NewBridgeNFTCaller(address common.Address, caller bind.ContractCaller) (*BridgeNFTCaller, error) {
	contract, err := bindBridgeNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTCaller{contract: contract}, nil
}

// NewBridgeNFTTransactor creates a new write-only instance of BridgeNFT, bound to a specific deployed contract.
func NewBridgeNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeNFTTransactor, error) {
	contract, err := bindBridgeNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTTransactor{contract: contract}, nil
}

// NewBridgeNFTFilterer creates a new log filterer instance of BridgeNFT, bound to a specific deployed contract.
func NewBridgeNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeNFTFilterer, error) {
	contract, err := bindBridgeNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTFilterer{contract: contract}, nil
}

// bindBridgeNFT binds a generic wrapper to an already deployed contract.
func bindBridgeNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeNFT *BridgeNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeNFT.Contract.BridgeNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeNFT *BridgeNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeNFT.Contract.BridgeNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeNFT *BridgeNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeNFT.Contract.BridgeNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeNFT *BridgeNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeNFT *BridgeNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeNFT *BridgeNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BridgeNFT *BridgeNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BridgeNFT *BridgeNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BridgeNFT.Contract.BalanceOf(&_BridgeNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BridgeNFT *BridgeNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BridgeNFT.Contract.BalanceOf(&_BridgeNFT.CallOpts, owner)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeNFT *BridgeNFTCaller) BridgeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "bridgeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeNFT *BridgeNFTSession) BridgeContract() (common.Address, error) {
	return _BridgeNFT.Contract.BridgeContract(&_BridgeNFT.CallOpts)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeNFT *BridgeNFTCallerSession) BridgeContract() (common.Address, error) {
	return _BridgeNFT.Contract.BridgeContract(&_BridgeNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BridgeNFT.Contract.GetApproved(&_BridgeNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BridgeNFT.Contract.GetApproved(&_BridgeNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BridgeNFT *BridgeNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BridgeNFT *BridgeNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BridgeNFT.Contract.IsApprovedForAll(&_BridgeNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BridgeNFT *BridgeNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BridgeNFT.Contract.IsApprovedForAll(&_BridgeNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeNFT *BridgeNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeNFT *BridgeNFTSession) Name() (string, error) {
	return _BridgeNFT.Contract.Name(&_BridgeNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeNFT *BridgeNFTCallerSession) Name() (string, error) {
	return _BridgeNFT.Contract.Name(&_BridgeNFT.CallOpts)
}

// NativeAsset is a free data retrieval call binding the contract method 0x74d32ad4.
//
// Solidity: function nativeAsset() view returns(address)
func (_BridgeNFT *BridgeNFTCaller) NativeAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "nativeAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeAsset is a free data retrieval call binding the contract method 0x74d32ad4.
//
// Solidity: function nativeAsset() view returns(address)
func (_BridgeNFT *BridgeNFTSession) NativeAsset() (common.Address, error) {
	return _BridgeNFT.Contract.NativeAsset(&_BridgeNFT.CallOpts)
}

// NativeAsset is a free data retrieval call binding the contract method 0x74d32ad4.
//
// Solidity: function nativeAsset() view returns(address)
func (_BridgeNFT *BridgeNFTCallerSession) NativeAsset() (common.Address, error) {
	return _BridgeNFT.Contract.NativeAsset(&_BridgeNFT.CallOpts)
}

// NativeBlockchainID is a free data retrieval call binding the contract method 0xf7253968.
//
// Solidity: function nativeBlockchainID() view returns(bytes32)
func (_BridgeNFT *BridgeNFTCaller) NativeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "nativeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeBlockchainID is a free data retrieval call binding the contract method 0xf7253968.
//
// Solidity: function nativeBlockchainID() view returns(bytes32)
func (_BridgeNFT *BridgeNFTSession) NativeBlockchainID() ([32]byte, error) {
	return _BridgeNFT.Contract.NativeBlockchainID(&_BridgeNFT.CallOpts)
}

// NativeBlockchainID is a free data retrieval call binding the contract method 0xf7253968.
//
// Solidity: function nativeBlockchainID() view returns(bytes32)
func (_BridgeNFT *BridgeNFTCallerSession) NativeBlockchainID() ([32]byte, error) {
	return _BridgeNFT.Contract.NativeBlockchainID(&_BridgeNFT.CallOpts)
}

// NativeBridge is a free data retrieval call binding the contract method 0x1a0b79bf.
//
// Solidity: function nativeBridge() view returns(address)
func (_BridgeNFT *BridgeNFTCaller) NativeBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "nativeBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeBridge is a free data retrieval call binding the contract method 0x1a0b79bf.
//
// Solidity: function nativeBridge() view returns(address)
func (_BridgeNFT *BridgeNFTSession) NativeBridge() (common.Address, error) {
	return _BridgeNFT.Contract.NativeBridge(&_BridgeNFT.CallOpts)
}

// NativeBridge is a free data retrieval call binding the contract method 0x1a0b79bf.
//
// Solidity: function nativeBridge() view returns(address)
func (_BridgeNFT *BridgeNFTCallerSession) NativeBridge() (common.Address, error) {
	return _BridgeNFT.Contract.NativeBridge(&_BridgeNFT.CallOpts)
}

// NativeTokenURI is a free data retrieval call binding the contract method 0x2d5c0446.
//
// Solidity: function nativeTokenURI() view returns(string)
func (_BridgeNFT *BridgeNFTCaller) NativeTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "nativeTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NativeTokenURI is a free data retrieval call binding the contract method 0x2d5c0446.
//
// Solidity: function nativeTokenURI() view returns(string)
func (_BridgeNFT *BridgeNFTSession) NativeTokenURI() (string, error) {
	return _BridgeNFT.Contract.NativeTokenURI(&_BridgeNFT.CallOpts)
}

// NativeTokenURI is a free data retrieval call binding the contract method 0x2d5c0446.
//
// Solidity: function nativeTokenURI() view returns(string)
func (_BridgeNFT *BridgeNFTCallerSession) NativeTokenURI() (string, error) {
	return _BridgeNFT.Contract.NativeTokenURI(&_BridgeNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BridgeNFT.Contract.OwnerOf(&_BridgeNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BridgeNFT *BridgeNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BridgeNFT.Contract.OwnerOf(&_BridgeNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeNFT *BridgeNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeNFT *BridgeNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeNFT.Contract.SupportsInterface(&_BridgeNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeNFT *BridgeNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeNFT.Contract.SupportsInterface(&_BridgeNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeNFT *BridgeNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeNFT *BridgeNFTSession) Symbol() (string, error) {
	return _BridgeNFT.Contract.Symbol(&_BridgeNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeNFT *BridgeNFTCallerSession) Symbol() (string, error) {
	return _BridgeNFT.Contract.Symbol(&_BridgeNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BridgeNFT *BridgeNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BridgeNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BridgeNFT *BridgeNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BridgeNFT.Contract.TokenURI(&_BridgeNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BridgeNFT *BridgeNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BridgeNFT.Contract.TokenURI(&_BridgeNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Approve(&_BridgeNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Approve(&_BridgeNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Burn(&_BridgeNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Burn(&_BridgeNFT.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactor) Mint(opts *bind.TransactOpts, account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "mint", account, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTSession) Mint(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Mint(&_BridgeNFT.TransactOpts, account, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) Mint(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.Mint(&_BridgeNFT.TransactOpts, account, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SafeTransferFrom(&_BridgeNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SafeTransferFrom(&_BridgeNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BridgeNFT *BridgeNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BridgeNFT *BridgeNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SafeTransferFrom0(&_BridgeNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SafeTransferFrom0(&_BridgeNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BridgeNFT *BridgeNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BridgeNFT *BridgeNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SetApprovalForAll(&_BridgeNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BridgeNFT.Contract.SetApprovalForAll(&_BridgeNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.TransferFrom(&_BridgeNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BridgeNFT *BridgeNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BridgeNFT.Contract.TransferFrom(&_BridgeNFT.TransactOpts, from, to, tokenId)
}

// BridgeNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BridgeNFT contract.
type BridgeNFTApprovalIterator struct {
	Event *BridgeNFTApproval // Event containing the contract specifics and raw log

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
func (it *BridgeNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeNFTApproval)
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
		it.Event = new(BridgeNFTApproval)
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
func (it *BridgeNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeNFTApproval represents a Approval event raised by the BridgeNFT contract.
type BridgeNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BridgeNFT *BridgeNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BridgeNFTApprovalIterator, error) {

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

	logs, sub, err := _BridgeNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTApprovalIterator{contract: _BridgeNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BridgeNFT *BridgeNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BridgeNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeNFTApproval)
				if err := _BridgeNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BridgeNFT *BridgeNFTFilterer) ParseApproval(log types.Log) (*BridgeNFTApproval, error) {
	event := new(BridgeNFTApproval)
	if err := _BridgeNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BridgeNFT contract.
type BridgeNFTApprovalForAllIterator struct {
	Event *BridgeNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BridgeNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeNFTApprovalForAll)
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
		it.Event = new(BridgeNFTApprovalForAll)
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
func (it *BridgeNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeNFTApprovalForAll represents a ApprovalForAll event raised by the BridgeNFT contract.
type BridgeNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BridgeNFT *BridgeNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BridgeNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BridgeNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTApprovalForAllIterator{contract: _BridgeNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BridgeNFT *BridgeNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BridgeNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BridgeNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeNFTApprovalForAll)
				if err := _BridgeNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_BridgeNFT *BridgeNFTFilterer) ParseApprovalForAll(log types.Log) (*BridgeNFTApprovalForAll, error) {
	event := new(BridgeNFTApprovalForAll)
	if err := _BridgeNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BridgeNFT contract.
type BridgeNFTTransferIterator struct {
	Event *BridgeNFTTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeNFTTransfer)
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
		it.Event = new(BridgeNFTTransfer)
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
func (it *BridgeNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeNFTTransfer represents a Transfer event raised by the BridgeNFT contract.
type BridgeNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BridgeNFT *BridgeNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BridgeNFTTransferIterator, error) {

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

	logs, sub, err := _BridgeNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeNFTTransferIterator{contract: _BridgeNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BridgeNFT *BridgeNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BridgeNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeNFTTransfer)
				if err := _BridgeNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BridgeNFT *BridgeNFTFilterer) ParseTransfer(log types.Log) (*BridgeNFTTransfer, error) {
	event := new(BridgeNFTTransfer)
	if err := _BridgeNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
