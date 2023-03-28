package timed

import (
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
	"log"
	"math/big"
	"strconv"
	"time"
)

type WorkerResult struct {
	opcode  string
	Address string
	Hash    string
	Gas     *big.Int
}

// worker 执行多线程合约创建任务
func worker(id int, data *entity.GoTestDb, jobs <-chan int, results chan<- *WorkerResult) {
	private := "web3.privateKey" + strconv.Itoa(id)
	loading, _ := utility.ReadConfigFile([]string{"web3.createBox721", private})
	createBox := deploy.LoadWithAddress(loading["web3.createBox721"], "createBox721", loading[private]).(*createBox721.CreateBox721)
	for range jobs {
		address, hash, gas, opcode := deploy.InteractiveNftContract(createBox, data, loading[private])
		time.Sleep(1 * time.Second)
		results <- &WorkerResult{opcode, address, hash, gas}
	}
}

// StartUp 启动线程
func StartUp(jobData []*entity.GoTestDb) {
	var numLoopsTwo = 5
	var integrate = make(map[string]*WorkerResult)

	//计算需要循环的次数
	numLoops := len(jobData) / 5
	if len(jobData)%5 != 0 {
		numLoops++
	}
	for i := 0; i < numLoops; i++ {
		if len(jobData)%5 > 0 && i == numLoops-1 {
			numLoopsTwo = len(jobData) % 5
		}
		jobs := make(chan int, numLoopsTwo)
		results := make(chan *WorkerResult, numLoopsTwo)
		for j := 0; j < numLoopsTwo; j++ {
			go worker(j, jobData[(i*5)+j], jobs, results)
		}
		for k := 0; k < numLoopsTwo; k++ {
			jobs <- k
		}
		close(jobs)
		for j := 0; j < numLoopsTwo; j++ {
			result := <-results
			integrate[result.opcode] = result
		}
	}
	processStructure(jobData, integrate)
}

// processStructure 格式化结构并插入数据库
func processStructure(jobData []*entity.GoTestDb, payload map[string]*WorkerResult) {
	for _, single := range jobData {
		single.ContractAddress = payload[single.Opcode].Address
		single.ContractHash = payload[single.Opcode].Hash
		single.GasUsed = payload[single.Opcode].Gas.Int64()
		single.CurrentStatus = 2
	}
	err := service.GoTestDb().UpdateJob(jobData)
	if err != nil {
		log.Println(err)
	}
}
