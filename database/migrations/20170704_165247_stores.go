package main

import (
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
	m.TableName = "stores"
	m.Engine = "InnoDB"
	m.Charset = "utf8"
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.UniCol("name_unique", "name").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("address").SetDataType("text").SetNullable(true)
	m.NewCol("contact").SetDataType("tinytext").SetNullable(true)
	migration.Register("Stores_20170704_165247", m)
}

// Run the migrations
func (m *Stores_20170704_165247) Up() {
	m.ModifyType = "create"
	sql := m.GetSQL()
	m.SQL(sql)
}

// Reverse the migrations
func (m *Stores_20170704_165247) Down() {
	m.ModifyType = "delete"
	sql := m.GetSQL()
	m.SQL(sql)
}
