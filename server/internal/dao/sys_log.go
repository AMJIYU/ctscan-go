// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// sysLogDao is the data access object for the table hg_sys_log.
// You can define custom methods on it to extend its functionality as needed.
type sysLogDao struct {
	*internal.SysLogDao
}

var (
	// SysLog is a globally accessible object for table hg_sys_log operations.
	SysLog = sysLogDao{internal.NewSysLogDao()}
)

// Add your custom methods and functionality below.
