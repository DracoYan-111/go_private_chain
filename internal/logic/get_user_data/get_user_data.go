package get_user_data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/model/entity"
	"go_private_chain/internal/service"
)

type (
	sGetUserData struct{}
)

func init() {
	service.RegisterGetUserData(New())
}

func New() service.IGetUserData {
	return &sGetUserData{}
}

// GetUserAddress 查询用户地址
func (s *sGetUserData) GetUserAddress(ctx context.Context, req string) (resData []string, err error) {

	// 通过昵称查询用户地址
	userData, err := dao.UserData.Ctx(ctx).One("user_nick", req)
	if err != nil {
		return nil, fmt.Errorf("查询用户地址失败(GetUserAddress):%s", err)
	}

	// 封装返回值
	if len(userData) != 0 {
		var data *entity.UserData
		err = structure(userData.Json(), &data)
		resData = append(resData, data.UserAddress, data.AccountHash)
		return resData, nil
	}

	return nil, fmt.Errorf("查询用户地址失败(GetUserAddress):用户不存在")
}

// GetUserCollection 查询用户已有藏品
func (s *sGetUserData) GetUserCollection(ctx context.Context, req string) (userCollectionList []entity.UserCollections, err error) {

	// 通过昵称查询用户藏品
	userCollectionArray, err := dao.ContractTrade.DB().GetAll(ctx,
		"SELECT B.* FROM user_data AS A JOIN contract_trade AS B ON A.user_address = B.user_address WHERE A.user_nick = ?", req)
	if err != nil {
		return nil, fmt.Errorf("查询用户藏品失败(GetUserCollection):%s", err)
	}

	// 检查数据库返回值
	if len(userCollectionArray) != 0 {

		// 信息格式化
		var collectionList []entity.ContractTrade
		err := structure(userCollectionArray.Json(), &collectionList)
		if err != nil {
			return nil, fmt.Errorf("信息格式化失败(GetUserCollection):%s", err)
		}
		collections := make(map[string]*entity.UserCollections)

		for _, v := range collectionList {
			if collection, ok := collections[v.ContractAddress]; ok {
				collection.CollectionList.TransactionHash = append(collection.CollectionList.TransactionHash, v.TransactionHash)
				collection.CollectionList.TokenId = append(collection.CollectionList.TokenId, v.TokenId)
				collection.CollectionList.TokenUri = append(collection.CollectionList.TokenUri, v.TokenUri)
			} else {
				collections[v.ContractAddress] = &entity.UserCollections{
					ContractAddress: v.ContractAddress,
					CollectionList: struct {
						TransactionHash []string `json:"TransactionHash"`
						TokenId         []string `json:"TokenId"`
						TokenUri        []string `json:"TokenUri"`
					}{
						TransactionHash: []string{v.TransactionHash},
						TokenId:         []string{v.TokenId},
						TokenUri:        []string{v.TokenUri},
					},
				}
			}
		}
		for _, collection := range collections {
			userCollectionList = append(userCollectionList, *collection)
		}
		return userCollectionList, nil
	}
	return nil, nil
}

func structure(req string, structure interface{}) error {
	// 将解密后的数据转换为结构体数据
	err := json.Unmarshal([]byte(req), &structure)
	if err != nil {
		return errors.New("转换为结构体失败")
	}
	return nil
}
