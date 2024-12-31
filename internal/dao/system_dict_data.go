// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemDictDataDao is internal type for wrapping internal DAO implements.
type internalSystemDictDataDao = *internal.SystemDictDataDao

// systemDictDataDao is the data access object for table system_dict_data.
// You can define custom methods on it to extend its functionality as you wish.
type systemDictDataDao struct {
	internalSystemDictDataDao
}

var (
	// SystemDictData is globally public accessible object for table system_dict_data operations.
	SystemDictData = systemDictDataDao{
		internal.NewSystemDictDataDao(),
	}
)

// Fill with you ideas below.