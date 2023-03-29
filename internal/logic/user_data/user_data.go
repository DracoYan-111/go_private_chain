package user_data

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
	"log"
)

type (
	sUserData struct{}
)

func init() {
	service.RegisterUserData(New())
}

func New() service.IUserData {
	return &sUserData{}
}

type UserNickName struct {
	NickName string `json:"name"`
}

// CreateUserAddress 创建一个新用户地址
func (s *sUserData) CreateUserAddress(ctx context.Context, req string) (string, error) {
	// 解析请求数据
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return "", fmt.Errorf("CreateUserAddress: %s", err)
	}

	//将任务插入数据库
	dbUserData := entity.UserData{CurrentStatus: 0}
	insertUserData, err := dao.UserData.Ctx(ctx).Data(dbUserData).Insert()

	// 创建用户合约
	private := "web3.privateKey0"
	loading, _ := utility.ReadConfigFile([]string{"web3.accountsFactory", private})
	createBox := deploy.LoadWithAddress(loading["web3.accountsFactory"], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)
	log.Println(aesDecrypt)
	userAddress, txHash, opcode := deploy.InteractiveAccountContract(createBox, aesDecrypt, loading[private])

	// 检查地址和opcode是否已经存在
	dbUserAddress, err := g.Model("user_data u,contract_trade c").Where(g.Map{
		"u.user_address": userAddress,
		"c.user_address": userAddress,
	}).WhereOr("c.opcode", opcode).All()
	if err != nil || dbUserAddress.Len() != 0 {
		return "", fmt.Errorf("CreateUserAddress: %s", err)
	}

	// 更新数据库
	id, err := insertUserData.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("CreateUserAddress: %s", err)
	}
	dbUserData.UserAddress = userAddress
	dbUserData.CurrentStatus = 2
	dbUserData.Id = int(id)
	return userAddress, dao.ContractTrade.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.ContractTrade.Ctx(ctx).Data(
			entity.ContractTrade{
				AccountHash: txHash,
				UserAddress: userAddress,
				Opcode:      opcode,
			}).Insert()
		_, err = dao.UserData.Ctx(ctx).Data(dbUserData).Where("id", id).Update()
		return err
	})
}
