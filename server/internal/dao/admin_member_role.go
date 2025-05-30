// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// adminMemberRoleDao is the data access object for the table hg_admin_member_role.
// You can define custom methods on it to extend its functionality as needed.
type adminMemberRoleDao struct {
	*internal.AdminMemberRoleDao
}

var (
	// AdminMemberRole is a globally accessible object for table hg_admin_member_role operations.
	AdminMemberRole = adminMemberRoleDao{internal.NewAdminMemberRoleDao()}
)

// Add your custom methods and functionality below.
