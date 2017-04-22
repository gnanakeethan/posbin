package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Sales struct {
	Id          int          `orm:"column(id);auto"`
	BillId      *Bills       `orm:"column(bill_id);rel(fk)"`
	InventoryId *Inventories `orm:"column(inventory_id);rel(fk)"`
	Units       float64      `orm:"column(units)"`
	Cost        float64      `orm:"column(cost);null"`
	Hide        int8         `orm:"column(hide);null"`
	UnitPrice   float64      `orm:"column(unit_price)"`
	Total       float64      `orm:"column(total)"`
	Discount    float64      `orm:"column(discount)"`
	Amount      float64      `orm:"column(amount)"`
	DeletedAt   time.Time    `orm:"column(deleted_at);type(timestamp);null"`
	CreatedAt   time.Time    `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time    `orm:"column(updated_at);type(timestamp);null"`
}

func (t *Sales) TableName() string {
	return "sales"
}

func init() {
	orm.RegisterModel(new(Sales))
}

// AddSales insert a new Sales into database and returns
// last inserted Id on success.
func AddSales(m *Sales) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSalesById retrieves Sales by Id. Returns error if
// Id doesn't exist
func GetSalesById(id int) (v *Sales, err error) {
	o := orm.NewOrm()
	v = &Sales{Id: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "InventoryId")
		return v, nil
	}
	return nil, err
}

// GetAllSales retrieves all Sales matches certain condition. Returns empty list if
// no records exist
func GetAllSales(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sales))
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

	var l []Sales
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

// UpdateSales updates Sales by Id and returns error if
// the record to be updated doesn't exist
func UpdateSalesById(m *Sales) (err error) {
	o := orm.NewOrm()
	v := Sales{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSales deletes Sales by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSales(id int) (err error) {
	o := orm.NewOrm()
	v := Sales{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sales{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
