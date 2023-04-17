package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractDataDao is the data access object for table contract_data.
type ContractDataDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ContractDataColumns // columns contains all the column names of Table for convenient usage.
}

// ContractDataColumns defines and stores column names for table contract_data.
type ContractDataColumns struct {
	Id              string // 自增ID
	Opcode          string // opcode
	ContractName    string // contract name
	ContractAddress string // contract address
	ContractHash    string // contract hash
	GasUsed         string // gas 使用量
	GasUsdt         string // 消耗的gas转为usdt
	ChainId         string // 链id
	CreatedAt       string // 合约创建时间
	CurrentStatus   string // 合约创建状态 0:任务提交 1:任务进行中 2:任务完成
}

// contractDataColumns holds the columns for table contract_data.
var contractDataColumns = ContractDataColumns{
	Id:              "id",
	Opcode:          "opcode",
	ContractName:    "contract_name",
	ContractAddress: "contract_address",
	ContractHash:    "contract_hash",
	GasUsed:         "gas_used",
	GasUsdt:         "gas_usdt",
	ChainId:         "chain_id",
	CreatedAt:       "created_at",
	CurrentStatus:   "current_status",
}

// NewContractDataDao creates and returns a new DAO object for table data access.
func NewContractDataDao() *ContractDataDao {
	return &ContractDataDao{
		group:   "default",
		table:   "contract_data",
		columns: contractDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContractDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContractDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContractDataDao) Columns() ContractDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContractDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContractDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContractDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
