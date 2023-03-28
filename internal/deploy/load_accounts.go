package deploy

import (
	"github.com/ethereum/go-ethereum/common"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/utility"
	"log"
	"math/big"
	"time"
)

// InteractiveAccountContract 创建用户地址
func InteractiveAccountContract(contract *accountsFactory.AccountsFactory, name, privateKeys string) (string, string, string, *big.Int) {
	auth, client := CreateConnection(privateKeys)
	opcode := utility.RandomNumber()
	accountsAddress := QueryAccountContract(opcode, name, contract)
	tx, err := contract.CreatePair(auth, opcode, name)
	if err != nil {
		log.Println("<==== loadAccounts:发起交易异常 ====>", err)
	}

	time.Sleep(9 * time.Second)

	gasUsed, err := TransactionNews(client, tx.Hash().Hex())
	if err != nil {
		log.Println(err)
	}
	gas := gasUsed.Mul(gasUsed, tx.GasPrice())

	return accountsAddress.Hex(), tx.Hash().Hex(), opcode.Text(10), gas
}

// QueryAccountContract 查询合约地址
func QueryAccountContract(_opcode *big.Int, _name string, contract *accountsFactory.AccountsFactory) common.Address {
	accountsAddress, err := contract.CalculateAddress(nil, _opcode, _name)
	if err != nil {
		log.Println("<==== loadAccounts:查询失败 ====>", err)
	}
	return accountsAddress
}
