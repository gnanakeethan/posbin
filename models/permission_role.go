package models

import "github.com/astaxie/beego/orm"

type PermissionRole struct {
	Id           int          `orm:"column(id);auto"`
	PermissionId *Permissions `orm:"column(permission_id);rel(fk)"`
	RoleId       *Roles       `orm:"column(role_id);rel(fk)"`
}

func (t *PermissionRole) TableName() string {
	return "permission_role"
}

func init() {
	orm.RegisterModel(new(PermissionRole))
}
