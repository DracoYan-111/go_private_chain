package timed

import (
	"github.com/robfig/cron"
	"go_private_chain/internal/service"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary() {
	cronJob := cron.New()
	spec := "*/60 * * * * ?"
	err := cronJob.AddFunc(spec, func() {
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
		return
	}
	cronJob.Start()
}
