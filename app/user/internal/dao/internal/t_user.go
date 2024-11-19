// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDao is the data access object for table t_user.
type TUserDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TUserColumns // columns contains all the column names of Table for convenient usage.
}

// TUserColumns defines and stores column names for table t_user.
type TUserColumns struct {
	Id     string // user id
	Name   string // user name
	Status string // user status
	Age    string // user age
}

// tUserColumns holds the columns for table t_user.
var tUserColumns = TUserColumns{
	Id:     "id",
	Name:   "name",
	Status: "status",
	Age:    "age",
}

// NewTUserDao creates and returns a new DAO object for table data access.
func NewTUserDao() *TUserDao {
	return &TUserDao{
		group:   "default",
		table:   "t_user",
		columns: tUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserDao) Columns() TUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
