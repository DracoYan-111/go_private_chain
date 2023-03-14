package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"go_private_chain/internal/consts"
	"go_private_chain/internal/model"
	"go_private_chain/internal/service"
)

// 创建type类型的sTesDBCtx结构体
type (
	sTesDBCtx struct{}
)

// init方法中调用service.RegisterTesDBCtx方法,传入的参数为service.ITesDBCtx类型的指针
func init() {
	service.RegisterTesDBCtx(New())
}

// New 方法返回service.ITesDBCtx类型的指针
func New() service.ITesDBCtx {
	return &sTesDBCtx{}
}

// Init 方法被s*sTesDBCtx调用,传入的参数为*ghttp.Request,*model.Context
func (s *sTesDBCtx) Init(r *ghttp.Request, ctx *model.Context) {
	r.SetCtxVar(consts.ContextKey, ctx)
}

// Get 方法被s*sTesDBCtx调用,传入的参数为model.Context;return *model.Context
func (s *sTesDBCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetTestDb 方法被s*sTesDBCtx调用,传入的参数为model.Context,*model.ContextGoTestDb
func (s *sTesDBCtx) SetTestDb(ctx context.Context, testDb *model.ContextGoTestDb) {
	s.Get(ctx).CGoTestDb = testDb
}
