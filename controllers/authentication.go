// Package controllers implement the controller functionality of the system
package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"github.com/gnanakeethan/posbin/src/auth"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

// Post method is used to send login requests
//
// @Description create Bills
// @Param	body		body 	authentication.AuthenticationRefreshRequest	true		"Authentication Request"
// @Success 200 {object} responses.Authentication
// @Failure 403 Forbidden Request
// @router /validate/ [post]
func (c *AuthenticationController) PostValidate() {
	var v requests.AuthenticationRefreshRequest
	var response responses.Authentication

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		logs.Info(v.Token)
		if v.Validate() {
			auth.ValidateToken(v, &response)
		}
		c.Data["json"] = response

	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post method is used to send login requests
//
// @Description create Bills
// @Param	body		body 	authentication.AuthenticationRefreshRequest	true		"Authentication Request"
// @Success 200 {object} responses.Authentication
// @Failure 403 Forbidden Request
// @router /validate/ [post]
func (c *AuthenticationController) Logout() {
	var v requests.AuthenticationRefreshRequest
	var response responses.Authentication

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Validate() {
			auth.InvalidateToken(v)
		}
		c.Data["json"] = response

	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post method is used to send login requests
//
// @Description create Bills
// @Param	body		body 	authentication.AuthenticationRefreshRequest	true		"Authentication Request"
// @Success 200 {object} responses.Authentication
// @Failure 403 Forbidden Request
// @router /refresh/ [post]
func (c *AuthenticationController) PostRefresh() {
	var v requests.AuthenticationRefreshRequest
	var response responses.Authentication
	go func() {
		o := orm.NewOrm()
		sql := fmt.Sprintf("delete from tokens where valid_thru < now()")

		_, err := o.Raw(sql).Exec()
		logs.Error(err)
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Validate() {
			auth.RefreshToken(v, &response)
			c.Data["json"] = response
		} else {

			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post method is used to send login requests
//
// @Description create Bills
// @Param	body		body 	authentication.AuthenticationRefreshRequest	true		"Authentication Request"
// @Success 200 {object} responses.Authentication
// @Failure 403 Forbidden Request
// @router /logout/ [post]
func (c *AuthenticationController) PostLogout() {
	var v requests.AuthenticationRefreshRequest
	var response responses.Authentication
	auth.ClearTokens()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Validate() {
			response.Success = false
			response.Token = ""
			go auth.InvalidateToken(v)
			c.Data["json"] = response
		}
	} else {
		c.Data["json"] = err.Error()
	}
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
		}
		c.Data["json"] = response

	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
