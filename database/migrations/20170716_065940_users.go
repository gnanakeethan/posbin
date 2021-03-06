package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20170716_065940 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20170716_065940{}
	m.Created = "20170716_065940"
	m.ddlSpec()
	migration.Register("Users_20170716_065940", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Users_20170716_065940) ddlSpec() {
	m.CreateTable("users", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.UniCol("username", "username").SetNullable(true).SetDataType("VARCHAR(255)").SetDefault("NULL")
	m.NewCol("email").SetNullable(true).SetDataType("VARCHAR(255)").SetDefault("NULL")
	m.NewCol("password").SetNullable(true).SetDataType("VARCHAR(255)").SetDefault("NULL")
	m.NewCol("remember_token").SetNullable(true).SetDataType("VARCHAR(255)").SetDefault("NULL")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
