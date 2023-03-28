package v1

import "github.com/gogf/gf/v2/frame/g"

type NewAdditionalTasksReq struct {
	g.Meta     `path:"/user/new-added-task" method:"post" tags:"AddedService" summary:"上传一个新的藏品增发任务"`
	Ciphertext string `v:"required"`
}

type NewAdditionalTasksRes struct {
	OK bool `summary:"任务状态"`
}
