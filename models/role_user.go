package models

type RoleUser struct {
	UserId *Users `orm:"column(user_id);rel(fk)"`
	RoleId *Roles `orm:"column(role_id);rel(fk)"`
}
