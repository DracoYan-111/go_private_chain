package deploy

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go_private_chain/contracts/accounts"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/contracts/box721"
	"go_private_chain/contracts/contractCall"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/utility"
	"log"
	"math/big"
)

type Structure struct {
	Name           string
	Symbol         string
	Minter         common.Address
	TokenURIPrefix string
}

// ContractDeployment 创建合约并返回合约地址
func ContractDeployment(privateKeys string) (string, string, *big.Int, int64) {
	auth, client := CreateConnection(privateKeys)

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)

	if balance.Int64() < 5e16 {
		log.Println("Deployment:Insufficient user balance", balance)

		return "", "", big.NewInt(0), 0
	}
	address, txData, _, err := box721.DeployBox721(
		auth,
		client,
		"structure.Name",
		"structure.Symbol",
	)

	if err != nil {
		log.Println("<==== Deployment:创建合约异常 ====>", err)

		return "", "", big.NewInt(0), 0
	}
	log.Println("structure.Name", "Deployment:开始部署:", txData.Hash().Hex())

	gasUsed, err := TransactionNews(client, txData.Hash().Hex())

	gas := gasUsed.Mul(gasUsed, txData.GasPrice())
	//time.Sleep(3 * time.Second)

	return address.Hex(), txData.Hash().Hex(), gas.Add(gas, big.NewInt(5e12)), 1

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
	case contractType == "contractCall":
		instance, err = contractCall.NewContractCall(common.HexToAddress(contractAddr), client)
	case contractType == "accounts":
		instance, err = accounts.NewAccounts(common.HexToAddress(contractAddr), client)
	case contractType == "accountsFactory":
		instance, err = accountsFactory.NewAccountsFactory(common.HexToAddress(contractAddr), client)
	}
	if err != nil {
		log.Println("<==== loadContract:生成合约实例失败 ====>", err)
	}
	return instance
}

// CreateConnection 创建区块链连接
func CreateConnection(privateKeys string) (*bind.TransactOpts, *ethclient.Client) {
	var client *ethclient.Client
	var err error

	loading, err := utility.ReadConfigFile([]string{"web3.rpcUrl"})
	if err != nil {
		log.Panicln("<==== Deployment:文件读取失败 ====>:", err)
	}
	client, err = ethclient.Dial(loading["web3.rpcUrl"])
	if err != nil {
		log.Println("<==== Deployment:连接到节点异常 ====>", err)
	}

	//创建私钥实例
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		log.Println("<==== Deployment:异常加载私钥 ====>", err)
	}

	//获取当前链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("<==== Deployment:获取链ID异常 ====>", err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	//获取当前用户的最新随机数
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Println("<==== Deployment:最新的随机数异常 ====>", err)
	}

	//预估gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("<==== Deployment:Gas 价格异常 ====>", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5100000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}

// TransactionNews 查看使用的gas
func TransactionNews(client *ethclient.Client, hash string) (*big.Int, error) {
	txHash := common.HexToHash(hash)

	_, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return new(big.Int).SetUint64(0), fmt.Errorf("<==== LoadContract:哈希交易检查失败 %s ====>", err)
	}
	if isPending {
		return new(big.Int).SetUint64(0), fmt.Errorf("<==== LoadContract:交易进行中 ====>")
	} else {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return new(big.Int).SetUint64(0), fmt.Errorf("<==== LoadContract:获取交易使用的gas量失败 %s ====>", err)
		}
		return new(big.Int).SetUint64(receipt.GasUsed), nil
	}
}
