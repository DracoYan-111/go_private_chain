package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractTrade is the golang structure for table contract_trade.
type ContractTrade struct {
	Id              int         `json:"id"               description:"主键，自增长ID"`
	Opcode          string      `json:"opcode"           description:"唯一操作码"`
	TransactionHash string      `json:"transaction_hash" description:"交易hash"`
	UserAddress     string      `json:"user_address"     description:"用户地址"`
	TokenId         int         `json:"token_id"         description:"token id"`
	TokenUri        string      `json:"token_uri"        description:"token Uri"`
	CreatedAt       *gtime.Time `json:"created_at"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updated_at"       description:"更新时间"`
	AccountHash     string      `json:"account_hash"     description:"用户创建时hash"`
}
