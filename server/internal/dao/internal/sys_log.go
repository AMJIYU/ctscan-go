// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLogDao is the data access object for the table hg_sys_log.
type SysLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysLogColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysLogColumns defines and stores column names for the table hg_sys_log.
type SysLogColumns struct {
	Id         string // 日志ID
	ReqId      string // 对外ID
	AppId      string // 应用ID
	MerchantId string // 商户ID
	MemberId   string // 用户ID
	Method     string // 提交类型
	Module     string // 访问模块
	Url        string // 提交url
	GetData    string // get数据
	PostData   string // post数据
	HeaderData string // header数据
	Ip         string // IP地址
	ProvinceId string // 省编码
	CityId     string // 市编码
	ErrorCode  string // 报错code
	ErrorMsg   string // 对外错误提示
	ErrorData  string // 报错日志
	UserAgent  string // UA信息
	TakeUpTime string // 请求耗时
	Timestamp  string // 响应时间
	Status     string // 状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// sysLogColumns holds the columns for the table hg_sys_log.
var sysLogColumns = SysLogColumns{
	Id:         "id",
	ReqId:      "req_id",
	AppId:      "app_id",
	MerchantId: "merchant_id",
	MemberId:   "member_id",
	Method:     "method",
	Module:     "module",
	Url:        "url",
	GetData:    "get_data",
	PostData:   "post_data",
	HeaderData: "header_data",
	Ip:         "ip",
	ProvinceId: "province_id",
	CityId:     "city_id",
	ErrorCode:  "error_code",
	ErrorMsg:   "error_msg",
	ErrorData:  "error_data",
	UserAgent:  "user_agent",
	TakeUpTime: "take_up_time",
	Timestamp:  "timestamp",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSysLogDao creates and returns a new DAO object for table data access.
func NewSysLogDao(handlers ...gdb.ModelHandler) *SysLogDao {
	return &SysLogDao{
		group:    "default",
		table:    "hg_sys_log",
		columns:  sysLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLogDao) Columns() SysLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
