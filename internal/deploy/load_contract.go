package deploy

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_private_chain/contracts/box721"
	"go_private_chain/contracts/createBox721"
	"log"
	"math/big"
	"math/rand"
	"time"
)

// GoInteractiveContract 互动合约
func GoInteractiveContract(contract *createBox721.CreateBox721, privateKeys string) {
	auth, _ := CreateConnection(privateKeys)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成19位随机数
	randNum := rand.Int63n(9000000000000000000) + 1000000000000000000

	tx, err := contract.CreatePair(auth, big.NewInt(randNum), "ttt", "ttt", common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	if err != nil {
		log.Println("<==== LoadContract:发起交易异常 ====>", err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())

}

// GoQueryContract 查询合约
func GoQueryContract(contract *box721.Box721) {
	name, err := contract.Name(nil)
	if err != nil {
		log.Println("<==== LoadContract:查询失败 ====>", err)
	}
	log.Println(name)
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

// goTransactionNews 查看使用的gas
func goTransactionNews(client *ethclient.Client, hash string) (*big.Int, error) {
	time.Sleep(7 * time.Second)

	txHash := common.HexToHash(hash)

	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Println(err)
	}

	if isPending {
		log.Println("Deployment:Transaction is being packaged")

		return new(big.Int).SetUint64(0), errors.New("Deployment:Transaction in progress")
	} else {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Println(err, "Deployment:Get the amount of gas used by the transaction")
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
