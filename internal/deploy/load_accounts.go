package deploy

import (
	"github.com/ethereum/go-ethereum/common"
	"go_private_chain/contracts/accountsFactory"
	"log"
	"math/big"
)

// InteractiveAccountContract 创建用户地址
func InteractiveAccountContract(contract *accountsFactory.AccountsFactory, name string, privateKeys string, opcode *big.Int) (string, string, error /*, *big.Int*/) {
	auth, _ := CreateConnection(privateKeys)
	accountsAddress := QueryAccountContract(opcode, name, contract)
	tx, err := contract.CreatePair(auth, opcode, name)
	if err != nil {
		log.Println("<==== loadAccounts:发起交易异常 ====>", err)
		return "", "", err
	}

	//time.Sleep(9 * time.Second)
	//
	//gasUsed, err := TransactionNews(client, tx.Hash().Hex())
	//if err != nil {
	//	log.Println(err)
	//}
	//gas := gasUsed.Mul(gasUsed, tx.GasPrice())

	return accountsAddress.Hex(), tx.Hash().Hex(), nil /*, gas*/
}

// QueryAccountContract 查询合约地址
func QueryAccountContract(_opcode *big.Int, _name string, contract *accountsFactory.AccountsFactory) common.Address {
	accountsAddress, err := contract.CalculateAddress(nil, _opcode, _name)
	if err != nil {
		log.Println("<==== loadAccounts:查询失败 ====>", err)
	}
	return accountsAddress
}
