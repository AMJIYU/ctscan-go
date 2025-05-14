// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserWxminiDao is the data access object for the table user_wxmini.
type UserWxminiDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserWxminiColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserWxminiColumns defines and stores column names for the table user_wxmini.
type UserWxminiColumns struct {
	Id         string //
	Openid     string // 微信openid
	Unionid    string // 微信unionid
	Nickname   string // 微信昵称
	Avatar     string // 头像
	UserId     string // 关联的用户ID
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Sessionkey string // 微信sessionkey
}

// userWxminiColumns holds the columns for the table user_wxmini.
var userWxminiColumns = UserWxminiColumns{
	Id:         "id",
	Openid:     "openid",
	Unionid:    "unionid",
	Nickname:   "nickname",
	Avatar:     "avatar",
	UserId:     "user_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Sessionkey: "sessionkey",
}

// NewUserWxminiDao creates and returns a new DAO object for table data access.
func NewUserWxminiDao(handlers ...gdb.ModelHandler) *UserWxminiDao {
	return &UserWxminiDao{
		group:    "aed",
		table:    "user_wxmini",
		columns:  userWxminiColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserWxminiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserWxminiDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserWxminiDao) Columns() UserWxminiColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserWxminiDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserWxminiDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserWxminiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
