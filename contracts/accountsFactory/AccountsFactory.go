// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountsFactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AccountsFactoryMetaData contains all meta data concerning the AccountsFactory contract.
var AccountsFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_opcode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"box721Address\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callAccounts\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"datas\",\"type\":\"bytes\"}],\"name\":\"callContractCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_opcode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b55806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c806330f020b214620000575780635fb998de146200008d578063b0ce95fe14620000c3578063f932b0f414620000f9575b600080fd5b6200007560048036038101906200006f9190620005a1565b6200012f565b6040516200008491906200064c565b60405180910390f35b620000ab6004803603810190620000a591906200074e565b62000158565b604051620000ba919062000847565b60405180910390f35b620000e16004803603810190620000db91906200074e565b620001dc565b604051620000f0919062000847565b60405180910390f35b620001176004803603810190620001119190620005a1565b6200029f565b6040516200012691906200064c565b60405180910390f35b600080600062000140858562000301565b91509150808251602084016000f59250505092915050565b6060630371066860e01b8383604051602401620001779291906200086b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905092915050565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051620002089190620008e1565b6000604051808303816000865af19150503d806000811462000247576040519150601f19603f3d011682016040523d82523d6000602084013e6200024c565b606091505b50915091508162000294576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200028b906200095b565b60405180910390fd5b809250505092915050565b6000806000620002b0858562000301565b91509150600060ff60f81b30838580519060200120604051602001620002da949392919062000a4d565b6040516020818303038152906040528051906020012060001c905080935050505092915050565b60606000604051806020016200031790620003e2565b6020820181038252601f19601f82011660405250836040516020016200033e919062000aef565b6040516020818303038152906040526040516020016200036092919062000b13565b6040516020818303038152906040529150600084846040516020016200038892919062000ba2565b6040516020818303038152906040528051906020012090508485620003ae919062000bfd565b81604051602001620003c292919062000c5e565b604051602081830303815290604052805190602001209150509250929050565b610e918062000c8f83390190565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620004198162000404565b81146200042557600080fd5b50565b60008135905062000439816200040e565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620004948262000449565b810181811067ffffffffffffffff82111715620004b657620004b56200045a565b5b80604052505050565b6000620004cb620003f0565b9050620004d9828262000489565b919050565b600067ffffffffffffffff821115620004fc57620004fb6200045a565b5b620005078262000449565b9050602081019050919050565b82818337600083830152505050565b60006200053a6200053484620004de565b620004bf565b90508281526020810184848401111562000559576200055862000444565b5b6200056684828562000514565b509392505050565b600082601f8301126200058657620005856200043f565b5b81356200059884826020860162000523565b91505092915050565b60008060408385031215620005bb57620005ba620003fa565b5b6000620005cb8582860162000428565b925050602083013567ffffffffffffffff811115620005ef57620005ee620003ff565b5b620005fd858286016200056e565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006348262000607565b9050919050565b620006468162000627565b82525050565b60006020820190506200066360008301846200063b565b92915050565b620006748162000627565b81146200068057600080fd5b50565b600081359050620006948162000669565b92915050565b600067ffffffffffffffff821115620006b857620006b76200045a565b5b620006c38262000449565b9050602081019050919050565b6000620006e7620006e1846200069a565b620004bf565b90508281526020810184848401111562000706576200070562000444565b5b6200071384828562000514565b509392505050565b600082601f8301126200073357620007326200043f565b5b813562000745848260208601620006d0565b91505092915050565b60008060408385031215620007685762000767620003fa565b5b6000620007788582860162000683565b925050602083013567ffffffffffffffff8111156200079c576200079b620003ff565b5b620007aa858286016200071b565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015620007f0578082015181840152602081019050620007d3565b8381111562000800576000848401525b50505050565b60006200081382620007b4565b6200081f8185620007bf565b935062000831818560208601620007d0565b6200083c8162000449565b840191505092915050565b6000602082019050818103600083015262000863818462000806565b905092915050565b60006040820190506200088260008301856200063b565b818103602083015262000896818462000806565b90509392505050565b600081905092915050565b6000620008b782620007b4565b620008c381856200089f565b9350620008d5818560208601620007d0565b80840191505092915050565b6000620008ef8284620008aa565b915081905092915050565b600082825260208201905092915050565b7f4d6574686f642063616c6c206661696c65640000000000000000000000000000600082015250565b600062000943601283620008fa565b915062000950826200090b565b602082019050919050565b60006020820190508181036000830152620009768162000934565b9050919050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b620009c8620009c2826200097d565b620009a9565b82525050565b60008160601b9050919050565b6000620009e882620009ce565b9050919050565b6000620009fc82620009db565b9050919050565b62000a1862000a128262000627565b620009ef565b82525050565b6000819050919050565b6000819050919050565b62000a4762000a418262000a1e565b62000a28565b82525050565b600062000a5b8287620009b3565b60018201915062000a6d828662000a03565b60148201915062000a7f828562000a32565b60208201915062000a91828462000a32565b60208201915081905095945050505050565b600081519050919050565b600062000abb8262000aa3565b62000ac78185620008fa565b935062000ad9818560208601620007d0565b62000ae48162000449565b840191505092915050565b6000602082019050818103600083015262000b0b818462000aae565b905092915050565b600062000b218285620008aa565b915062000b2f8284620008aa565b91508190509392505050565b6000819050919050565b62000b5a62000b548262000404565b62000b3b565b82525050565b600081905092915050565b600062000b788262000aa3565b62000b84818562000b60565b935062000b96818560208601620007d0565b80840191505092915050565b600062000bb0828562000b45565b60208201915062000bc2828462000b6b565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000c0a8262000404565b915062000c178362000404565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000c535762000c5262000bce565b5b828202905092915050565b600062000c6c828562000b45565b60208201915062000c7e828462000a32565b602082019150819050939250505056fe60806040523480156200001157600080fd5b5060405162000e9138038062000e91833981810160405281019062000037919062000390565b620000576200004b6200007760201b60201c565b6200007f60201b60201c565b80600190805190602001906200006f92919062000143565b505062000446565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b828054620001519062000410565b90600052602060002090601f016020900481019282620001755760008555620001c1565b82601f106200019057805160ff1916838001178555620001c1565b82800160010185558215620001c1579182015b82811115620001c0578251825591602001919060010190620001a3565b5b509050620001d09190620001d4565b5090565b5b80821115620001ef576000816000905550600101620001d5565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200025c8262000211565b810181811067ffffffffffffffff821117156200027e576200027d62000222565b5b80604052505050565b600062000293620001f3565b9050620002a1828262000251565b919050565b600067ffffffffffffffff821115620002c457620002c362000222565b5b620002cf8262000211565b9050602081019050919050565b60005b83811015620002fc578082015181840152602081019050620002df565b838111156200030c576000848401525b50505050565b6000620003296200032384620002a6565b62000287565b9050828152602081018484840111156200034857620003476200020c565b5b62000355848285620002dc565b509392505050565b600082601f83011262000375576200037462000207565b5b81516200038784826020860162000312565b91505092915050565b600060208284031215620003a957620003a8620001fd565b5b600082015167ffffffffffffffff811115620003ca57620003c962000202565b5b620003d8848285016200035d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200042957607f821691505b6020821081141562000440576200043f620003e1565b5b50919050565b610a3b80620004566000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80638228863b116100505780638228863b146100a65780638da5cb5b146100c4578063f2fde38b146100e257610067565b8063037106681461006c578063715018a61461009c575b600080fd5b61008660048036038101906100819190610613565b6100fe565b60405161009391906106f7565b60405180910390f35b6100a46101c2565b005b6100ae6101d6565b6040516100bb919061076e565b60405180910390f35b6100cc610264565b6040516100d9919061079f565b60405180910390f35b6100fc60048036038101906100f791906107ba565b61028d565b005b6060610108610311565b6000808473ffffffffffffffffffffffffffffffffffffffff16846040516101309190610823565b6000604051808303816000865af19150503d806000811461016d576040519150601f19603f3d011682016040523d82523d6000602084013e610172565b606091505b5091509150816101b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ae90610886565b60405180910390fd5b809250505092915050565b6101ca610311565b6101d4600061038f565b565b600180546101e3906108d5565b80601f016020809104026020016040519081016040528092919081815260200182805461020f906108d5565b801561025c5780601f106102315761010080835404028352916020019161025c565b820191906000526020600020905b81548152906001019060200180831161023f57829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610295610311565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610305576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fc90610979565b60405180910390fd5b61030e8161038f565b50565b610319610453565b73ffffffffffffffffffffffffffffffffffffffff16610337610264565b73ffffffffffffffffffffffffffffffffffffffff161461038d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610384906109e5565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061049a8261046f565b9050919050565b6104aa8161048f565b81146104b557600080fd5b50565b6000813590506104c7816104a1565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610520826104d7565b810181811067ffffffffffffffff8211171561053f5761053e6104e8565b5b80604052505050565b600061055261045b565b905061055e8282610517565b919050565b600067ffffffffffffffff82111561057e5761057d6104e8565b5b610587826104d7565b9050602081019050919050565b82818337600083830152505050565b60006105b66105b184610563565b610548565b9050828152602081018484840111156105d2576105d16104d2565b5b6105dd848285610594565b509392505050565b600082601f8301126105fa576105f96104cd565b5b813561060a8482602086016105a3565b91505092915050565b6000806040838503121561062a57610629610465565b5b6000610638858286016104b8565b925050602083013567ffffffffffffffff8111156106595761065861046a565b5b610665858286016105e5565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106a957808201518184015260208101905061068e565b838111156106b8576000848401525b50505050565b60006106c98261066f565b6106d3818561067a565b93506106e381856020860161068b565b6106ec816104d7565b840191505092915050565b6000602082019050818103600083015261071181846106be565b905092915050565b600081519050919050565b600082825260208201905092915050565b600061074082610719565b61074a8185610724565b935061075a81856020860161068b565b610763816104d7565b840191505092915050565b600060208201905081810360008301526107888184610735565b905092915050565b6107998161048f565b82525050565b60006020820190506107b46000830184610790565b92915050565b6000602082840312156107d0576107cf610465565b5b60006107de848285016104b8565b91505092915050565b600081905092915050565b60006107fd8261066f565b61080781856107e7565b935061081781856020860161068b565b80840191505092915050565b600061082f82846107f2565b915081905092915050565b7f4d6574686f642063616c6c206661696c65640000000000000000000000000000600082015250565b6000610870601283610724565b915061087b8261083a565b602082019050919050565b6000602082019050818103600083015261089f81610863565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806108ed57607f821691505b60208210811415610901576109006108a6565b5b50919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610963602683610724565b915061096e82610907565b604082019050919050565b6000602082019050818103600083015261099281610956565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006109cf602083610724565b91506109da82610999565b602082019050919050565b600060208201905081810360008301526109fe816109c2565b905091905056fea2646970667358221220ee78770a99c093202e7bbef7fc66646d8eee260264e0796e7756e2cd0c35e64d64736f6c63430008090033a26469706673582212207299beb12e56b7b513abb3e82ea3c39eab2c3c43ff7d1375d43e921525caced664736f6c63430008090033",
}

// AccountsFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountsFactoryMetaData.ABI instead.
var AccountsFactoryABI = AccountsFactoryMetaData.ABI

// AccountsFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountsFactoryMetaData.Bin instead.
var AccountsFactoryBin = AccountsFactoryMetaData.Bin

// DeployAccountsFactory deploys a new Ethereum contract, binding an instance of AccountsFactory to it.
func DeployAccountsFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountsFactory, error) {
	parsed, err := AccountsFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountsFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountsFactory{AccountsFactoryCaller: AccountsFactoryCaller{contract: contract}, AccountsFactoryTransactor: AccountsFactoryTransactor{contract: contract}, AccountsFactoryFilterer: AccountsFactoryFilterer{contract: contract}}, nil
}

// AccountsFactory is an auto generated Go binding around an Ethereum contract.
type AccountsFactory struct {
	AccountsFactoryCaller     // Read-only binding to the contract
	AccountsFactoryTransactor // Write-only binding to the contract
	AccountsFactoryFilterer   // Log filterer for contract events
}

// AccountsFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsFactorySession struct {
	Contract     *AccountsFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsFactoryCallerSession struct {
	Contract *AccountsFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AccountsFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsFactoryTransactorSession struct {
	Contract     *AccountsFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AccountsFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsFactoryRaw struct {
	Contract *AccountsFactory // Generic contract binding to access the raw methods on
}

// AccountsFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsFactoryCallerRaw struct {
	Contract *AccountsFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsFactoryTransactorRaw struct {
	Contract *AccountsFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountsFactory creates a new instance of AccountsFactory, bound to a specific deployed contract.
func NewAccountsFactory(address common.Address, backend bind.ContractBackend) (*AccountsFactory, error) {
	contract, err := bindAccountsFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountsFactory{AccountsFactoryCaller: AccountsFactoryCaller{contract: contract}, AccountsFactoryTransactor: AccountsFactoryTransactor{contract: contract}, AccountsFactoryFilterer: AccountsFactoryFilterer{contract: contract}}, nil
}

// NewAccountsFactoryCaller creates a new read-only instance of AccountsFactory, bound to a specific deployed contract.
func NewAccountsFactoryCaller(address common.Address, caller bind.ContractCaller) (*AccountsFactoryCaller, error) {
	contract, err := bindAccountsFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsFactoryCaller{contract: contract}, nil
}

// NewAccountsFactoryTransactor creates a new write-only instance of AccountsFactory, bound to a specific deployed contract.
func NewAccountsFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsFactoryTransactor, error) {
	contract, err := bindAccountsFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsFactoryTransactor{contract: contract}, nil
}

// NewAccountsFactoryFilterer creates a new log filterer instance of AccountsFactory, bound to a specific deployed contract.
func NewAccountsFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFactoryFilterer, error) {
	contract, err := bindAccountsFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFactoryFilterer{contract: contract}, nil
}

// bindAccountsFactory binds a generic wrapper to an already deployed contract.
func bindAccountsFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountsFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountsFactory *AccountsFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountsFactory.Contract.AccountsFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountsFactory *AccountsFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountsFactory.Contract.AccountsFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountsFactory *AccountsFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountsFactory.Contract.AccountsFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountsFactory *AccountsFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountsFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountsFactory *AccountsFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountsFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountsFactory *AccountsFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountsFactory.Contract.contract.Transact(opts, method, params...)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xf932b0f4.
//
// Solidity: function calculateAddress(uint256 _opcode, string _name) view returns(address)
func (_AccountsFactory *AccountsFactoryCaller) CalculateAddress(opts *bind.CallOpts, _opcode *big.Int, _name string) (common.Address, error) {
	var out []interface{}
	err := _AccountsFactory.contract.Call(opts, &out, "calculateAddress", _opcode, _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateAddress is a free data retrieval call binding the contract method 0xf932b0f4.
//
// Solidity: function calculateAddress(uint256 _opcode, string _name) view returns(address)
func (_AccountsFactory *AccountsFactorySession) CalculateAddress(_opcode *big.Int, _name string) (common.Address, error) {
	return _AccountsFactory.Contract.CalculateAddress(&_AccountsFactory.CallOpts, _opcode, _name)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xf932b0f4.
//
// Solidity: function calculateAddress(uint256 _opcode, string _name) view returns(address)
func (_AccountsFactory *AccountsFactoryCallerSession) CalculateAddress(_opcode *big.Int, _name string) (common.Address, error) {
	return _AccountsFactory.Contract.CalculateAddress(&_AccountsFactory.CallOpts, _opcode, _name)
}

// CallContractCall is a free data retrieval call binding the contract method 0x5fb998de.
//
// Solidity: function callContractCall(address contractAddress, bytes datas) pure returns(bytes data)
func (_AccountsFactory *AccountsFactoryCaller) CallContractCall(opts *bind.CallOpts, contractAddress common.Address, datas []byte) ([]byte, error) {
	var out []interface{}
	err := _AccountsFactory.contract.Call(opts, &out, "callContractCall", contractAddress, datas)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CallContractCall is a free data retrieval call binding the contract method 0x5fb998de.
//
// Solidity: function callContractCall(address contractAddress, bytes datas) pure returns(bytes data)
func (_AccountsFactory *AccountsFactorySession) CallContractCall(contractAddress common.Address, datas []byte) ([]byte, error) {
	return _AccountsFactory.Contract.CallContractCall(&_AccountsFactory.CallOpts, contractAddress, datas)
}

// CallContractCall is a free data retrieval call binding the contract method 0x5fb998de.
//
// Solidity: function callContractCall(address contractAddress, bytes datas) pure returns(bytes data)
func (_AccountsFactory *AccountsFactoryCallerSession) CallContractCall(contractAddress common.Address, datas []byte) ([]byte, error) {
	return _AccountsFactory.Contract.CallContractCall(&_AccountsFactory.CallOpts, contractAddress, datas)
}

// CallAccounts is a paid mutator transaction binding the contract method 0xb0ce95fe.
//
// Solidity: function callAccounts(address box721Address, bytes data) returns(bytes)
func (_AccountsFactory *AccountsFactoryTransactor) CallAccounts(opts *bind.TransactOpts, box721Address common.Address, data []byte) (*types.Transaction, error) {
	return _AccountsFactory.contract.Transact(opts, "callAccounts", box721Address, data)
}

// CallAccounts is a paid mutator transaction binding the contract method 0xb0ce95fe.
//
// Solidity: function callAccounts(address box721Address, bytes data) returns(bytes)
func (_AccountsFactory *AccountsFactorySession) CallAccounts(box721Address common.Address, data []byte) (*types.Transaction, error) {
	return _AccountsFactory.Contract.CallAccounts(&_AccountsFactory.TransactOpts, box721Address, data)
}

// CallAccounts is a paid mutator transaction binding the contract method 0xb0ce95fe.
//
// Solidity: function callAccounts(address box721Address, bytes data) returns(bytes)
func (_AccountsFactory *AccountsFactoryTransactorSession) CallAccounts(box721Address common.Address, data []byte) (*types.Transaction, error) {
	return _AccountsFactory.Contract.CallAccounts(&_AccountsFactory.TransactOpts, box721Address, data)
}

// CreatePair is a paid mutator transaction binding the contract method 0x30f020b2.
//
// Solidity: function createPair(uint256 _opcode, string _name) returns(address pair)
func (_AccountsFactory *AccountsFactoryTransactor) CreatePair(opts *bind.TransactOpts, _opcode *big.Int, _name string) (*types.Transaction, error) {
	return _AccountsFactory.contract.Transact(opts, "createPair", _opcode, _name)
}

// CreatePair is a paid mutator transaction binding the contract method 0x30f020b2.
//
// Solidity: function createPair(uint256 _opcode, string _name) returns(address pair)
func (_AccountsFactory *AccountsFactorySession) CreatePair(_opcode *big.Int, _name string) (*types.Transaction, error) {
	return _AccountsFactory.Contract.CreatePair(&_AccountsFactory.TransactOpts, _opcode, _name)
}

// CreatePair is a paid mutator transaction binding the contract method 0x30f020b2.
//
// Solidity: function createPair(uint256 _opcode, string _name) returns(address pair)
func (_AccountsFactory *AccountsFactoryTransactorSession) CreatePair(_opcode *big.Int, _name string) (*types.Transaction, error) {
	return _AccountsFactory.Contract.CreatePair(&_AccountsFactory.TransactOpts, _opcode, _name)
}