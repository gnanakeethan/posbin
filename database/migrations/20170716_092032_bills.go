package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Bills_20170716_092032 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Bills_20170716_092032{}
	m.Created = "20170716_092032"
	m.ddlSpec()
	migration.Register("Bills_20170716_092032", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Bills_20170716_092032) ddlSpec() {
	m.CreateTable("bills", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("customer_id", "id", "customers").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.NewCol("cost").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("gross_total").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("discount").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("net_total").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("balance").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("card_paid").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.NewCol("cash_paid").SetDataType("DOUBLE(16,4)").SetNullable(false).SetDefault("1")
	m.ForeignCol("user_id", "id", "users").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.ForeignCol("terminal_id", "id", "terminals").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.ForeignCol("store_id", "id", "stores").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("closed").SetDataType("TINYINT(1)").SetNullable(false)
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("deleted_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
}
