package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Scales_20170716_084718 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Scales_20170716_084718{}
	m.Created = "20170716_084718"
	m.ddlSpec()
	migration.Register("Scales_20170716_084718", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Scales_20170716_084718) ddlSpec() {
	m.CreateTable("scales", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("description").SetNullable(true).SetDataType("VARCHAR(255)")
	m.NewCol("unit").SetNullable(true).SetDataType("VARCHAR(10)")

	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.ForeignCol("store_id", "id", "stores").SetDataType("int(10)").SetUnsigned(true).SetDefault("NULL")

}
