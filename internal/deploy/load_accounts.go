package deploy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_private_chain/contracts/accountsFactory"
	"log"
	"math/big"
)

// InteractiveAccountContract 创建用户地址
func InteractiveAccountContract(contract *accountsFactory.AccountsFactory, name string, privateKeys string, opcode *big.Int) (string, string, error) {
	auth, _ := CreateConnection(privateKeys)
	accountsAddress := QueryAccountContract(opcode, name, contract)
	if accountsAddress == common.HexToAddress("") {
		return "", "", fmt.Errorf("loadAccounts:预计算用户地址失败")
	}
	tx, err := contract.CreatePair(auth, opcode, name)
	if err != nil {
		return "", "", fmt.Errorf("loadAccounts:发起交易异常 %s", err)
	}
	return accountsAddress.Hex(), tx.Hash().Hex(), nil
}

// QueryAccountContract 查询合约地址
func QueryAccountContract(opcode *big.Int, name string, contract *accountsFactory.AccountsFactory) common.Address {
	accountsAddress, err := contract.CalculateAddress(nil, opcode, name)
	if err != nil {
		log.Println("<==== loadAccounts:查询失败 ====>", err)
		return common.HexToAddress("")
	}
	return accountsAddress
}
