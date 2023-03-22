package timed

import (
	"fmt"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/utility"
	"log"
	"strconv"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		private := "web3.privateKey" + strconv.Itoa(id)
		loading, _ := utility.ReadConfigFile([]string{private})
		a := deploy.LoadWithAddress("0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1", "createBox721", loading[private]).(*createBox721.CreateBox721)
		deploy.GoInteractiveContract(a, loading[private])
		fmt.Println("worker", id, "开始任务", j)

		//time.Sleep(3 * time.Second)

		fmt.Println(a)
		fmt.Println("worker", id, "结束任务", j)
		results <- j * 2
	}
}

func Mains(jobData []*entity.GoTestDb) {
	//计算需要循环的次数
	numLoops := len(jobData) / 5
	if len(jobData)%5 != 0 {
		numLoops++
	}
	var numLoopsTwo = 5
	for i := 0; i < numLoops; i++ {
		log.Println("=======+++++++++++======")
		if len(jobData)%5 > 0 && i == numLoops-1 {
			numLoopsTwo = len(jobData) % 5
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
