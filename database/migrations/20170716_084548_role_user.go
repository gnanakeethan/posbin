package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RoleUser_20170716_084548 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RoleUser_20170716_084548{}
	m.Created = "20170716_084548"
	m.ddlSpec()
	migration.Register("RoleUser_20170716_084548", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *RoleUser_20170716_084548) ddlSpec() {
	m.CreateTable("role_user", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	role := m.ForeignCol("role_id", "id", "roles").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	user := m.ForeignCol("user_id", "id", "users").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)

	//Set Unique
	unique := migration.Unique{Definition: "role_user"}
	unique.AddColumnsToUnique(role, user)
	m.AddUnique(&unique)

}
