// Package v1 针对v1版本的操作合约接口
package v1

import "github.com/gogf/gf/v2/frame/g"

type NewJobTaskReq struct {
	g.Meta     `path:"/job/new-job-task" method:"post" tags:"CreateNftContract" summary:"上传一个创建新藏品的工作任务"`
	Ciphertext string `v:"required"`
}

type NewJobTaskRes struct {
	OK bool `v:"required" summary:"任务状态"`
}
