package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Stocks_20170716_090934 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Stocks_20170716_090934{}
	m.Created = "20170716_090934"
	m.ddlSpec()
	migration.Register("Stocks_20170716_090934", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Stocks_20170716_090934) ddlSpec() {
	m.CreateTable("stocks", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("inventory_id", "id", "inventories").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("available_stock").SetDataType("DOUBLE(8,2)").SetNullable(false).SetDefault("1")

}
