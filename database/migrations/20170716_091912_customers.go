package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Customers_20170716_091912 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Customers_20170716_091912{}
	m.Created = "20170716_091912"
	m.ddlSpec()
	migration.Register("Customers_20170716_091912", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Customers_20170716_091912) ddlSpec() {
	m.CreateTable("customers", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(255)")
	m.NewCol("contact_no").SetDataType("VARCHAR(255)")
	m.NewCol("address").SetDataType("VARCHAR(255)")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
}
