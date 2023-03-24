package deploy

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_private_chain/contracts/box721"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/model/entity"
	"go_private_chain/utility"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

// InteractiveContract 创建Box721合约
func InteractiveContract(contract *createBox721.CreateBox721, jobData *entity.GoTestDb, privateKeys string) (string, string, *big.Int, string) {
	auth, client := CreateConnection(privateKeys)
	opcode, _ := new(big.Int).SetString(jobData.Opcode, 10)
	// 防止重复修改
	if jobData.ContractAddress != "" && jobData.ContractHash != "" {
		return jobData.ContractAddress, jobData.ContractHash, big.NewInt(jobData.GasUsed), jobData.Opcode
	}
	loading, _ := utility.ReadConfigFile([]string{"web3.createBox721"})
	contractAddress := QueryContract(opcode, jobData.ContractName, jobData.ContractName, common.HexToAddress(loading["web3.createBox721"]), contract)
	tx, err := contract.CreatePair(auth, opcode, jobData.ContractName, jobData.ContractName, common.HexToAddress(loading["web3.createBox721"]))
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

func BulkIssuance(contract *createBox721.CreateBox721) {
	//loading, _ := utility.ReadConfigFile([]string{"web3.web3.privateKey1"})
	//auth, _ := CreateConnection(loading["web3.privateKey1"])

}

func mintNft(contract *box721.Box721) {

	//myAddress := []common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890"), common.HexToAddress("0x1234567890123456789012345678901234567890")}
	//myValue := []*big.Int{big.NewInt(0).SetInt64(1)}
	//myUri := []string{""}
	//opts := bind.CallOpts{}
	//encodedParams, err := contract.BatchSafeMint(nil, myAddress, myValue, myUri)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 构建函数调用数据
	//functionSelector := box721.Box721Caller{}.
	//data := append(functionSelector[:4], encodedParams...)

}

func Qianming() {
	// 打开本地文件
	file, err := os.Open("contracts/contract/Box721_sol_createBox721.abi")
	if err != nil {
		// 处理错误
	}
	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		// 处理错误
	}
	factoryABI, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		fmt.Println("Failed to parse ABI:", err)
		return
	}

	data, err := factoryABI.Pack("transferOwnership", common.HexToAddress("0x5bdAAD14EA49603fA1170175BB1B14eF707b1104"))
	if err != nil {
		fmt.Println("Failed to encode function data:", err)
		return
	}

	fmt.Printf("签名信息: 0x%x\n", data)
}

// QueryContract 查询合约地址
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
