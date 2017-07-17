package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Tokens_20170716_091716 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Tokens_20170716_091716{}
	m.Created = "20170716_091716"
	m.ddlSpec()
	migration.Register("Tokens_20170716_091716", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Tokens_20170716_091716) ddlSpec() {
	m.CreateTable("tokens", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("user_id", "id", "users").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	m.NewCol("key").SetDataType("MEDIUMTEXT").SetNullable(false)
	m.NewCol("valid_thru").SetDataType("DATETIME").SetNullable(false)
	m.NewCol("valid").SetDataType("TINYINT(1)").SetNullable(false)

}
