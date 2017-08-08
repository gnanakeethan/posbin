package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Products_20170716_085355 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Products_20170716_085355{}
	m.Created = "20170716_085355"
	m.ddlSpec()
	migration.Register("Products_20170716_085355", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Products_20170716_085355) ddlSpec() {
	m.CreateTable("products", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("product_code").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("barcode").SetDataType("VARCHAR(255)").SetNullable(true)
	m.NewCol("name").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("singular").SetDataType("tinyint(1)").SetDefault("NULL")
	m.NewCol("priority").SetDataType("tinyint(1)").SetDefault("NULL")
	m.NewCol("ordering").SetDataType("int(10)").SetNullable(false).SetUnsigned(true)
	m.NewCol("image").SetDataType("BLOB").SetDefault("NULL")
	m.NewCol("image_type").SetDataType("VARCHAR(20)").SetDefault("NULL")
	m.ForeignCol("scale_id", "id", "scales").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.ForeignCol("store_id", "id", "stores").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("service").SetDataType("TINYINT(1)").SetNullable(false).SetDefault("0")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
