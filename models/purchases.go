package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Purchases struct {
	Id          int          `orm:"column(id);auto"`
	InventoryId *Inventories `orm:"column(inventory_id);rel(fk)"`
	AverageCost float64      `orm:"column(average_cost)"`
	Units       float64      `orm:"column(units)"`
	CreatedAt   time.Time    `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time    `orm:"column(updated_at);type(timestamp);null"`
}

func (t *Purchases) TableName() string {
	return "purchases"
}

func init() {
	orm.RegisterModel(new(Purchases))
}

// AddPurchases insert a new Purchases into database and returns
// last inserted Id on success.
func AddPurchases(m *Purchases) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPurchasesById retrieves Purchases by Id. Returns error if
// Id doesn't exist
func GetPurchasesById(id int) (v *Purchases, err error) {
	o := orm.NewOrm()

	v = &Purchases{Id: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "InventoryId")
		return v, nil
	}
	return nil, err
}

// GetAllPurchases retrieves all Purchases matches certain condition. Returns empty list if
// no records exist
func GetAllPurchases(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Purchases))
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

	var l []Purchases
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

// UpdatePurchases updates Purchases by Id and returns error if
// the record to be updated doesn't exist
func UpdatePurchasesById(m *Purchases) (err error) {
	o := orm.NewOrm()
	v := Purchases{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePurchases deletes Purchases by Id and returns error if
// the record to be deleted doesn't exist
func DeletePurchases(id int) (err error) {
	o := orm.NewOrm()
	v := Purchases{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Purchases{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
