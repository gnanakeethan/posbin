package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func gC(conf string) string {
	return beego.AppConfig.String(conf)
}

// HomeController operations for Authentication
type HomeController struct {
	beego.Controller
}

func (c *HomeController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("AllBlock", c.Index)

}

// @router / [any]
func (c *HomeController) Index() {

	logs.Info(c.Ctx.Input.IsSecure())
	if !c.Ctx.Input.IsSecure() {
		url := "https://" + gC("HTTPSAddr") + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/" + "index.html"
		logs.Info(url)
		c.Ctx.Redirect(302, url)
	}

}
