package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Inventories_20170716_090814 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Inventories_20170716_090814{}
	m.Created = "20170716_090814"
	m.ddlSpec()
	migration.Register("Inventories_20170716_090814", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Inventories_20170716_090814) ddlSpec() {
	m.CreateTable("inventories", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("product_id", "id", "products").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("batch_no").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("service").SetDataType("TINYINT(1)").SetNullable(false).SetDefault("0")

}
