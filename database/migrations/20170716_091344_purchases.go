package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Purchases_20170716_091344 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Purchases_20170716_091344{}
	m.Created = "20170716_091344"
	m.ddlSpec()
	migration.Register("Purchases_20170716_091344", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Purchases_20170716_091344) ddlSpec() {
	m.CreateTable("purchases", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("inventory_id", "id", "inventories").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("average_cost").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.NewCol("units").SetDataType("DOUBLE(12,2)").SetNullable(false).SetDefault("1")
	m.ForeignCol("stock_flow_id", "id", "stock_flows").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.ForeignCol("store_id", "id", "stores").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
