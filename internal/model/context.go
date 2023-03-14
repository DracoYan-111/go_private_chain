package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session   *ghttp.Session   // Session in context.
	CGoTestDb *ContextGoTestDb // CGoTestDb in context.
}

type ContextGoTestDb struct {
	Opcode       string // 任务唯一标识
	ContractName string // 合约名称
	ChainId      int64  // 链ID
}
