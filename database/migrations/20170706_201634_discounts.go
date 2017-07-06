package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Discounts_20170706_201634 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Discounts_20170706_201634{}
	m.Created = "20170706_201634"
	m.TableName = "discounts"
	m.Engine = "InnoDB"
	m.Charset = "utf8"
	m.ForeignCol("store_id", "id", "stores").SetDataType("int(10)").SetUnsigned(true).SetDefault("NULL")
	migration.Register("Discounts_20170706_201634", m)
}

// Run the migrations
func (m *Discounts_20170706_201634) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.Migrate("alter")
}

// Reverse the migrations
func (m *Discounts_20170706_201634) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.Migrate("reverse")
}
