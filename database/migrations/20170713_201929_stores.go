package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Stores_20170713_201929 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Stores_20170713_201929{}
	m.Created = "20170713_201929"
	m.ddlSpec()
	migration.Register("Stores_20170713_201929", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Stores_20170713_201929) ddlSpec() {
	m.CreateTable("stores", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.UniCol("name_unique", "name").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("address").SetDataType("text").SetNullable(true).SetDefault("NULL")
	m.NewCol("contact").SetDataType("tinytext").SetNullable(true).SetDefault("NULL")
}
