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

func init() {
	service.RegisterTesDBCtx(New())
}

func New() service.ITesDBCtx {
	return &sTesDBCtx{}
}

// Init 方法被s*sTesDBCtx调用,传入的参数为*ghttp.Request,*model.Context
func (s *sTesDBCtx) Init(req *ghttp.Request, ctx *model.Context) {
	req.SetCtxVar(consts.ContextKey, ctx)
}

// Get 方法被s*sTesDBCtx调用,传入的参数为context.Context返回 *model.Context
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

// SetCiphertext 方法被s*sTesDBCtx调用,传入的参数为context.Context,*string
func (s *sTesDBCtx) SetCiphertext(ctx context.Context, ctxGoTestDb *string) {
	s.Get(ctx).Ciphertext = ctxGoTestDb
}
