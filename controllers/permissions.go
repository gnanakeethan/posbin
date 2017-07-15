package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/gnanakeethan/posbin/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// oprations for Permissions
type PermissionsController struct {
	beego.Controller
}

func (c *PermissionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Permissions
// @Param	body		body 	models.Permissions	true		"body for Permissions content"
// @Success 201 {int} models.Permissions
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionsController) Post() {
	var v models.Permissions
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPermissions(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Permissions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Permissions
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PermissionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPermissionsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Permissions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Permissions
// @Failure 403
// @router / [get]
func (c *PermissionsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllPermissions(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Permissions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Permissions	true		"body for Permissions content"
// @Success 200 {object} models.Permissions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PermissionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Permissions{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdatePermissionsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Permissions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PermissionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePermissions(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Permissions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /update_permissions/ [get]
func (c *PermissionsController) Routes() {
	// Encoding the map
	content := beego.PrintTree()
	routelist := make(map[string]map[string]string)
	routelist["get"] = make(map[string]string)
	routelist["put"] = make(map[string]string)
	routelist["post"] = make(map[string]string)
	routelist["delete"] = make(map[string]string)
	for _, el := range content["Data"].(map[string]interface{}) {
		elp := el.(*[][]string)
		for _, p := range *elp {
			p[0] = strings.Replace(p[0], ":id/", "id/", -1)
			p[0] = strings.Replace(p[0], ":id", "id/", -1)
			p[0] = strings.TrimPrefix(p[0], "/v1/")
			p[1] = strings.Replace(p[1], "map[", "", -1)
			p[1] = strings.Replace(p[1], "]", "", -1)
			method := strings.ToLower(strings.Split(p[1], ":")[0])
			routelist[method][p[0]] = method
		}
	}
	routemap := make(map[string]string)
	for _, method := range routelist {
		for key, lap := range method {
			routemap[strings.Replace(key, "/", "_", -1)+lap] = lap
		}
	}
	o := orm.NewOrm()
	for route, method := range routemap {

		permission := &models.Permissions{Name: route}
		permission.DisplayName = strings.Replace(route, "_", " ", -1)
		permission.DisplayName = strings.Title(permission.DisplayName)
		switch method {
		case "get":
			permission.Description = "CAN RETRIEVE " + strings.Replace(strings.TrimSuffix(route, "get"), "_", "/", -1)
		case "post":
			permission.Description = "CAN CREATE " + strings.Replace(strings.TrimSuffix(route, "post"), "_", "/", -1)
		case "put":
			permission.Description = "CAN UPDATE " + strings.Replace(strings.TrimSuffix(route, "put"), "_", "/", -1)
		case "delete":
			permission.Description = "CAN DELETE " + strings.Replace(strings.TrimSuffix(route, "delete"), "_", "/", -1)
		}
		o.InsertOrUpdate(permission, "name")
		logs.Info(permission)
	}
	c.Data["json"] = routemap
	c.ServeJSON()
}
