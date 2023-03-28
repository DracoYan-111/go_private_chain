package dao

import (
	"go_private_chain/internal/dao/internal"
)

// internalUserDataDao is internal type for wrapping internal DAO implements.
type internalUserDataDao = *internal.UserDataDao

// userDataDao is the data access object for table user_data
// You can define custom methods on it to extend its functionality as you wish.
type userDataDao struct {
	internalUserDataDao
}

var (
	// UserData is globally public accessible object for table user_data operations.
	UserData = userDataDao{
		internal.NewUserDataDao(),
	}
)

// Fill with you ideas below.
