package main

import "github.com/astaxie/beego/migration"

// DO NOT MODIFY
type Stores_20170704_165247 struct {
	migration.Migration
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
	m.Migrate("create")
}

// Reverse the migrations
func (m *Stores_20170704_165247) Down() {
	m.Migrate("delete")
}
