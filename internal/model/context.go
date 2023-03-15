package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session    *ghttp.Session // Session in context.
	Ciphertext *string        // Ciphertext in context.
}
