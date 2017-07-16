package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Sales_20170716_092314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Sales_20170716_092314{}
	m.Created = "20170716_092314"
	m.ddlSpec()
	migration.Register("Sales_20170716_092314", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Sales_20170716_092314) ddlSpec() {
	m.CreateTable("sales", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("bill_id", "id", "bills").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.ForeignCol("inventory_id", "id", "inventories").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.NewCol("units").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("cost").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("hide").SetDataType("TINYINT(1)").SetNullable(false)
	m.NewCol("unit_price").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("total").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("discount").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("amount").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
