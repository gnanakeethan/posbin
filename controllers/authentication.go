// Package controllers implement the controller functionality of the system
package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/gnanakeethan/posbin/requests"
	"github.com/gnanakeethan/posbin/responses"
	"github.com/gnanakeethan/posbin/src/auth"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

var AuthenticateUser = func(ctx *context.Context) {
	versionString := "/v1/"
	request := ctx.Request.RequestURI
	request = strings.Replace(request, versionString, "", 1)
	if !strings.HasPrefix(request, "auth") {
		response := responses.Authentication{}
		token := ctx.Request.Header.Get("Authorization")
		if token != "" {
			auth.ValidateToken(requests.AuthenticationRefreshRequest{Token: token}, &response)
		}
		logs.Info(token)

		if response.Success {

		} else {
			if beego.AppConfig.String("runmode") != "dev" {
				ctx.Abort(401, string("ftesin"))
			}
		}

	}
}

// URLMapping is used to map default routes of the controller
func (c *AuthenticationController) URLMapping() {

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
		if v.Validate() {
			auth.ValidateToken(v, &response)
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
// @router /refresh/ [post]
func (c *AuthenticationController) PostRefresh() {
	var v requests.AuthenticationRefreshRequest
	var response responses.Authentication
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
