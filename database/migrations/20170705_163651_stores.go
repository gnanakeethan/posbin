package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/migration"
	"github.com/gnanakeethan/posbin/database/migrations/migrator"
)

// DO NOT MODIFY
type Stores_20170705_163651 struct {
	migration.Migration
	migrator.Migrate
}

// DO NOT MODIFY
func init() {
	m := &Stores_20170705_163651{}
	m.Created = "20170705_163651"
	migration.Register("Stores_20170705_163651", m)

	m.TableName = "stores"
	m.NewCol("secd_id").SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	rename := m.RenameColumn("sec_id", "ter_id").SetOldNullable(false).SetOldDataType("INT(10)").SetOldUnsigned(true)
	rename.SetNullable(true).SetDataType("INT(11)").SetUnsigned(false)

}

// Run the migrations
func (m *Stores_20170705_163651) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.ModifyType = "alter"
	sql := m.GetSQL()
	m.SQL(sql)
	beego.BeeLogger.Info(fmt.Sprint(sql))
}

// Reverse the migrations
func (m *Stores_20170705_163651) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.ModifyType = "reverse"
	sql := m.GetSQL()
	m.SQL(sql)
	beego.BeeLogger.Info(fmt.Sprint(sql))
}
