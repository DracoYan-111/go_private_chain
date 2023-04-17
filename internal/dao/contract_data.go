package dao

import (
	"go_private_chain/internal/dao/internal"
)

// internalContractDataDao is internal type for wrapping internal DAO implements.
type internalContractDataDao = *internal.ContractDataDao

// contractDataDao is the data access object for table contract_data.
// You can define custom methods on it to extend its functionality as you wish.
type contractDataDao struct {
	internalContractDataDao
}

var (
	// ContractData is globally public accessible object for table contract_data operations.
	ContractData = contractDataDao{
		internal.NewContractDataDao(),
	}
)

// Fill with you ideas below.
