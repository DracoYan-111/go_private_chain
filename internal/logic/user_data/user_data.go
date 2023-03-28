package user_data

import (
	"context"
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

func (s *sUserData) CreateUserAddress(ctx context.Context, req string) error {
	// 解析请求数据
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return err
	}
	log.Println(aesDecrypt, "=====================")

	return nil
}
