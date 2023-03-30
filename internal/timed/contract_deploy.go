package timed

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"go_private_chain/internal/service"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary() {
	//cronJob := cron.New()
	//spec := "*/60 * * * * ?"
	//err := cronJob.AddFunc(spec, func() {
	//	job, err := service.GoTestDb().UndoneJob()
	//	if err != nil {
	//		log.Panicln("UpdateLibrary:数据库户获取异常", err)
	//	}
	//	if len(job) > 0 {
	//		log.Println("-------------任务执行开始--------------")
	//		StartUp(job)
	//		log.Println("-------------任务执行结束--------------")
	//	}
	//})
	//if err != nil {
	//	return
	//}
	//cronJob.Start()

	var (
		err error
		ctx = gctx.New()
		//logger = glog.New()
	)
	//logger.SetLevel(glog.LEVEL_DEBU)
	//gcron.SetLogger(logger)
	_, err = gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
		//		g.Log().Info(ctx, "test")
		job, err := service.GoTestDb().UndoneJob()
		if err != nil {
			log.Panicln("UpdateLibrary:数据库户获取异常", err)
		}
		if len(job) > 0 {
			log.Println("-------------任务执行开始--------------")
			StartUp(job)
			log.Println("-------------任务执行结束--------------")
		}
	})
	if err != nil {
		panic(err)
	}
}
