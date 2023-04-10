package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDataDao is the data access object for table user_data.
type UserDataDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserDataColumns // columns contains all the column names of Table for convenient usage.
}

// UserDataColumns defines and stores column names for table user_data.
type UserDataColumns struct {
	Id            string //
	UserId        string // 用户id
	Opcode        string // 唯一操作码
	UserNick      string // 用户昵称
	AccountHash   string // 用户创建时hash
	UserAddress   string // 用户地址
	ParentId      string // 父级信息id
	CurrentStatus string // 状态信息 0:任务已加入 1:任务已完成
}

// userDataColumns holds the columns for table user_data.
var userDataColumns = UserDataColumns{
	Id:            "id",
	UserId:        "user_id",
	Opcode:        "opcode",
	UserNick:      "user_nick",
	AccountHash:   "account_hash",
	UserAddress:   "user_address",
	ParentId:      "parent_id",
	CurrentStatus: "current_status",
}

// NewUserDataDao creates and returns a new DAO object for table data access.
func NewUserDataDao() *UserDataDao {
	return &UserDataDao{
		group:   "default",
		table:   "user_data",
		columns: userDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDataDao) Columns() UserDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
