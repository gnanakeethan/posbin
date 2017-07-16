package main

import "github.com/astaxie/beego/migration"

// DO NOT MODIFY
type PermissionRole_20170716_083244 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PermissionRole_20170716_083244{}
	m.Created = "20170716_083244"
	m.ddlSpec()
	migration.Register("PermissionRole_20170716_083244", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *PermissionRole_20170716_083244) ddlSpec() {
	m.CreateTable("permission_role", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	role := m.ForeignCol("role_id", "id", "roles").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)
	permission := m.ForeignCol("permission_id", "id", "permissions").SetDataType("INT(10)").SetNullable(true).SetUnsigned(true)

	//Set Unique
	unique := migration.Unique{Definition: "role_permission"}
	unique.AddColumnsToUnique(role, permission)
	m.AddUnique(&unique)

}
