package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController operations for Error
type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["json"] = map[string]interface{}{"errorCode": 404, "errorStr": "NOT FOUND"}
	this.ServeJSON()
}
func (this *ErrorController) Error401() {
	this.Data["json"] = map[string]interface{}{"errorCode": 401, "errorStr": "Unauthorized Request"}
	this.ServeJSON()
}
