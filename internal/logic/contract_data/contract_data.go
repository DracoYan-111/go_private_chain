package contract_data

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
)

type (
	sContractData struct{}
)

func init() {
	service.RegisterContractData(New())
}

func New() service.IContractData {
	return &sContractData{}
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
func (s *sContractData) CreateJob(ctx context.Context, req string) error {
	// 将解密后的数据转换为结构体数据
	var temps []Temporary
	err := utility.DecryptStructure(req, &temps)
	if err != nil {
		return fmt.Errorf("解密密文失败(CreateJo):%s", err)
	}

	// 检查结构体内容是否正确
	var tempTwos []TemporaryTwo
	for _, v := range temps {
		if v.Opcode == "" || v.ContractName == "" || len(v.Opcode) != 19 || v.ChainId == 0 {
			return fmt.Errorf("内容检查失败(CreateJo):传入内容格式不正确")
		} else {
			tempTwo := TemporaryTwo{
				Opcode:       v.Opcode,
				ContractName: v.ContractName,
				ChainId:      v.ChainId,
			}
			tempTwos = append(tempTwos, tempTwo)
		}

		all, err := dao.ContractData.Ctx(ctx).Where("opcode", v.Opcode).All()
		if err != nil || len(all) > 0 {
			return errors.New("重复检查失败(CreateJob):opcode重复")
		}
	}

	// 查询是否存在重复

	// 插入数据库
	return dao.ContractData.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.ContractData.Ctx(ctx).Data(tempTwos).Batch(len(tempTwos)).Insert()
		return err
	})
}
