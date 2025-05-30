// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// sysProvincesDao is the data access object for the table hg_sys_provinces.
// You can define custom methods on it to extend its functionality as needed.
type sysProvincesDao struct {
	*internal.SysProvincesDao
}

var (
	// SysProvinces is a globally accessible object for table hg_sys_provinces operations.
	SysProvinces = sysProvincesDao{internal.NewSysProvincesDao()}
)

// Add your custom methods and functionality below.
