package migrations

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Discountables_20170716_090213 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Discountables_20170716_090213{}
	m.Created = "20170716_090213"
	m.ddlSpec()
	migration.Register("Discountables_20170716_090213", m)
}

/*
	refer beego/migration/doc.go
*/
func (m *Discountables_20170716_090213) ddlSpec() {
	m.CreateTable("discountables", "InnoDB", "utf8")
	m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT(10)").SetUnsigned(true)
	m.ForeignCol("discount_id", "id", "discounts").SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("discountable_id").SetNullable(true).SetDataType("INT(10)").SetUnsigned(true).SetDefault("NULL")
	m.NewCol("discountable_type").SetDataType("VARCHAR(255)").SetNullable(false)
	m.NewCol("created_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")
	m.NewCol("updated_at").SetNullable(true).SetDataType("DATETIME").SetDefault("NULL")

}
