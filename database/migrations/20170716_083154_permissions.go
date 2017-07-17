package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Permissions_20170716_083154 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Permissions_20170716_083154{}
	m.Created = "20170716_083154"
	m.ddlSpec()
	migration.Register("Permissions_20170716_083154", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Permissions_20170716_083154) ddlSpec() {
	m.CreateTable("permissions", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.NewCol("name").SetNullable(false).SetDataType("VARCHAR(255)")
	m.NewCol("display_name").SetNullable(true).SetDataType("VARCHAR(255)")
	m.NewCol("description").SetNullable(true).SetDataType("VARCHAR(255)")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
