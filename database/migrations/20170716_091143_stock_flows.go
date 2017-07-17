package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type StockFlows_20170716_091143 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &StockFlows_20170716_091143{}
	m.Created = "20170716_091143"
	m.ddlSpec()
	migration.Register("StockFlows_20170716_091143", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *StockFlows_20170716_091143) ddlSpec() {
	m.CreateTable("stock_flows", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("stock_id", "id", "stocks").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("flow").SetDataType("DOUBLE(4,2)").SetNullable(false).SetDefault("1")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
