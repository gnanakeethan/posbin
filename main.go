package main

import (
	_ "github.com/gnanakeethan/posbin/routers"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gnanakeethan/posbin/database"
)

func init() {
	dbdriver := gC("dbdriver")
	conn := gC(dbdriver+"user") + ":" + gC(dbdriver+"pass") + "@tcp(" + gC(dbdriver+"urls") + ")/" + gC(dbdriver+"db") + "?charset=utf8&loc=Asia%2FColombo"
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
	database.Migrate()

	//Enables GZip functionality
	beego.BConfig.EnableGzip = true
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js", ".html"}

	//Listens on 0.0.0.0 instead of ::0
	beego.BConfig.Listen.ListenTCP4 = true

	//Configuring HTTPs Mode
	beego.BConfig.Listen.HTTPSCertFile = gC("certfile")
	beego.BConfig.Listen.HTTPSKeyFile = gC("certkey")

	//Expose Swagger Directory
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.Run()
}
