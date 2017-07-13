package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Products_20170713_202036 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Products_20170713_202036{}
	m.Created = "20170713_202036"
	m.ddlSpec()
	migration.Register("Products_20170713_202036", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Products_20170713_202036) ddlSpec() {
	m.CreateTable("products", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)

}
