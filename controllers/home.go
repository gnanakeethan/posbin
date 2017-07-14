package controllers

import (
	"strconv"
	"strings"

	"fmt"
	"sort"

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

// @router / [any]
func (c *HomeController) Index() {

	logs.Info(c.Ctx.Input.IsSecure())
	if !c.Ctx.Input.IsSecure() {
		url := "https://" + gC("HTTPSAddr") + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/" + "index.html"
		logs.Info(url)
		c.Ctx.Redirect(302, url)
	}

}

// @router /routes/ [get]
func (c *HomeController) Routes() {
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
			p[0] = strings.TrimSuffix(p[0], ":id")
			p[0] = strings.TrimPrefix(p[0], "/v1/")
			p[1] = strings.Replace(p[1], "map[", "", -1)
			p[1] = strings.Replace(p[1], "]", "", -1)
			method := strings.ToLower(strings.Split(p[1], ":")[0])
			routelist[method][p[0]] = method
		}
	}
	routemap := []string{}
	for _, method := range routelist {
		for key, lap := range method {
			routemap = append(routemap, strings.Replace(key, "/", "_", -1)+lap)
			//logs.Info(lap)
		}
		//logs.Info(el)
	}

	fmt.Println("UNSORTED")
	sort.Sort(Alphabetic(routemap))
	c.Data["json"] = routemap

	c.ServeJSON()
}

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}
