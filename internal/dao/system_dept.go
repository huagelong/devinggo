// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemDeptDao is internal type for wrapping internal DAO implements.
type internalSystemDeptDao = *internal.SystemDeptDao

// systemDeptDao is the data access object for table system_dept.
// You can define custom methods on it to extend its functionality as you wish.
type systemDeptDao struct {
	internalSystemDeptDao
}

var (
	// SystemDept is globally public accessible object for table system_dept operations.
	SystemDept = systemDeptDao{
		internal.NewSystemDeptDao(),
	}
)

// Fill with you ideas below.
