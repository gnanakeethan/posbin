package database

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/ddds")
}

func Migrate() {
	//migration.Refresh()
}
