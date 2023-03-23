package deploy

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_private_chain/contracts/box721"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/model/entity"
	"log"
	"math/big"
	"time"
)

// InteractiveContract 互动合约
func InteractiveContract(contract *createBox721.CreateBox721, jobData *entity.GoTestDb, privateKeys string) (string, string, *big.Int, string) {
	auth, client := CreateConnection(privateKeys)
	opcode, _ := new(big.Int).SetString(jobData.Opcode, 10)

	contractAddress := QueryContract(opcode, jobData.ContractName, jobData.ContractName, common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"), contract)
	tx, err := contract.CreatePair(auth, opcode, jobData.ContractName, jobData.ContractName, common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	if err != nil {
		log.Println("<==== LoadContract:发起交易异常 ====>", err)
	}

	time.Sleep(9 * time.Second)

	gasUsed, err := transactionNews(client, tx.Hash().Hex())
	if err != nil {
		log.Println(err)
	}
	gas := gasUsed.Mul(gasUsed, tx.GasPrice())

	return contractAddress.Hex(), tx.Hash().Hex(), gas, jobData.Opcode
}

// QueryContract 查询合约
func QueryContract(_opcode *big.Int, _name string, _symbol string, _minter common.Address, contract *createBox721.CreateBox721) common.Address {
	contractAddress, err := contract.CalculateAddress(nil, _opcode, _name, _symbol, _minter)
	if err != nil {
		log.Println("<==== LoadContract:查询失败 ====>", err)
	}
	return contractAddress
}

// GoCreateAndGenerate 通过地址创建合约并生成合约实例
//func GoCreateAndGenerate(structure Structure, t *testing.T) *box721.Box721 {
//	//contractAddr := GoContractDeployment(structure)
//	_, address, _, _ := GoContractDeployment(structure)
//	example := GoLoadWithAddress(address, t)
//	GoQueryContract(example, t)
//	GoInteractiveContract(example, t)
//	return example
//}

// transactionNews 查看使用的gas
func transactionNews(client *ethclient.Client, hash string) (*big.Int, error) {
	txHash := common.HexToHash(hash)

	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return new(big.Int).SetUint64(0), errors.New("<==== LoadContract:哈希交易检查失败 ====>")
	}
	if isPending {
		return new(big.Int).SetUint64(0), errors.New("<==== LoadContract:交易进行中 ====>")
	} else {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return new(big.Int).SetUint64(0), errors.New("<==== LoadContract:获取交易使用的gas量失败 ====>")
		}
		return new(big.Int).SetUint64(receipt.GasUsed), nil
	}
}

// LoadWithAddress 通过地址生成合约实例
func LoadWithAddress(contractAddr, contractType, privateKeys string) interface{} {
	_, client := CreateConnection(privateKeys)

	var instance interface{}
	var err error
	switch {
	case contractType == "box721":
		instance, err = box721.NewBox721(common.HexToAddress(contractAddr), client)
	case contractType == "createBox721":
		instance, err = createBox721.NewCreateBox721(common.HexToAddress(contractAddr), client)
	}
	if err != nil {
		log.Println("<==== loadContract:生成合约实例失败 ====>", err)
	}
	return instance
}
