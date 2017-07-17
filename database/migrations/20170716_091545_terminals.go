package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Terminals_20170716_091545 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Terminals_20170716_091545{}
	m.Created = "20170716_091545"
	m.ddlSpec()
	migration.Register("Terminals_20170716_091545", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Terminals_20170716_091545) ddlSpec() {
	m.CreateTable("terminals", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("user_id", "id", "users").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.NewCol("machine").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("printer").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("balance").SetDataType("DOUBLE(24,8)").SetNullable(false).SetDefault("1")
	m.ForeignCol("store_id", "id", "stores").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
