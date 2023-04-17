package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractData is the golang structure for table contract_data.
type ContractData struct {
	Id              int         `json:"id"               description:"自增ID"`
	Opcode          string      `json:"opcode"           description:"opcode"`
	ContractName    string      `json:"contract_name"    description:"contract name"`
	ContractAddress string      `json:"contract_address" description:"contract address"`
	ContractHash    string      `json:"contract_hash"    description:"contract hash"`
	GasUsed         int64       `json:"gas_used"         description:"gas 使用量"`
	GasUsdt         float64     `json:"gas_usdt"         description:"消耗的gas转为usdt"`
	ChainId         int64       `json:"chain_id"         description:"链id"`
	CreatedAt       *gtime.Time `json:"created_at"       description:"合约创建时间"`
	CurrentStatus   int         `json:"current_status"   description:"合约创建状态 0:任务提交 1:任务进行中 2:任务完成"`
}
