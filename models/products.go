package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Products struct {
	Id          int            `orm:"column(id);auto"`
	ProductCode string         `orm:"column(product_code);size(255);null"`
	Barcode     string         `orm:"column(barcode);size(255);null"`
	Name        string         `orm:"column(name);size(255)"`
	Singular    int8           `orm:"column(singular)"`
	Inventory   []*Inventories `orm:"reverse(many)"`
	Priority    int8           `orm:"column(priority);null"`
	Ordering    uint           `orm:"column(ordering);null"`
	Image       string         `orm:"column(image);null"`
	ImageType   string         `orm:"column(image_type);null"`
	ScaleId     *Scales        `orm:"column(scale_id);rel(fk)"`
	CreatedAt   time.Time      `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time      `orm:"column(updated_at);type(timestamp);null"`
	Service     int8           `orm:"column(service);null"`
	StoreId     *Stores        `orm:"column(store_id);rel(fk)"`
}

func (t *Products) TableName() string {
	return "products"
}

func init() {
	orm.RegisterModel(new(Products))
}

// AddProducts insert a new Products into database and returns
// last inserted Id on success.
func AddProducts(m *Products) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductsById retrieves Products by Id. Returns error if
// Id doesn't exist
func GetProductsById(id int) (v *Products, err error) {
	o := orm.NewOrm()
	v = &Products{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "Inventory")
		return v, nil
	}
	return nil, err
}

// GetAllProducts retrieves all Products matches certain condition. Returns empty list if
// no records exist
func GetAllProducts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Products))
	// query k=v
	likeMap := map[string]bool{
		"Name":        true,
		"Barcode":     true,
		"ProductCode": true,
	}

	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if likeMap[k] {
			//__like condition
			qs = qs.Filter(k+"__icontains", v)
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []Products
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

// GetAllProducts retrieves all Products matches certain condition. Returns empty list if
// no records exist
func GetBillProducts(query map[string]string, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var products []Products
	if query["Barcode"] != "" {
		query["Barcode"] = strings.Replace(query["Barcode"], " ", "%", 10)
		query["Barcode"] = "%" + query["Barcode"] + "%"
		_, err = o.Raw("select * from products where barcode like ? limit ?", query["Barcode"], limit).QueryRows(&products)
	} else {
		query["Name"] = strings.Replace(query["Name"], " ", "%", 10)
		query["Name"] = "%" + query["Name"] + "%"
		_, err = o.Raw("select * from products where name like ? or product_code like ?  limit ?", query["Name"], query["Name"], limit).QueryRows(&products)
	}
	for _, v := range products {
		o.LoadRelated(&v, "Inventory")
		ml = append(ml, v)
	}
	return ml, nil
}

// UpdateProducts updates Products by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductsById(m *Products) (err error) {
	o := orm.NewOrm()
	v := Products{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProducts deletes Products by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProducts(id int) (err error) {
	o := orm.NewOrm()
	v := Products{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Products{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
