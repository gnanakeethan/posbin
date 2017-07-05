package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/migration"
	"github.com/gnanakeethan/posbin/database/migrations/migrator"
)

// DO NOT MODIFY
type Stores_20170704_165247 struct {
	migration.Migration
	migrator.Migrate
}

// DO NOT MODIFY
func init() {
	m := &Stores_20170704_165247{}
	m.Created = "20170704_165247"
	migration.Register("Stores_20170704_165247", m)
	m.TableName = "stores"
	m.Engine = "InnoDB"
	m.Charset = "utf8"
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.PriCol("sec_id").SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.UniCol("name_unique", "name").SetDataType("VARCHAR(255)").SetNullable(true)
	m.UniCol("name_unique", "cals").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("address").SetDataType("text").SetNullable(true)
	m.NewCol("contact").SetDataType("tinytext").SetNullable(true)
}

// Run the migrations
func (m *Stores_20170704_165247) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.ModifyType = "create"
	sql := m.GetSQL()
	beego.BeeLogger.Info(fmt.Sprint(sql))
	m.SQL(sql)
}

// Reverse the migrations
func (m *Stores_20170704_165247) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.ModifyType = "delete"
	sql := m.GetSQL()
	beego.BeeLogger.Info(fmt.Sprint(sql))
	m.SQL(sql)
}
