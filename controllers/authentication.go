// Package controllers implement the controller functionality of the system
package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/gnanakeethan/posbin/authentication"
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
// @Description create Bills
// @Param	body		body 	authentication.AuthenticationRequest	true		"Authentication Request"
// @Success 201 Successfully Logged In
// @Failure 401 Unauthorized Request
// @Failure 403 Forbidden Request
// @router / [post]
func (c *AuthenticationController) Post() {
	var v authentication.AuthenticationRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Validate() {
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
