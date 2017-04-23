package controllers

import (
	"github.com/astaxie/beego"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

func (c *AuthenticationController) URLMapping() {
	c.Mapping("Index", c.Index)
	//c.Mapping("AllBlock", c.AllBlock)
}

// @router / [get]
func (c *AuthenticationController) Index() {
	c.Data["json"] = "Success"
	c.ServeJSON()
}
