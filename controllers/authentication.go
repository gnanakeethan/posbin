// Package controllers implement the controller functionality of the system
package controllers

import (
	"github.com/astaxie/beego"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

// URLMapping is used to map default routes of the controller
func (c *AuthenticationController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Post", c.Post)
	// c.Mapping("AllBlock", c.AllBlock)
}

// Index method is used to catch the root URL request
//
// @router / [get]
func (c *AuthenticationController) Index() {
	c.Data["json"] = "Success"
	c.ServeJSON()
}

// Post method is used to send login requests
//
// @router / [post]
func (c *AuthenticationController) Post() {

	// email := c.GetString("email")
	// password := c.GetString("password")
	c.Data["json"] = "Success"
	c.ServeJSON()
}
