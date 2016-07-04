package models

type PermissionRole struct {
	PermissionId *Permissions `orm:"column(permission_id);rel(fk)"`
	RoleId       *Roles       `orm:"column(role_id);rel(fk)"`
}
