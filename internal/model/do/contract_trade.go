package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractTrade is the golang structure of table contract_trade for DAO operations like Where/Data.
type ContractTrade struct {
	g.Meta          `orm:"table:contract_trade, do:true"`
	Id              interface{} // 主键，自增长ID
	Opcode          interface{} // 唯一操作码
	TransactionHash interface{} // 交易hash
	UserAddress     interface{} // 用户地址
	TokenId         interface{} // token id
	TokenUri        interface{} // token Uri
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	AccountHash     interface{} // 用户创建时hash
}
