package deploy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/contracts/contractCall"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/consts"
	"go_private_chain/internal/model/entity"
	"go_private_chain/utility"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// InteractiveNftContract 创建Box721合约
func InteractiveNftContract(contract *createBox721.CreateBox721, jobData *entity.GoTestDb, privateKeys string) (string, string, *big.Int, string) {
	auth, client := CreateConnection(privateKeys)
	opcode, _ := new(big.Int).SetString(jobData.Opcode, 10)

	// 防止重复修改
	if jobData.ContractAddress != "" || jobData.ContractHash != "" {
		return jobData.ContractAddress, jobData.ContractHash, big.NewInt(jobData.GasUsed), jobData.Opcode
	}
	contractAddress := QueryNftContract(opcode, jobData.ContractName, jobData.ContractName, contract)
	if contractAddress == common.HexToAddress("") {
		return "", "", big.NewInt(0), ""
	}
	tx, err := contract.CreatePair(auth, opcode, jobData.ContractName, jobData.ContractName)
	if err != nil {
		log.Println("<==== LoadContract:发起交易异常 ====>", err)
		return "", "", big.NewInt(0), ""
	}

	// 等待链上确定
	time.Sleep(9 * time.Second)

	gasUsed, err := TransactionNews(client, tx.Hash().Hex())
	if err != nil {
		log.Println(err)
	}
	gas := gasUsed.Mul(gasUsed, tx.GasPrice())

	return contractAddress.Hex(), tx.Hash().Hex(), gas, jobData.Opcode
}

// BulkIssuance 批量增发方法
func BulkIssuance(createBox721 *createBox721.CreateBox721, box721Address common.Address, tos []common.Address, tokenIds []string, uris []string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	private := consts.PrivateKey + strconv.Itoa(rand.Intn(5))
	loading, _ := utility.ReadConfigFile([]string{private})
	auth, _ := CreateConnection(loading[private])
	sig, err := Signature(tos, tokenIds, uris, loading[private])

	if err != nil {
		return "", err
	}
	callBox721, err := createBox721.CallBox721(auth, box721Address, sig)
	if err != nil {
		return "", err
	}

	return callBox721.Hash().String(), nil

}

// BulkTransfer 批量转移
func BulkTransfer(userAddress, receiveAddress []common.Address, correct []*big.Int, contractAddress common.Address) (string, error) {
	rand.Seed(time.Now().UnixNano())
	private := consts.PrivateKey + strconv.Itoa(rand.Intn(5))
	loading, _ := utility.ReadConfigFile([]string{consts.ContractCall, consts.AccountsFactory, private})
	// 获取批量转移字节码
	contractCallContract := LoadWithAddress(loading[consts.ContractCall], "contractCall", loading[private]).(*contractCall.ContractCall)
	callData, err := contractCallContract.BatchSecurityTransferCall(nil, userAddress, receiveAddress, correct)
	if err != nil {
		return "", fmt.Errorf("BatchSecurityTransferCall: %s", err)
	}
	// 获取调用账户字节码
	accountsFactoryContract := LoadWithAddress(loading[consts.AccountsFactory], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)
	calldata, err := accountsFactoryContract.CallContractCall(nil, contractAddress, callData)
	if err != nil {
		return "", fmt.Errorf("CallContractCall: %s", err)
	}
	// 调用用户合约
	auth, _ := CreateConnection(loading[private])

	accounts, err := accountsFactoryContract.CallAccounts(auth, userAddress[0], calldata)
	if err != nil {
		return "", fmt.Errorf("CallAccounts: %s", err)
	}
	return accounts.Hash().Hex(), nil
}

// Signature 获取方法签名信息
func Signature(tos []common.Address, tokenIds []string, uris []string, private string) ([]byte, error) {
	loading, _ := utility.ReadConfigFile([]string{consts.ContractCall})
	createBox := LoadWithAddress(loading[consts.ContractCall], "contractCall", private).(*contractCall.ContractCall)

	call, err := createBox.BatchSafeMintCall(nil, tos, tokenIds, uris)
	if err != nil {
		return nil, err
	}
	return call, nil
}

// QueryNftContract 查询合约地址
func QueryNftContract(opcode *big.Int, name string, symbol string, contract *createBox721.CreateBox721) common.Address {
	contractAddress, err := contract.CalculateAddress(nil, opcode, name, symbol)
	if err != nil {
		log.Println("<==== LoadContract:查询失败 ====>", err)
		return common.HexToAddress("")
	}
	return contractAddress
}
