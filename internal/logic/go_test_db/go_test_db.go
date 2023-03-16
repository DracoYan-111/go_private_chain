package go_test_db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
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

// Temporary 处理json数据
type Temporary struct {
	Opcode       string `json:"id"`
	ContractName string `json:"name"`
	ChainId      int64  `json:"chainId"`
}

// TemporaryTwo 处理数据库插入数据
type TemporaryTwo struct {
	Opcode       string `json:"opcode"`
	ContractName string `json:"contractName"`
	ChainId      int64  `json:"chainId"`
}

// CreateJob 是新合约部署的任务接口
func (s *sGoTestDb) CreateJob(ctx context.Context, req string) error {
	// 解析请求数据
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return err
	}

	// 将解密后的数据转换为结构体数据
	var temps []Temporary
	err = json.Unmarshal([]byte(aesDecrypt), &temps)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %s", err)
	}

	// 检查结构体内容是否正确
	var tempTwos []TemporaryTwo
	for _, v := range temps {
		if v.Opcode == "" || v.ContractName == "" || len(v.Opcode) != 19 || v.ChainId == 0 {
			return errors.New("illegal content")
		} else {
			tempTwo := TemporaryTwo{
				Opcode:       v.Opcode,
				ContractName: v.ContractName,
				ChainId:      v.ChainId,
			}
			tempTwos = append(tempTwos, tempTwo)
		}
	}

	// 插入数据库
	return dao.GoTestDb.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GoTestDb.Ctx(ctx).Data(tempTwos).Batch(len(tempTwos)).Insert()
		return err
	})
}
