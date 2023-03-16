package dao

import (
	"go_private_chain/internal/dao/internal"
)

// internalGoTestDbDao is internal type for wrapping internal DAO implements.
type internalGoTestDbDao = *internal.GoTestDbDao

// goTestDbDao is the data access object for table go_test_db.
// You can define custom methods on it to extend its functionality as you wish.
type goTestDbDao struct {
	internalGoTestDbDao
}

var (
	// GoTestDb is globally public accessible object for table go_test_db operations.
	GoTestDb = goTestDbDao{
		internal.NewGoTestDbDao(),
	}
)

// Fill with you ideas below.
// Model is core struct implementing the DAO for ORM.
