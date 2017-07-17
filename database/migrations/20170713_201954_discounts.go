package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Discounts_20170713_201954 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Discounts_20170713_201954{}
	m.Created = "20170713_201954"
	m.ddlSpec()
	migration.Register("Discounts_20170713_201954", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Discounts_20170713_201954) ddlSpec() {
	m.CreateTable("discounts", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetDataType("VARCHAR(255) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("from").SetDataType("datetime").SetNullable(false)
	m.NewCol("to").SetDataType("datetime").SetNullable(false)
	m.NewCol("type").SetDataType("enum('value','percentage') COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("value").SetDataType("DOUBLE(8,2)").SetNullable(false)
	m.ForeignCol("store_id", "id", "stores").SetDataType("int(10)").SetUnsigned(true).SetDefault("NULL")

}
