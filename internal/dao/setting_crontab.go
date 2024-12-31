// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSettingCrontabDao is internal type for wrapping internal DAO implements.
type internalSettingCrontabDao = *internal.SettingCrontabDao

// settingCrontabDao is the data access object for table setting_crontab.
// You can define custom methods on it to extend its functionality as you wish.
type settingCrontabDao struct {
	internalSettingCrontabDao
}

var (
	// SettingCrontab is globally public accessible object for table setting_crontab operations.
	SettingCrontab = settingCrontabDao{
		internal.NewSettingCrontabDao(),
	}
)

// Fill with you ideas below.