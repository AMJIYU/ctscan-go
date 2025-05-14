// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AedLocationsDao is the data access object for the table aed_locations.
type AedLocationsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AedLocationsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AedLocationsColumns defines and stores column names for the table aed_locations.
type AedLocationsColumns struct {
	Id        string //
	Longitude string // 经度
	Latitude  string // 纬度
	Address   string // 地址
	CreatedAt string //
	UpdatedAt string //
}

// aedLocationsColumns holds the columns for the table aed_locations.
var aedLocationsColumns = AedLocationsColumns{
	Id:        "id",
	Longitude: "longitude",
	Latitude:  "latitude",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAedLocationsDao creates and returns a new DAO object for table data access.
func NewAedLocationsDao(handlers ...gdb.ModelHandler) *AedLocationsDao {
	return &AedLocationsDao{
		group:    "aed",
		table:    "aed_locations",
		columns:  aedLocationsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AedLocationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AedLocationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AedLocationsDao) Columns() AedLocationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AedLocationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AedLocationsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AedLocationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
