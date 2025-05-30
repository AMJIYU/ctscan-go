// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayLogDao is the data access object for the table hg_pay_log.
type PayLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PayLogColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PayLogColumns defines and stores column names for the table hg_pay_log.
type PayLogColumns struct {
	Id            string // 主键
	MemberId      string // 会员ID
	AppId         string // 应用ID
	AddonsName    string // 插件名称
	OrderSn       string // 关联订单号
	OrderGroup    string // 组别[默认统一支付类型]
	Openid        string // openid
	MchId         string // 商户支付账户
	Subject       string // 订单标题
	Detail        string // 支付商品详情
	AuthCode      string // 刷卡码
	OutTradeNo    string // 商户订单号
	TransactionId string // 交易号
	PayType       string // 支付类型
	PayAmount     string // 支付金额
	ActualAmount  string // 实付金额
	PayStatus     string // 支付状态
	PayAt         string // 支付时间
	TradeType     string // 交易类型
	RefundSn      string // 退款单号
	IsRefund      string // 是否退款
	Custom        string // 自定义参数
	CreateIp      string // 创建者IP
	PayIp         string // 支付者IP
	NotifyUrl     string // 支付通知回调地址
	ReturnUrl     string // 买家付款成功跳转地址
	TraceIds      string // 链路ID集合
	Status        string // 状态
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
}

// payLogColumns holds the columns for the table hg_pay_log.
var payLogColumns = PayLogColumns{
	Id:            "id",
	MemberId:      "member_id",
	AppId:         "app_id",
	AddonsName:    "addons_name",
	OrderSn:       "order_sn",
	OrderGroup:    "order_group",
	Openid:        "openid",
	MchId:         "mch_id",
	Subject:       "subject",
	Detail:        "detail",
	AuthCode:      "auth_code",
	OutTradeNo:    "out_trade_no",
	TransactionId: "transaction_id",
	PayType:       "pay_type",
	PayAmount:     "pay_amount",
	ActualAmount:  "actual_amount",
	PayStatus:     "pay_status",
	PayAt:         "pay_at",
	TradeType:     "trade_type",
	RefundSn:      "refund_sn",
	IsRefund:      "is_refund",
	Custom:        "custom",
	CreateIp:      "create_ip",
	PayIp:         "pay_ip",
	NotifyUrl:     "notify_url",
	ReturnUrl:     "return_url",
	TraceIds:      "trace_ids",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewPayLogDao creates and returns a new DAO object for table data access.
func NewPayLogDao(handlers ...gdb.ModelHandler) *PayLogDao {
	return &PayLogDao{
		group:    "default",
		table:    "hg_pay_log",
		columns:  payLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PayLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PayLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PayLogDao) Columns() PayLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PayLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PayLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PayLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
