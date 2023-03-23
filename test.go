package main

import (
	"fmt"
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "开始任务", j)
		time.Sleep(1 * time.Second)
		fmt.Println("worker", id, "结束任务", j)
		results <- j * 2
	}
}

func main() {
	//// 模拟从数据库获取到的对象列表
	objects := []string{"object2", "object2", "object2", "object2", "object2", "object2", "object2", "object2", "object2", "object3", "object4", "object5", "object6", "object7", "object8", "object9", "object10"}

	//计算需要循环的次数
	numLoops := len(objects) / 5
	if len(objects)%5 != 0 {
		numLoops++
	}
	var numLoopsTwo = 5
	for i := 0; i < numLoops; i++ {
		log.Println("=======+++++++++++======")
		if len(objects)%5 > 0 && i == numLoops-1 {
			numLoopsTwo = len(objects) % 5
		}
		jobs := make(chan int, numLoopsTwo)
		results := make(chan int, numLoopsTwo)
		for j := 0; j < numLoopsTwo; j++ {
			go worker(j, jobs, results)
		}
		for k := 0; k < numLoopsTwo; k++ {
			jobs <- k
		}
		close(jobs)
		for j := 0; j < numLoopsTwo; j++ {
			<-results
		}
	}
}
