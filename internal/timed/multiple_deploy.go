package timed

import (
	"fmt"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/deploy"
	"go_private_chain/utility"
	"strconv"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		private := "web3.privateKey" + strconv.Itoa(id)
		loading, _ := utility.ReadConfigFile([]string{private}, "manifest/config/")
		//
		a := deploy.LoadWithAddress("0x959922bE3CAee4b8Cd9a407cc3ac1C251C2007B1", "createBox721", loading[private]).(*createBox721.CreateBox721)
		deploy.GoInteractiveContract(a, loading[private])
		//fmt.Println("worker", id, "开始任务", j)
		//time.Sleep(5 * time.Second)
		fmt.Println(a)
		//fmt.Println("worker", id, "结束任务", j)
		results <- j * 2
	}
}

func Mains() {
	const numJobs = 3
	for i := 0; i < 5; i++ {
		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)

		for w := 1; w <= numJobs; w++ {
			go worker(w, jobs, results)
		}

		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)

		for a := 1; a <= numJobs; a++ {
			<-results
		}
	}
}
