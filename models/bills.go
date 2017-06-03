package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Bills struct {
	Id         int        `orm:"column(id);auto"`
	CustomerId *Customers `orm:"column(customer_id);rel(fk)"`
	Cost       float64    `orm:"column(cost);null"`
	GrossTotal float64    `orm:"column(gross_total);null"`
	Discount   float64    `orm:"column(discount);null"`
	NetTotal   float64    `orm:"column(net_total);null"`
	Balance    float64    `orm:"column(balance);null"`
	CardPaid   float64    `orm:"column(card_paid);null"`
	CashPaid   float64    `orm:"column(cash_paid);null"`
	UserId     *Users     `orm:"column(user_id);rel(fk)" required:"true"`
	TerminalId *Terminals `orm:"column(terminal_id);rel(fk)" required:"true"`
	DeletedAt  time.Time  `orm:"column(deleted_at);type(timestamp);null"`
	CreatedAt  time.Time  `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt  time.Time  `orm:"column(updated_at);type(timestamp);null"`
}

func (t *Bills) TableName() string {
	return "bills"
}

func init() {
	orm.RegisterModel(new(Bills))
}

// AddBills insert a new Bills into database and returns
// last inserted Id on success.
func AddBills(m *Bills) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBillsById retrieves Bills by Id. Returns error if
// Id doesn't exist
func GetBillsById(id int) (v *Bills, err error) {
	o := orm.NewOrm()
	v = &Bills{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetPayableBills() (bills []Bills, err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("SELECT T0.* FROM `bills` T0 WHERE T0.`balance` >= 0 AND T0.`net_total` >= T0.`cash_paid`+T0.`card_paid`")
	update := fmt.Sprintf("UPDATE " +
		"bills b," +
		"(select bill_id, SUM(total) AS total,SUM(discount) AS discount,SUM(cost) as cost FROM sales GROUP BY bill_id) AS bill_total " +
		"SET " +
		"b.cost = bill_total.cost," +
		"b.gross_total = bill_total.total," +
		"b.net_total = bill_total.total-bill_total.discount, " +
		"b.discount=bill_total.discount," +
		"b.balance=(bill_total.total-bill_total.discount)-(b.card_paid+b.cash_paid) " +
		"WHERE " +
		"b.id=bill_total.bill_id")
	_, err = o.Raw(update).Exec()
	_, err = o.Raw(sql).QueryRows(&bills)
	if err == nil {
		return bills, err
	}
	return nil, err
}

// GetAllBills retrieves all Bills matches certain condition. Returns empty list if
// no records exist
func GetAllBills(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Bills))
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

	var l []Bills
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

// UpdateBills updates Bills by Id and returns error if
// the record to be updated doesn't exist
func UpdateBillsById(m *Bills) (err error) {
	o := orm.NewOrm()
	v := Bills{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBills deletes Bills by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBills(id int) (err error) {
	o := orm.NewOrm()
	v := Bills{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Bills{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
