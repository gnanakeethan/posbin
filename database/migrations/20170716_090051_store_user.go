package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type StoreUser_20170716_090051 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &StoreUser_20170716_090051{}
	m.Created = "20170716_090051"
	m.ddlSpec()
	migration.Register("StoreUser_20170716_090051", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *StoreUser_20170716_090051) ddlSpec() {
	m.CreateTable("store_user", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	store := m.ForeignCol("store_id", "id", "stores").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	user := m.ForeignCol("user_id", "id", "users").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)

	//Set Unique
	unique := migration.Unique{Definition: "store_user"}
	unique.AddColumnsToUnique(store, user)
	m.AddUnique(&unique)

}
