package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractTrade is the golang structure for table contract_trade.
type ContractTrade struct {
	Id              int         `json:"id"               description:"主键，自增长ID"`
	TransactionHash string      `json:"transaction_hash" description:"交易hash"`
	ContractAddress string      `json:"contract_address" description:"合约地址"`
	UserAddress     string      `json:"user_address"     description:"用户地址"`
	TokenId         string      `json:"token_id"         description:"token id"`
	TokenUri        string      `json:"token_uri"        description:"token Uri"`
	CreatedAt       *gtime.Time `json:"created_at"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updated_at"       description:"更新时间"`
}
