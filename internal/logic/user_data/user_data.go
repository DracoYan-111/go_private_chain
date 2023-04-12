package user_data

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/contracts/accountsFactory"
	"go_private_chain/contracts/box721"
	"go_private_chain/contracts/createBox721"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/deploy"
	"go_private_chain/internal/model/entity"
	"go_private_chain/internal/service"
	"go_private_chain/utility"
	"math/big"
	"math/rand"
	"strconv"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	private := "web3.accountsKey.privateKey" + strconv.Itoa(rand.Intn(5))
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
	ContractAddress common.Address   `json:"contractAddress"`
	UserAddrArray   []common.Address `json:"userAddrArray"`
	TokenIdArray    []string         `json:"tokenIdArray"`
	UriArray        []string         `json:"uriArray"`
}

// BatchCastingNft 新的批量创建nft任务
func (s *sUserData) BatchCastingNft(ctx context.Context, req string) (string, []string, error) {
	// 将解密后的数据转换为结构体数据
	var temp additionalInfo
	utility.DecryptStructure(req, &temp)

	// 检查tokenId唯一性
	if len(temp.UserAddrArray) == len(temp.UriArray) && len(temp.UriArray) == len(temp.TokenIdArray) {
		all, err := dao.ContractTrade.Ctx(ctx).Where("token_id", temp.TokenIdArray).Where("contract_address", temp.ContractAddress.Hex()).All()
		if err != nil {
			return "", nil, err
		} else if len(all) > 0 {
			return "", nil, fmt.Errorf("合约的tokenID已存在")
		}
	}

	// 创建用户合约
	rand.Seed(time.Now().UnixNano())
	private := "web3.accountsKey.privateKey" + strconv.Itoa(rand.Intn(5))
	loading, _ := utility.ReadConfigFile([]string{"web3.createBox721", private})
	createBox := deploy.LoadWithAddress(loading["web3.createBox721"], "createBox721", loading[private]).(*createBox721.CreateBox721)
	transactionHash, err := deploy.BulkIssuance(createBox, temp.ContractAddress, temp.UserAddrArray, temp.TokenIdArray, temp.UriArray)
	if err != nil {
		return "", nil, err
	}

	// 将内容更新到数据库
	dbAdditionalInfo := make([]entity.ContractTrade, 0)
	for i := range temp.UriArray {
		dbAdditionalInfo = append(dbAdditionalInfo, entity.ContractTrade{
			TransactionHash: transactionHash,
			ContractAddress: temp.ContractAddress.Hex(),
			UserAddress:     temp.UserAddrArray[i].Hex(),
			TokenId:         temp.TokenIdArray[i],
			TokenUri:        temp.UriArray[i],
		})
	}

	return transactionHash, temp.TokenIdArray, dao.ContractTrade.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.ContractTrade.Ctx(ctx).Data(dbAdditionalInfo).Batch(len(dbAdditionalInfo)).Insert()
		return err
	})
}

type transferNftInfo struct {
	ContractAddress common.Address   `json:"contractAddress"`
	UserAddress     common.Address   `json:"userAddress"`
	ReceiveAddress  []common.Address `json:"receiveAddrArray"`
	TokenIdArray    []string         `json:"tokenIdArray"`
}

// BatchTransferNft 新的批量转移nft任务
func (s *sUserData) BatchTransferNft(ctx context.Context, req string) (string, []string, error) {
	// 将解密后的数据转换为结构体数据
	var temp transferNftInfo
	utility.DecryptStructure(req, &temp)

	// 检查当前用户余额是否正常
	rand.Seed(time.Now().UnixNano())
	private := "web3.accountsKey.privateKey" + strconv.Itoa(rand.Intn(5))
	loading, _ := utility.ReadConfigFile([]string{"web3.contractCall", "web3.accountsFactory", private})
	box721Contract := deploy.LoadWithAddress(temp.ContractAddress.Hex(), "box721", loading[private]).(*box721.Box721)
	internalId, externalId, _, err := box721Contract.UserAllTokenIndexes(nil, temp.UserAddress)
	if err != nil {
		return "", nil, fmt.Errorf("UserAllTokenIndexes: %s", err)
	} else if len(temp.ReceiveAddress) > len(internalId) {
		return "", nil, fmt.Errorf("转账数量超过余额")
	}

	// 检查用户要转移的id是否拥有所有权
	var userAddress []common.Address
	var correct []*big.Int
	var tokenIdArray []string
	for i, str1 := range temp.TokenIdArray {
		for _, str2 := range externalId {
			if str1 == str2 {
				correct = append(correct, internalId[i])
				userAddress = append(userAddress, temp.UserAddress)
				tokenIdArray = append(tokenIdArray, str1)
			}
		}
	}

	// 进行转移方法
	if len(correct) > 0 {
		transfer, err := deploy.BulkTransfer(userAddress, temp.ReceiveAddress, correct, temp.ContractAddress)
		if err != nil {
			return "", nil, err
		}
		return transfer, tokenIdArray, nil
	}
	return "", nil, nil
}
