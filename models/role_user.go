package models

import "github.com/astaxie/beego/orm"

type RoleUser struct {
	Id     int    `orm:"column(id);auto"`
	UserId *Users `orm:"column(user_id);rel(fk)"`
	RoleId *Roles `orm:"column(role_id);rel(fk)"`
}

func (t *RoleUser) TableName() string {
	return "role_user"
}

func init() {
	orm.RegisterModel(new(RoleUser))
}
