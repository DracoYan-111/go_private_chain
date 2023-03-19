package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoTestDbDao is the data access object for table go_test_db.
type GoTestDbDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns GoTestDbColumns // columns contains all the column names of Table for convenient usage.
}

// GoTestDbColumns defines and stores column names for table go_test_db.
type GoTestDbColumns struct {
	Id              string // 自增ID
	Opcode          string // opcode
	ContractName    string // contracts name
	ContractAddress string // contracts address
	ContractHash    string // contracts hash
	GasUsed         string // gas 使用量
	GasUsdt         string // 消耗的gas转为usdt
	ChainId         string // 链id
	CreatedAt       string // 合约创建时间
	CurrentStatus   string // 合约创建状态 0:任务提交 1:任务进行中 2:任务完成
}

// goTestDbColumns holds the columns for table go_test_db.
var goTestDbColumns = GoTestDbColumns{
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

// NewGoTestDbDao creates and returns a new DAO object for table data access.
func NewGoTestDbDao() *GoTestDbDao {
	return &GoTestDbDao{
		group:   "default",
		table:   "go_test_db",
		columns: goTestDbColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoTestDbDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoTestDbDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoTestDbDao) Columns() GoTestDbColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoTestDbDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoTestDbDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoTestDbDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
