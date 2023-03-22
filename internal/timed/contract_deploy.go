package timed

import (
	"github.com/robfig/cron"
	"go_private_chain/internal/service"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary() {
	log.Println("??????????????????????????")
	cronJob := cron.New()
	spec := "*/30 * * * * ?"
	err := cronJob.AddFunc(spec, func() {
		log.Println("-------------开始--------------")
		job, err := service.GoTestDb().UndoneJob()
		if err != nil {
			log.Panicln("数据库户获取异常", err)
		}
		if len(job) > 0 {
			Mains(job)
		}

		log.Println("-------------结束--------------")
	})
	if err != nil {
		return
	}
	cronJob.Start()
}
