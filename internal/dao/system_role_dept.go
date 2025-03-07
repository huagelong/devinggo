// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemRoleDeptDao is internal type for wrapping internal DAO implements.
type internalSystemRoleDeptDao = *internal.SystemRoleDeptDao

// systemRoleDeptDao is the data access object for table system_role_dept.
// You can define custom methods on it to extend its functionality as you wish.
type systemRoleDeptDao struct {
	internalSystemRoleDeptDao
}

var (
	// SystemRoleDept is globally public accessible object for table system_role_dept operations.
	SystemRoleDept = systemRoleDeptDao{
		internal.NewSystemRoleDeptDao(),
	}
)

// Fill with you ideas below.
