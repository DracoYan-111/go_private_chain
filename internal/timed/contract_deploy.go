package timed

import (
	"github.com/robfig/cron"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary() {

	cronJob := cron.New()
	spec := "*/500 * * * *"
	err := cronJob.AddFunc(spec, func() {
		log.Println("=============开始==============")
		Mains()

		log.Println("-------------结束--------------")

	})
	if err != nil {
		return
	}
	cronJob.Start()
}
