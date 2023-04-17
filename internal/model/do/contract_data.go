package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractData is the golang structure of table contract_data for DAO operations like Where/Data.
type ContractData struct {
	g.Meta          `orm:"table:contract_data, do:true"`
	Id              interface{} // 自增ID
	Opcode          interface{} // opcode
	ContractName    interface{} // contract name
	ContractAddress interface{} // contract address
	ContractHash    interface{} // contract hash
	GasUsed         interface{} // gas 使用量
	GasUsdt         interface{} // 消耗的gas转为usdt
	ChainId         interface{} // 链id
	CreatedAt       *gtime.Time // 合约创建时间
	CurrentStatus   interface{} // 合约创建状态 0:任务提交 1:任务进行中 2:任务完成
}
