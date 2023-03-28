package user_data

import (
	"context"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/internal/deploy"
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

func (s *sUserData) CreateUserAddress(ctx context.Context, req string) error {
	// 解析请求数据
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return err
	}

	private := "web3.privateKey0"
	loading, _ := utility.ReadConfigFile([]string{"web3.accountsFactory", private})
	createBox := deploy.LoadWithAddress(loading["web3.accountsFactory"], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)
	log.Println(aesDecrypt)
	a, _, _ := deploy.InteractiveAccountContract(createBox, aesDecrypt, loading[private])

	log.Println(a, "++++++++++++++++++++++++")
	return nil
}
