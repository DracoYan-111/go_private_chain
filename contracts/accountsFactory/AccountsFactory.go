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
	Bin: "0x608060405234801561001057600080fd5b50612099806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c806330f020b214620000575780635fb998de146200008d578063b0ce95fe14620000c3578063f932b0f414620000f9575b600080fd5b6200007560048036038101906200006f9190620005a1565b6200012f565b6040516200008491906200064c565b60405180910390f35b620000ab6004803603810190620000a591906200074e565b62000158565b604051620000ba919062000847565b60405180910390f35b620000e16004803603810190620000db91906200074e565b620001dc565b604051620000f0919062000847565b60405180910390f35b620001176004803603810190620001119190620005a1565b6200029f565b6040516200012691906200064c565b60405180910390f35b600080600062000140858562000301565b91509150808251602084016000f59250505092915050565b6060630371066860e01b8383604051602401620001779291906200086b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905092915050565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051620002089190620008e1565b6000604051808303816000865af19150503d806000811462000247576040519150601f19603f3d011682016040523d82523d6000602084013e6200024c565b606091505b50915091508162000294576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200028b906200095b565b60405180910390fd5b809250505092915050565b6000806000620002b0858562000301565b91509150600060ff60f81b30838580519060200120604051602001620002da949392919062000a4d565b6040516020818303038152906040528051906020012060001c905080935050505092915050565b60606000604051806020016200031790620003e2565b6020820181038252601f19601f82011660405250836040516020016200033e919062000aef565b6040516020818303038152906040526040516020016200036092919062000b13565b6040516020818303038152906040529150600084846040516020016200038892919062000ba2565b6040516020818303038152906040528051906020012090508485620003ae919062000bfd565b81604051602001620003c292919062000c5e565b604051602081830303815290604052805190602001209150509250929050565b6113d58062000c8f83390190565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620004198162000404565b81146200042557600080fd5b50565b60008135905062000439816200040e565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620004948262000449565b810181811067ffffffffffffffff82111715620004b657620004b56200045a565b5b80604052505050565b6000620004cb620003f0565b9050620004d9828262000489565b919050565b600067ffffffffffffffff821115620004fc57620004fb6200045a565b5b620005078262000449565b9050602081019050919050565b82818337600083830152505050565b60006200053a6200053484620004de565b620004bf565b90508281526020810184848401111562000559576200055862000444565b5b6200056684828562000514565b509392505050565b600082601f8301126200058657620005856200043f565b5b81356200059884826020860162000523565b91505092915050565b60008060408385031215620005bb57620005ba620003fa565b5b6000620005cb8582860162000428565b925050602083013567ffffffffffffffff811115620005ef57620005ee620003ff565b5b620005fd858286016200056e565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006348262000607565b9050919050565b620006468162000627565b82525050565b60006020820190506200066360008301846200063b565b92915050565b620006748162000627565b81146200068057600080fd5b50565b600081359050620006948162000669565b92915050565b600067ffffffffffffffff821115620006b857620006b76200045a565b5b620006c38262000449565b9050602081019050919050565b6000620006e7620006e1846200069a565b620004bf565b90508281526020810184848401111562000706576200070562000444565b5b6200071384828562000514565b509392505050565b600082601f8301126200073357620007326200043f565b5b813562000745848260208601620006d0565b91505092915050565b60008060408385031215620007685762000767620003fa565b5b6000620007788582860162000683565b925050602083013567ffffffffffffffff8111156200079c576200079b620003ff565b5b620007aa858286016200071b565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015620007f0578082015181840152602081019050620007d3565b8381111562000800576000848401525b50505050565b60006200081382620007b4565b6200081f8185620007bf565b935062000831818560208601620007d0565b6200083c8162000449565b840191505092915050565b6000602082019050818103600083015262000863818462000806565b905092915050565b60006040820190506200088260008301856200063b565b818103602083015262000896818462000806565b90509392505050565b600081905092915050565b6000620008b782620007b4565b620008c381856200089f565b9350620008d5818560208601620007d0565b80840191505092915050565b6000620008ef8284620008aa565b915081905092915050565b600082825260208201905092915050565b7f4d6574686f642063616c6c206661696c65640000000000000000000000000000600082015250565b600062000943601283620008fa565b915062000950826200090b565b602082019050919050565b60006020820190508181036000830152620009768162000934565b9050919050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b620009c8620009c2826200097d565b620009a9565b82525050565b60008160601b9050919050565b6000620009e882620009ce565b9050919050565b6000620009fc82620009db565b9050919050565b62000a1862000a128262000627565b620009ef565b82525050565b6000819050919050565b6000819050919050565b62000a4762000a418262000a1e565b62000a28565b82525050565b600062000a5b8287620009b3565b60018201915062000a6d828662000a03565b60148201915062000a7f828562000a32565b60208201915062000a91828462000a32565b60208201915081905095945050505050565b600081519050919050565b600062000abb8262000aa3565b62000ac78185620008fa565b935062000ad9818560208601620007d0565b62000ae48162000449565b840191505092915050565b6000602082019050818103600083015262000b0b818462000aae565b905092915050565b600062000b218285620008aa565b915062000b2f8284620008aa565b91508190509392505050565b6000819050919050565b62000b5a62000b548262000404565b62000b3b565b82525050565b600081905092915050565b600062000b788262000aa3565b62000b84818562000b60565b935062000b96818560208601620007d0565b80840191505092915050565b600062000bb0828562000b45565b60208201915062000bc2828462000b6b565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000c0a8262000404565b915062000c178362000404565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000c535762000c5262000bce565b5b828202905092915050565b600062000c6c828562000b45565b60208201915062000c7e828462000a32565b602082019150819050939250505056fe60806040523480156200001157600080fd5b50604051620013d5380380620013d5833981810160405281019062000037919062000390565b620000576200004b6200007760201b60201c565b6200007f60201b60201c565b80600190805190602001906200006f92919062000143565b505062000446565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b828054620001519062000410565b90600052602060002090601f016020900481019282620001755760008555620001c1565b82601f106200019057805160ff1916838001178555620001c1565b82800160010185558215620001c1579182015b82811115620001c0578251825591602001919060010190620001a3565b5b509050620001d09190620001d4565b5090565b5b80821115620001ef576000816000905550600101620001d5565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200025c8262000211565b810181811067ffffffffffffffff821117156200027e576200027d62000222565b5b80604052505050565b600062000293620001f3565b9050620002a1828262000251565b919050565b600067ffffffffffffffff821115620002c457620002c362000222565b5b620002cf8262000211565b9050602081019050919050565b60005b83811015620002fc578082015181840152602081019050620002df565b838111156200030c576000848401525b50505050565b6000620003296200032384620002a6565b62000287565b9050828152602081018484840111156200034857620003476200020c565b5b62000355848285620002dc565b509392505050565b600082601f83011262000375576200037462000207565b5b81516200038784826020860162000312565b91505092915050565b600060208284031215620003a957620003a8620001fd565b5b600082015167ffffffffffffffff811115620003ca57620003c962000202565b5b620003d8848285016200035d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200042957607f821691505b6020821081141562000440576200043f620003e1565b5b50919050565b610f7f80620004566000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638228863b11610076578063bc197c811161005b578063bc197c811461017e578063f23a6e61146101ae578063f2fde38b146101de576100a3565b80638228863b146101425780638da5cb5b14610160576100a3565b806301ffc9a7146100a857806303710668146100d8578063150b7a0214610108578063715018a614610138575b600080fd5b6100c260048036038101906100bd919061063c565b6101fa565b6040516100cf9190610684565b60405180910390f35b6100f260048036038101906100ed9190610843565b610230565b6040516100ff9190610927565b60405180910390f35b610122600480360381019061011d91906109df565b6102f4565b60405161012f9190610a76565b60405180910390f35b610140610309565b005b61014a61031d565b6040516101579190610ae6565b60405180910390f35b6101686103ab565b6040516101759190610b17565b60405180910390f35b61019860048036038101906101939190610b88565b6103d4565b6040516101a59190610a76565b60405180910390f35b6101c860048036038101906101c39190610c64565b6103ec565b6040516101d59190610a76565b60405180910390f35b6101f860048036038101906101f39190610cfe565b610402565b005b60008060e01b6301ffc9a760e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614159050919050565b606061023a610486565b6000808473ffffffffffffffffffffffffffffffffffffffff16846040516102629190610d67565b6000604051808303816000865af19150503d806000811461029f576040519150601f19603f3d011682016040523d82523d6000602084013e6102a4565b606091505b5091509150816102e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e090610dca565b60405180910390fd5b809250505092915050565b600063150b7a0260e01b905095945050505050565b610311610486565b61031b6000610504565b565b6001805461032a90610e19565b80601f016020809104026020016040519081016040528092919081815260200182805461035690610e19565b80156103a35780601f10610378576101008083540402835291602001916103a3565b820191906000526020600020905b81548152906001019060200180831161038657829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600063bc197c8160e01b905098975050505050505050565b600063f23a6e6160e01b90509695505050505050565b61040a610486565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190610ebd565b60405180910390fd5b61048381610504565b50565b61048e6105c8565b73ffffffffffffffffffffffffffffffffffffffff166104ac6103ab565b73ffffffffffffffffffffffffffffffffffffffff1614610502576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f990610f29565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610619816105e4565b811461062457600080fd5b50565b60008135905061063681610610565b92915050565b600060208284031215610652576106516105da565b5b600061066084828501610627565b91505092915050565b60008115159050919050565b61067e81610669565b82525050565b60006020820190506106996000830184610675565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106ca8261069f565b9050919050565b6106da816106bf565b81146106e557600080fd5b50565b6000813590506106f7816106d1565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61075082610707565b810181811067ffffffffffffffff8211171561076f5761076e610718565b5b80604052505050565b60006107826105d0565b905061078e8282610747565b919050565b600067ffffffffffffffff8211156107ae576107ad610718565b5b6107b782610707565b9050602081019050919050565b82818337600083830152505050565b60006107e66107e184610793565b610778565b90508281526020810184848401111561080257610801610702565b5b61080d8482856107c4565b509392505050565b600082601f83011261082a576108296106fd565b5b813561083a8482602086016107d3565b91505092915050565b6000806040838503121561085a576108596105da565b5b6000610868858286016106e8565b925050602083013567ffffffffffffffff811115610889576108886105df565b5b61089585828601610815565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108d95780820151818401526020810190506108be565b838111156108e8576000848401525b50505050565b60006108f98261089f565b61090381856108aa565b93506109138185602086016108bb565b61091c81610707565b840191505092915050565b6000602082019050818103600083015261094181846108ee565b905092915050565b6000819050919050565b61095c81610949565b811461096757600080fd5b50565b60008135905061097981610953565b92915050565b600080fd5b600080fd5b60008083601f84011261099f5761099e6106fd565b5b8235905067ffffffffffffffff8111156109bc576109bb61097f565b5b6020830191508360018202830111156109d8576109d7610984565b5b9250929050565b6000806000806000608086880312156109fb576109fa6105da565b5b6000610a09888289016106e8565b9550506020610a1a888289016106e8565b9450506040610a2b8882890161096a565b935050606086013567ffffffffffffffff811115610a4c57610a4b6105df565b5b610a5888828901610989565b92509250509295509295909350565b610a70816105e4565b82525050565b6000602082019050610a8b6000830184610a67565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610ab882610a91565b610ac28185610a9c565b9350610ad28185602086016108bb565b610adb81610707565b840191505092915050565b60006020820190508181036000830152610b008184610aad565b905092915050565b610b11816106bf565b82525050565b6000602082019050610b2c6000830184610b08565b92915050565b60008083601f840112610b4857610b476106fd565b5b8235905067ffffffffffffffff811115610b6557610b6461097f565b5b602083019150836020820283011115610b8157610b80610984565b5b9250929050565b60008060008060008060008060a0898b031215610ba857610ba76105da565b5b6000610bb68b828c016106e8565b9850506020610bc78b828c016106e8565b975050604089013567ffffffffffffffff811115610be857610be76105df565b5b610bf48b828c01610b32565b9650965050606089013567ffffffffffffffff811115610c1757610c166105df565b5b610c238b828c01610b32565b9450945050608089013567ffffffffffffffff811115610c4657610c456105df565b5b610c528b828c01610989565b92509250509295985092959890939650565b60008060008060008060a08789031215610c8157610c806105da565b5b6000610c8f89828a016106e8565b9650506020610ca089828a016106e8565b9550506040610cb189828a0161096a565b9450506060610cc289828a0161096a565b935050608087013567ffffffffffffffff811115610ce357610ce26105df565b5b610cef89828a01610989565b92509250509295509295509295565b600060208284031215610d1457610d136105da565b5b6000610d22848285016106e8565b91505092915050565b600081905092915050565b6000610d418261089f565b610d4b8185610d2b565b9350610d5b8185602086016108bb565b80840191505092915050565b6000610d738284610d36565b915081905092915050565b7f4d6574686f642063616c6c206661696c65640000000000000000000000000000600082015250565b6000610db4601283610a9c565b9150610dbf82610d7e565b602082019050919050565b60006020820190508181036000830152610de381610da7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e3157607f821691505b60208210811415610e4557610e44610dea565b5b50919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610ea7602683610a9c565b9150610eb282610e4b565b604082019050919050565b60006020820190508181036000830152610ed681610e9a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610f13602083610a9c565b9150610f1e82610edd565b602082019050919050565b60006020820190508181036000830152610f4281610f06565b905091905056fea2646970667358221220cdcae9231c9c6d9e9a8352c5de90afe236e48cf6b818c7d936408a8447e248b964736f6c63430008090033a2646970667358221220b0fa52089ccd0f5a2eee88395976a5b3ed3a1a8dc20d192c0a7b7eb85882414164736f6c63430008090033",
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
