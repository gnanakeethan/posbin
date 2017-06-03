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

type Inventories struct {
	Id             int               `orm:"column(id);auto"`
	BatchNo        string            `orm:"column(batch_no);size(255)"`
	Expiry         time.Time         `orm:"column(expiry);type(date);null"`
	ProductId      *Products         `orm:"column(product_id);rel(fk)"`
	InventoryScale []*InventoryScale `orm:"reverse(many)"`
	Purchases      []*Purchases      `orm:"reverse(many)"`
	CreatedAt      time.Time         `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt      time.Time         `orm:"column(updated_at);type(timestamp);null"`
	Service        int8              `orm:"column(service);null"`
}

func (t *Inventories) TableName() string {
	return "inventories"
}

func init() {
	orm.RegisterModel(new(Inventories))
}

// AddInventories insert a new Inventories into database and returns
// last inserted Id on success.
func AddInventories(m *Inventories) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInventoriesById retrieves Inventories by Id. Returns error if
// Id doesn't exist
func GetInventoriesById(id int) (v *Inventories, err error) {
	o := orm.NewOrm()
	v = &Inventories{Id: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "ProductId")
		_, err = o.LoadRelated(v, "Purchases")
		return v, nil
	}
	return nil, err
}

// GetAllInventories retrieves all Inventories matches certain condition. Returns empty list if
// no records exist
func GetAllInventories(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Inventories))
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

	var l []Inventories
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		newl := []Inventories{}
		for _, el := range l {
			o.LoadRelated(&el, "InventoryScale")
			logs.Info(el)
			newl = append(newl, el)
		}
		l = newl
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

// UpdateInventories updates Inventories by Id and returns error if
// the record to be updated doesn't exist
func UpdateInventoriesById(m *Inventories) (err error) {
	o := orm.NewOrm()
	v := Inventories{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInventories deletes Inventories by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInventories(id int) (err error) {
	o := orm.NewOrm()
	v := Inventories{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Inventories{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
