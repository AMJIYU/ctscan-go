// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// adminCashDao is the data access object for the table hg_admin_cash.
// You can define custom methods on it to extend its functionality as needed.
type adminCashDao struct {
	*internal.AdminCashDao
}

var (
	// AdminCash is a globally accessible object for table hg_admin_cash operations.
	AdminCash = adminCashDao{internal.NewAdminCashDao()}
)

// Add your custom methods and functionality below.
