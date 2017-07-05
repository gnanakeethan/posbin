package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/migration"
	"github.com/gnanakeethan/posbin/database/migrations/migrator"
)

// DO NOT MODIFY
type Stores_20170704_165247 struct {
	migration.Migration
	migrator.Migrate
}

// DO NOT MODIFY
func init() {
	m := &Stores_20170704_165247{}
	m.Created = "20170704_165247"
	migration.Register("Stores_20170704_165247", m)
	m.getMigration()
}

// Run the migrations
func (m *Stores_20170704_165247) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.ModifyType = "create"
	sql := m.GetSQL()
	beego.BeeLogger.Info(fmt.Sprint(sql))
	m.SQL(sql)
}

func (m *Stores_20170704_165247) getMigration() {
	m.TableName = "stores"
	m.Engine = "InnoDB"
	m.Charset = "utf8"

	id := &migrator.Column{Name:"id",Unsign:"UNSIGNED",Inc:"AUTO_INCREMENT",Null:"NOT NULL",DataType:"INT(10)"}
	name := &migrator.Column{Name: "name", DataType: "varchar(255)"}
	address := &migrator.Column{Name: "address", DataType: "text"}
	contact := &migrator.Column{Name: "contact_no", DataType: "tinytext"}

	unique := &migrator.Unique{Definition: "name_index"}
	unique.AddColumns(name)

	m.AddColumns(id, name, address, contact).AddPrimary(id).AddUnique(unique)

}

// Reverse the migrations
func (m *Stores_20170704_165247) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.ModifyType = "delete"
	sql := m.GetSQL()
	beego.BeeLogger.Info(fmt.Sprint(sql))
	m.SQL(sql)

}
