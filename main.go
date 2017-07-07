package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/gnanakeethan/posbin/routers"
	"golang.org/x/net/html"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
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

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

var preload = func(ctx *context.Context) {
	path := gC("publicdir") + ctx.Request.URL.Path
	f, _ := os.Open(path)
	replace, _ := regexp.Compile("([a-z-]+).html$")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	rootpath := replace.ReplaceAll([]byte(path), []byte(""))
	r4 := bufio.NewReader(f)

	z := html.NewTokenizer(r4)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "link"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, urlt := getHref(t)
			if !ok {
				continue
			}

			urld, _ := url.Parse(string(rootpath) + string(urlt))

			pushpath, _ := filepath.Abs(urld.EscapedPath())
			pushpath = strings.Replace(pushpath, dir+"/"+gC("publicdir"), "", -1)
			// ctx.Output.Context.ResponseWriter.Header().Add("Link", fmt.Sprintf(" <%s>; rel=preload", pushpath))
		}
	}
}

func main() {

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

	//Static Path Configurations
	beego.SetStaticPath("/service-worker.js", gC("publicdir")+"/service-worker.js")
	beego.SetStaticPath("/manifest.json", "public/manifest.json")
	beego.SetStaticPath("/index.html", gC("publicdir")+"/index.html")
	beego.SetStaticPath("/", gC("publicdir")+"/index.html")
	beego.BConfig.WebConfig.StaticDir["/bower_components"] = gC("publicdir") + "/bower_components"
	beego.BConfig.WebConfig.StaticDir["/src"] = gC("publicdir") + "/src"

	beego.InsertFilter("/*", beego.BeforeStatic, preload)

	beego.Run()
}
