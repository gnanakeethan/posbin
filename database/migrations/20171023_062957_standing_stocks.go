package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type StandingStocks_20171023_062957 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &StandingStocks_20171023_062957{}
	m.Created = "20171023_062957"
	m.ddlSpec()
	migration.Register("StandingStocks_20171023_062957", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *StandingStocks_20171023_062957) ddlSpec() {
	m.CreateTable("standing_stocks", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("quantity").SetDataType("double").SetUnsigned(true)
	m.ForeignCol("inventory_id", "id", "inventory").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)

}
