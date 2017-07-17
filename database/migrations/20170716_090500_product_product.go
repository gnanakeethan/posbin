package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ProductProduct_20170716_090500 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ProductProduct_20170716_090500{}
	m.Created = "20170716_090500"
	m.ddlSpec()
	migration.Register("ProductProduct_20170716_090500", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *ProductProduct_20170716_090500) ddlSpec() {
	m.CreateTable("product_product", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	parent := m.ForeignCol("parent_id", "id", "products").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	child := m.ForeignCol("child_id", "id", "products").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")

	m.NewCol("units").SetDataType("DOUBLE(8,2)").SetNullable(false).SetDefault("1")
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

	//Set Unique
	unique := migration.Unique{Definition: "parent_child"}
	unique.AddColumnsToUnique(parent, child)
	m.AddUnique(&unique)

}
