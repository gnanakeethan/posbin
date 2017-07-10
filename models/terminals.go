package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/imdario/mergo"
)

type Terminals struct {
	Id        int       `orm:"column(id);auto"`
	UserId    *Users    `orm:"column(user_id);rel(one);"`
	Machine   string    `orm:"column(machine);size(255)"`
	Printer   string    `orm:"column(printer);size(255)"`
	Balance   float64   `orm:"column(balance)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	StoreId   *Stores   `orm:"column(store_id);rel(fk)"`
}

func (t *Terminals) TableName() string {
	return "terminals"
}

func init() {
	orm.RegisterModel(new(Terminals))
}

// AddTerminals insert a new Terminals into database and returns
// last inserted Id on success.
func AddTerminals(m *Terminals) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTerminalsById retrieves Terminals by Id. Returns error if
// Id doesn't exist
func GetTerminalsById(id int) (v *Terminals, err error) {
	o := orm.NewOrm()
	v = &Terminals{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTerminals retrieves all Terminals matches certain condition. Returns empty list if
// no records exist
func GetAllTerminals(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Terminals))
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

	var l []Terminals
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

// UpdateTerminals updates Terminals by Id and returns error if
// the record to be updated doesn't exist
func UpdateTerminalsById(m *Terminals) (err error) {
	o := orm.NewOrm()
	v := Terminals{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		sql := "update terminals set user_id null where user_id=?"
		if v.UserId != nil {
			o.Raw(sql, v.UserId.Id)
		}
		if err = mergo.Merge(m, &v); err == nil {
			var num int64
			if num, err = o.Update(m); err == nil {
				fmt.Println("Number of records updated in database:", num)
			}
		}
	}
	return
}

// DeleteTerminals deletes Terminals by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTerminals(id int) (err error) {
	o := orm.NewOrm()
	v := Terminals{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Terminals{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
