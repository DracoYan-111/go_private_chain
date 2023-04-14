package timed

import (
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/consts"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/utility"
	"log"
	"math/big"
	"strconv"
	"time"
)

type WorkerJobResult struct {
	opcode  string
	Address string
	Hash    string
	Gas     *big.Int
}

// workerJob 执行多线程合约创建任务
func workerJob(id int, data *entity.GoTestDb, jobs <-chan int, results chan<- *WorkerJobResult) {
	private := consts.PrivateKey + strconv.Itoa(id)
	loading, _ := utility.ReadConfigFile([]string{consts.CreateBox721, private})
	createBox := deploy.LoadWithAddress(loading[consts.CreateBox721], "createBox721", loading[private]).(*createBox721.CreateBox721)
	for range jobs {
		address, hash, gas, opcode := deploy.InteractiveNftContract(createBox, data, loading[private])
		time.Sleep(1 * time.Second)
		results <- &WorkerJobResult{opcode, address, hash, gas}
	}
}

// StartJobUp 启动线程
func StartJobUp(jobData []*entity.GoTestDb) {
	var numLoopsTwo = 5
	var integrate = make(map[string]*WorkerJobResult)

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
		results := make(chan *WorkerJobResult, numLoopsTwo)
		for j := 0; j < numLoopsTwo; j++ {
			go workerJob(j, jobData[(i*5)+j], jobs, results)
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
	processJobStructure(jobData, integrate)
}

// processJobStructure 格式化结构并插入数据库
func processJobStructure(jobData []*entity.GoTestDb, payload map[string]*WorkerJobResult) {
	for _, single := range jobData {
		single.ContractAddress = payload[single.Opcode].Address
		single.ContractHash = payload[single.Opcode].Hash
		single.GasUsed = payload[single.Opcode].Gas.Int64()
		single.CurrentStatus = 2
	}
	for _, gtd := range jobData {
		_, err := dao.GoTestDb.DB().Model("go_test_db").Data(gtd).Where("id = ?", gtd.Id).Update()
		if err != nil {
			log.Println("update go_test_db error", err)
		}
	}
	log.Println(len(jobData), "条数据已经更新")
}

type WorkerUserResult struct {
	opcode      string
	userAddress string
	txHash      string
	dbId        int
}

// workerUser 执行多线程用户创建任务
func workerUser(id int, data *entity.UserData, jobs <-chan int, results chan<- *WorkerUserResult) {
	private := consts.PrivateKey + strconv.Itoa(id)
	loading, _ := utility.ReadConfigFile([]string{consts.AccountsFactory, private})
	createAccounts := deploy.LoadWithAddress(loading[consts.AccountsFactory], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)
	for range jobs {
		n := new(big.Int)
		n, _ = n.SetString(data.Opcode, 10)
		userAddress, txHash, _ := deploy.InteractiveAccountContract(createAccounts, data.UserNick, loading[private], n)
		time.Sleep(1 * time.Second)
		results <- &WorkerUserResult{data.Opcode, userAddress, txHash, data.Id}
	}
}

// StartUserUp 启动线程
func StartUserUp(jobData []*entity.UserData) {
	var numLoopsTwo = 14
	var integrate = make(map[string]*WorkerUserResult)

	//计算需要循环的次数
	numLoops := len(jobData) / 14
	if len(jobData)%14 != 0 {
		numLoops++
	}
	for i := 0; i < numLoops; i++ {
		if len(jobData)%14 > 0 && i == numLoops-1 {
			numLoopsTwo = len(jobData) % 14
		}
		jobs := make(chan int, numLoopsTwo)
		results := make(chan *WorkerUserResult, numLoopsTwo)
		for j := 0; j < numLoopsTwo; j++ {
			go workerUser(j, jobData[(i*14)+j], jobs, results)
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
	processUserStructure(jobData, integrate)
}

func processUserStructure(jobData []*entity.UserData, payload map[string]*WorkerUserResult) {
	for _, single := range jobData {
		single.UserAddress = payload[single.Opcode].userAddress
		single.AccountHash = payload[single.Opcode].txHash
		if payload[single.Opcode].txHash != "" {
			single.CurrentStatus = 2
		}
	}
	for _, gtd := range jobData {
		_, err := dao.GoTestDb.DB().Model("user_data").Data(gtd).Where("id", gtd.Id).Update()
		if err != nil {
			log.Println("update user_data error:", err)
		}
	}
	log.Println(len(jobData), "条数据已更新")
}
