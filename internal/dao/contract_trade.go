package dao

import (
	"go_private_chain/internal/dao/internal"
)

// internalContractTradeDao is internal type for wrapping internal DAO implements.
type internalContractTradeDao = *internal.ContractTradeDao

// contractTradeDao is the data access object for table contract_trade.
// You can define custom methods on it to extend its functionality as you wish.
type contractTradeDao struct {
	internalContractTradeDao
}

var (
	// ContractTrade is globally public accessible object for table contract_trade operations.
	ContractTrade = contractTradeDao{
		internal.NewContractTradeDao(),
	}
)
