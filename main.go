package main

import (
	_ "github.com/gnanakeethan/pospo5/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/posres")

}

func main() {
	orm.Debug = true

	beego.BConfig.EnableGzip = true
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.ListenTCP4 = true
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSAddr = ""
	beego.BConfig.Listen.HTTPSPort = 8443
	beego.BConfig.Listen.HTTPSCertFile = "conf/cert.pem"
	beego.BConfig.Listen.HTTPSKeyFile = "conf/key.pem"
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.WebConfig.StaticDir["/app"] = "public"
	beego.InsertFilter("/v1/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
