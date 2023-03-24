package main

import (
	"encoding/hex"
	"fmt"
	soliditysha3 "github.com/miguelmota/go-solidity-sha3"
	"go_private_chain/internal/deploy"
	"log"
	"math/big"
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
	deploy.Qianming()
}

func encodeMintFunction() {
	toAddress := "0x1234567890123456789012345678901234567890"
	amount := big.NewInt(1000000000000000000) // 1 ether
	functionName := []byte("transferOwnership")
	argTypes := []string{"address"}
	args := []interface{}{toAddress, amount}
	data := make([]byte, 0)
	data = append(data, soliditysha3.SoliditySHA3(functionName, argTypes)...)
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			data = append(data, soliditysha3.SoliditySHA3([]byte(v))...)
		case int:
			data = append(data, soliditysha3.SoliditySHA3(big.NewInt(int64(v)))...)
		case *big.Int:
			data = append(data, soliditysha3.SoliditySHA3(v)...)
		default:
			panic(fmt.Sprintf("Unsupported type: %T", v))
		}
	}
	hexData := hex.EncodeToString(data)
	fmt.Println(hexData)
}
