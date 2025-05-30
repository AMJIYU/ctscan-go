// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysEmsLogDao is the data access object for the table hg_sys_ems_log.
type SysEmsLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysEmsLogColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysEmsLogColumns defines and stores column names for the table hg_sys_ems_log.
type SysEmsLogColumns struct {
	Id        string // 主键
	Event     string // 事件
	Email     string // 邮箱地址，多个用;隔开
	Code      string // 验证码
	Times     string // 验证次数
	Content   string // 邮件内容
	Ip        string // ip地址
	Status    string // 状态(1未验证,2已验证)
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysEmsLogColumns holds the columns for the table hg_sys_ems_log.
var sysEmsLogColumns = SysEmsLogColumns{
	Id:        "id",
	Event:     "event",
	Email:     "email",
	Code:      "code",
	Times:     "times",
	Content:   "content",
	Ip:        "ip",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysEmsLogDao creates and returns a new DAO object for table data access.
func NewSysEmsLogDao(handlers ...gdb.ModelHandler) *SysEmsLogDao {
	return &SysEmsLogDao{
		group:    "default",
		table:    "hg_sys_ems_log",
		columns:  sysEmsLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysEmsLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysEmsLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysEmsLogDao) Columns() SysEmsLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysEmsLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysEmsLogDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysEmsLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
