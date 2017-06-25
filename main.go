package main

import (
	_ "github.com/gnanakeethan/pospo5/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbdriver := gC("dbdriver")
	conn := gC(dbdriver+"user") + ":" + gC(dbdriver+"pass") + "@tcp(" + gC(dbdriver+"urls") + ")/" + gC(dbdriver+"db")
	orm.RegisterDataBase("default", gC("dbdriver"), conn)
	if gC("runmode") == "dev" {
		orm.Debug = true
	}

}
func gC(conf string) string {
	return beego.AppConfig.String(conf)
}

func main() {

	beego.BConfig.EnableGzip = true

	beego.BConfig.Listen.ListenTCP4 = true
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true

	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}
	beego.BConfig.Listen.HTTPSCertFile = gC("certfile")
	beego.BConfig.Listen.HTTPSKeyFile = gC("certkey")

	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.SetStaticPath("/service-worker.js", gC("publicdir")+"/service-worker.js")
	beego.SetStaticPath("/manifest.json", "public/manifest.json")
	beego.SetStaticPath("/index.html", gC("publicdir")+"/index.html")
	beego.SetStaticPath("/", gC("publicdir")+"/index.html")

	beego.BConfig.WebConfig.StaticDir["/bower_components"] = gC("publicdir") + "/bower_components"
	beego.BConfig.WebConfig.StaticDir["/src"] = gC("publicdir") + "/src"

	beego.InsertFilter("/v1/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
