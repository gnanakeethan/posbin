package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id            int        `orm:"column(id);auto"`
	Username      string     `orm:"column(username);size(255)"`
	Email         string     `orm:"column(email);size(255)"`
	Password      string     `orm:"column(password);size(255)" json:"-"`
	RememberToken string     `orm:"column(remember_token);size(100);null" json:"-"`
	CreatedAt     time.Time  `orm:"column(created_at);type(timestamp);null" json:"-"`
	UpdatedAt     time.Time  `orm:"column(updated_at);type(timestamp);null" json:"-"`
	TerminalId    *Terminals `orm:"null;reverse(one)"`
	Roles         []*Roles   `orm:"rel(m2m);rel_through(github.com/gnanakeethan/posbin/models.RoleUser)"`
	Stores        []*Stores  `orm:"reverse(many);rel_table(store_user)"`
}
type UsersPassword struct {
	Id       int    `orm:"column(id);auto"`
	Password string `orm:"column(password);size(255)"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "Roles")
		for _, el := range v.Roles {
			o.LoadRelated(el, "Permissions")
			logs.Info(el)
		}
		o.LoadRelated(v, "Stores")
		logs.Info(v);
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
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

	var l []Users
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}

	m2m := o.QueryM2M(&v, "Roles")
	m2m.Clear()
	for _, p := range m.Roles {
		logs.Info(p)
		if p != nil {
			m2m.Add(p)
		}
	}
	m2m = o.QueryM2M(&v, "Stores")
	m2m.Clear()
	for _, q := range m.Stores {
		logs.Info(q)
		if q != nil {
			m2m.Add(q)
		}
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (user *Users) HasRole(role string) bool {
	o := orm.NewOrm()
	var roles []*Roles
	o.Raw("select r.* from users u "+
		" inner join role_user ru on ru.user_id=u.id"+
		" inner join roles r on r.id=ru.role_id "+
		" where u.id=? and r.name=?"+
		"", user.Id, role).QueryRows(&roles)

	if len(roles) > 0 {
		return true
	}
	return false
}
