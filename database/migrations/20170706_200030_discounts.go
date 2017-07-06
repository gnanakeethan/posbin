package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Discounts_20170706_200030 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Discounts_20170706_200030{}
	m.Created = "20170706_200030"
	m.TableName = "discounts"
	m.Engine = "InnoDB"
	m.Charset = "utf8"
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(255) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("from").SetDataType("datetime").SetNullable(false)
	m.NewCol("to").SetDataType("datetime").SetNullable(false)
	m.NewCol("type").SetDataType("enum('value','percentage') COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("value").SetDataType("DOUBLE(8,2)").SetNullable(false)
	migration.Register("Discounts_20170706_200030", m)
}

// Run the migrations
func (m *Discounts_20170706_200030) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.Migrate("create")

}

// Reverse the migrations
func (m *Discounts_20170706_200030) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.Migrate("delete")

}
