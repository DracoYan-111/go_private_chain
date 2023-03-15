package go_test_db

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
	"log"
)

type (
	sGoTestDb struct{}
)

func init() {
	service.RegisterGoTestDb(New())
}

func New() service.IGoTestDb {
	return &sGoTestDb{}
}

// 处理json数据
type temporary struct {
	Opcode       string `json:"id"`
	ContractName string `json:"name"`
	ChainId      int64  `json:"chainId"`
}

// 处理数据库插入数据
type temporaryTwo struct {
	Opcode       string `json:"opcode"`
	ContractName string `json:"contract_name"`
	ChainId      int64  `json:"chain_id"`
}

// CreateJob Is the task interface for new contract deployment.
func (s *sGoTestDb) CreateJob(ctx context.Context, req string) error {
	//// 解析请求数据
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return err
	}

	// 将解密后的数据转换为结构体数据
	var temp []temporary
	var tempTwos []temporaryTwo
	err = json.Unmarshal([]byte(aesDecrypt), &temp)

	// 检查结构体内容是否正确
	for i, v := range temp {
		if v.Opcode == "" || v.ContractName == "" || len(v.Opcode) != 19 || v.ChainId == 0 {
			return errors.New("illegal content")
		} else {
			tempTwo := temporaryTwo{
				Opcode:       temp[i].Opcode,
				ContractName: temp[i].ContractName,
				ChainId:      temp[i].ChainId,
			}
			tempTwos = append(tempTwos, tempTwo)
		}
	}
	log.Println(tempTwos, len(temp), "------------")

	// 插入数据库
	return dao.GoTestDb.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GoTestDb.Ctx(ctx).Data(tempTwos).Batch(len(tempTwos)).Insert()
		return err
	})
}
