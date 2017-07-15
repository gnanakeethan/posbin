package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Permissions struct {
	Id          int       `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(255)"`
	DisplayName string    `orm:"column(display_name);size(255);null"`
	Description string    `orm:"column(description);size(255);null"`
	CreatedAt   time.Time `orm:"auto_now_add;column(created_at);type(datetime);null"`
	UpdatedAt   time.Time `orm:"auto_now;column(updated_at);type(datetime);null"`
	Roles       []*Roles  `orm:"reverse(many)"`
}

func (t *Permissions) TableName() string {
	return "permissions"
}

func init() {
	orm.RegisterModel(new(Permissions))
}

// AddPermissions insert a new Permissions into database and returns
// last inserted Id on success.
func AddPermissions(m *Permissions) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPermissionsForUser retrieves all permissions of the user
func GetPermissionsForUser(id int) (v []*Permissions, err error) {
	o := orm.NewOrm()
	query := "select p.* from users u inner join role_user ru on ru.user_id=u.id " +
		"inner join roles r on ru.role_id=r.id " +
		"inner join permission_role rp on rp.role_id=r.id " +
		"inner join permissions p on p.id=rp.permission_id " +
		"where u.id = ?"
	if _, err := o.Raw(query, id).QueryRows(&v); err == nil {
		return v, nil
	}
	return nil, err

}

// GetPermissionsById retrieves Permissions by Id. Returns error if
// Id doesn't exist
func GetPermissionsById(id int) (v *Permissions, err error) {
	o := orm.NewOrm()
	v = &Permissions{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPermissions retrieves all Permissions matches certain condition. Returns empty list if
// no records exist
func GetAllPermissions(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permissions))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Permissions
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePermissions updates Permissions by Id and returns error if
// the record to be updated doesn't exist
func UpdatePermissionsById(m *Permissions) (err error) {
	o := orm.NewOrm()
	v := Permissions{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePermissions deletes Permissions by Id and returns error if
// the record to be deleted doesn't exist
func DeletePermissions(id int) (err error) {
	o := orm.NewOrm()
	v := Permissions{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Permissions{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
