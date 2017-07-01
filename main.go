package main

import (
	"github.com/gnanakeethan/posbin/controllers"
	_ "github.com/gnanakeethan/posbin/routers"

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

// gC grabs the application configuration using beego method
func gC(conf string) string {
	return beego.AppConfig.String(conf)
}

func main() {

	//Enables GZip functionality
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}

	//Listens on 0.0.0.0 instead of ::0
	beego.BConfig.Listen.ListenTCP4 = true

	//Configuring HTTPs Mode
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSCertFile = gC("certfile")
	beego.BConfig.Listen.HTTPSKeyFile = gC("certkey")

	//Expose Swagger Directory
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.WebConfig.DirectoryIndex = true

	//Static Path Configurations
	beego.SetStaticPath("/service-worker.js", gC("publicdir")+"/service-worker.js")
	beego.SetStaticPath("/manifest.json", "public/manifest.json")
	beego.SetStaticPath("/index.html", gC("publicdir")+"/index.html")
	beego.SetStaticPath("/", gC("publicdir")+"/index.html")
	beego.BConfig.WebConfig.StaticDir["/bower_components"] = gC("publicdir") + "/bower_components"
	beego.BConfig.WebConfig.StaticDir["/src"] = gC("publicdir") + "/src"

	beego.InsertFilter("/v1/*", beego.BeforeRouter, controllers.AuthenticateUser, true)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	beego.Run()
}
