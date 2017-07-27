package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Roles_20170716_071411 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Roles_20170716_071411{}
	m.Created = "20170716_071411"
	m.ddlSpec()
	migration.Register("Roles_20170716_071411", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Roles_20170716_071411) ddlSpec() {
	m.CreateTable("roles", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetNullable(false).SetDataType("VARCHAR(255)")
	m.NewCol("display_name").SetNullable(true).SetDataType("VARCHAR(255)")
	m.NewCol("description").SetNullable(true).SetDataType("VARCHAR(255)")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
}
