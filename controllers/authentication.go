// Package controllers implement the controller functionality of the system
package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"github.com/gnanakeethan/posbin/src/auth"
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
// @Success 200 {object} responses.Authentication
// @Failure 403 Forbidden Request
// @router / [post]
func (c *AuthenticationController) Post() {
	var v requests.AuthenticationRequest
	var response responses.Authentication
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Validate() {
			auth.Authenticate(v, &response)
			c.Data["json"] = response
		} else {

			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
