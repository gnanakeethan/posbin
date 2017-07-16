package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InventoryScale_20170716_091044 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InventoryScale_20170716_091044{}
	m.Created = "20170716_091044"
	m.ddlSpec()
	migration.Register("InventoryScale_20170716_091044", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *InventoryScale_20170716_091044) ddlSpec() {
	m.CreateTable("inventory_scale", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("inventory_id", "id", "inventories").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.ForeignCol("scale_id", "id", "scales").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("units").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("price").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")

}
