package model

type UserCreateInput struct {
	Opcode       string // 任务唯一标识
	ContractName string // 合约名称
	ChainId      int64  // 链ID
}
