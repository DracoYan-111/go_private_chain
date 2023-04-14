package timed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/model/entity"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary() {
	var (
		err error
		ctx = gctx.New()
	)
	_, err = gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
		if job, _ := UndoneJob(); len(job) > 0 {
			fmt.Println("-------------任务执行开始--------------")
			StartJobUp(job)
			fmt.Println("-------------任务执行结束--------------")
		}
		if user, _ := UndoneUser(); len(user) > 0 {
			log.Println("-------------用户执行开始--------------")
			StartUserUp(user)
			fmt.Println("-------------用户执行结束--------------")
		}
	})
	if err != nil {
		panic(err)
	}
}

// UndoneJob 查询未完成的任务
func UndoneJob() (temps []*entity.GoTestDb, err error) {
	job, err := dao.GoTestDb.DB().Model("go_test_db").All("current_status", 0)
	if err != nil {
		log.Panicln("UndoneUser:数据库户获取异常", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(job.Json()), &temps)
	if err != nil {
		return nil, err
	}
	return temps, nil
}

// UndoneUser 查询未完成的用户
func UndoneUser() (temps []*entity.UserData, err error) {
	user, err := dao.UserData.DB().Model("user_data").All("current_status", 0)
	if err != nil {
		log.Panicln("UndoneUser:数据库户获取异常", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(user.Json()), &temps)
	if err != nil {
		return nil, err
	}
	return temps, nil
}
