// db/init.go
package db

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() {
	// username:password@tcp(host:port)/dbname?charset=utf8
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=true"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RunSyncdb("default", false, true) // auto-creates tables (dev only).
}
