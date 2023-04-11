package deploy

import (
	"github.com/ethereum/go-ethereum/common"
	"go_private_chain/contracts/contractCall"
	"go_private_chain/contracts/createBox721"
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
	private := "web3.accountsKey.privateKey" + strconv.Itoa(rand.Intn(5))
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

// Signature 获取方法签名信息
func Signature(tos []common.Address, tokenIds []string, uris []string, private string) ([]byte, error) {
	loading, _ := utility.ReadConfigFile([]string{"web3.contractCall"})
	createBox := LoadWithAddress(loading["web3.contractCall"], "contractCall", private).(*contractCall.ContractCall)

	call, err := createBox.BatchSafeMintCall(nil, tos, tokenIds, uris)
	if err != nil {
		return nil, err
	}
	return call, nil
}

// QueryNftContract 查询合约地址
func QueryNftContract(_opcode *big.Int, _name string, _symbol string, contract *createBox721.CreateBox721) common.Address {
	contractAddress, err := contract.CalculateAddress(nil, _opcode, _name, _symbol)
	if err != nil {
		log.Println("<==== LoadContract:查询失败 ====>", err)
		return common.HexToAddress("")
	}
	return contractAddress
}
