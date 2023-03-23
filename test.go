package main

import (
	"fmt"
	"log"
	"time"
)

var objects = []string{"object 21", "object 22", "object 23", "object 24", "object 25", "object 26", "object 27", "object 28", "object 29", "object 3", "object 4", "object 5", "object 6", "object 7", "object 8", "object 9", "object 10"}

func lalal(id int) int {
	log.Println(id, "))))))))))))))))))))))))))))))))))")
	time.Sleep(10 * time.Second)
	return id + 1000000000
}

// bbbbbbbbbbb []string{"object25", "object23", "object24", "object21", "object25", "object26", "object27", "object28", "object29", "object3", "object4", "object5", "object6", "object7", "object8", "object9", "object10"}
func worker(id int, object string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", object)
		fmt.Println("worker", object, "开始任务+", j)
		a := lalal(id)
		//time.Sleep(10 * time.Second)
		fmt.Println("worker", object, "结束任务-", a, "++++++++++++++++")
		results <- j * 2
	}
}

func main() {
	//// 模拟从数据库获取到的对象列表

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
			go worker(j, objects[(i*5)+j], jobs, results)
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
