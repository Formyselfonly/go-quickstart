package dao

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// User is the data access object for user operations
var User = userDao{}

type userDao struct{}

// Table returns the table name for user
func (dao *userDao) Table() string {
	return "users"
}

// DB returns the database connection
func (dao *userDao) DB() gdb.DB {
	return g.DB()
}

// Ctx returns a new context with the database connection
func (dao *userDao) Ctx(ctx g.Ctx) *gdb.Model {
	return dao.DB().Model(dao.Table()).Safe().Ctx(ctx)
}
