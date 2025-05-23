// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSettingGenerateColumnsDao is internal type for wrapping internal DAO implements.
type internalSettingGenerateColumnsDao = *internal.SettingGenerateColumnsDao

// settingGenerateColumnsDao is the data access object for table setting_generate_columns.
// You can define custom methods on it to extend its functionality as you wish.
type settingGenerateColumnsDao struct {
	internalSettingGenerateColumnsDao
}

var (
	// SettingGenerateColumns is globally public accessible object for table setting_generate_columns operations.
	SettingGenerateColumns = settingGenerateColumnsDao{
		internal.NewSettingGenerateColumnsDao(),
	}
)

// Fill with you ideas below.
