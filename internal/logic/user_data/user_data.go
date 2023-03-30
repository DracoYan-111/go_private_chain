package user_data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
	"log"
	"math/big"
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

	// 检查opcode是否已经存在
	opcode := utility.RandomNumber()
	dbUserAddress, err := dao.UserData.Ctx(ctx).Where("opcode", opcode).All()
	if err != nil || dbUserAddress.Len() != 0 {
		return "", fmt.Errorf("CreateUserAddress: %s", err)
	}

	//将任务插入数据库
	dbUserData := entity.UserData{
		Opcode:        opcode.Text(10),
		UserNick:      aesDecrypt,
		CurrentStatus: 0,
	}
	insertUserData, err := dao.UserData.Ctx(ctx).Data(dbUserData).Insert()
	if err != nil {
		return "", fmt.Errorf("CreateUserAddress: %s", err)
	}

	// 创建用户合约
	private := "web3.accountsKey.privateKey0"
	loading, _ := utility.ReadConfigFile([]string{"web3.accountsFactory", private})
	createBox := deploy.LoadWithAddress(loading["web3.accountsFactory"], "accountsFactory", loading[private]).(*accountsFactory.AccountsFactory)
	userAddress, txHash, err := deploy.InteractiveAccountContract(createBox, aesDecrypt, loading[private], opcode)
	if err != nil {
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
	dbUserData.AccountHash = txHash
	return userAddress, dao.UserData.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.UserData.Ctx(ctx).Data(dbUserData).Where("id", id).Update()
		return err
	})
}

type additionalInfo struct {
	ContractAddress common.Address   `json:"contract"`
	Tos             []common.Address `json:"tos"`
	TokenIds        []*big.Int       `json:"tokenIds"`
	Uris            []string         `json:"uris"`
}

// BatchCastingNft 新的批量创建nft任务
func (s *sUserData) BatchCastingNft(ctx context.Context, req string) (string, error) {
	aesDecrypt, err := utility.AesDecrypt(req)
	if err != nil {
		return "", err
	}

	// 将解密后的数据转换为结构体数据
	var temp additionalInfo
	err = json.Unmarshal([]byte(aesDecrypt), &temp)
	if err != nil {
		return "", err
	}

	if len(temp.Tos) == len(temp.Uris) && len(temp.Uris) == len(temp.TokenIds) {
		// 创建用户合约
		private := "web3.accountsKey.privateKey0"
		loading, _ := utility.ReadConfigFile([]string{"web3.createBox721", private})
		createBox := deploy.LoadWithAddress(loading["web3.createBox721"], "createBox721", loading[private]).(*createBox721.CreateBox721)
		transactionHash, err := deploy.BulkIssuance(createBox, temp.ContractAddress, temp.Tos, temp.TokenIds, temp.Uris)
		if err != nil {
			return "", err
		}

		dbAdditionalInfo := make([]entity.ContractTrade, 0)
		for i := range temp.Uris {
			dbAdditionalInfo = append(dbAdditionalInfo, entity.ContractTrade{
				TransactionHash: transactionHash,
				UserAddress:     temp.Tos[i].Hex(),
				TokenId:         int(temp.TokenIds[i].Int64()),
				TokenUri:        temp.Uris[i],
			})
		}
		log.Println(dbAdditionalInfo)

		return transactionHash, dao.ContractTrade.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.ContractTrade.Ctx(ctx).Data(dbAdditionalInfo).Batch(len(dbAdditionalInfo)).Insert()
			return err
		})
	}

	return "", nil
}
