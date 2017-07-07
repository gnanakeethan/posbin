package main

import (
	"bufio"
	"log"
	"net/http"
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
func getHref(t html.Token) (bool, string, string) {
	href := ""
	impt := ""
	ok := false
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
		if a.Key == "rel" {
			impt = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return ok, href, impt
}
func recurvPreload(path string, paths map[string]string, level *int) {

	if *level > 10 {
		return
	}
	f, _ := os.Open(path)
	replace, _ := regexp.Compile("([a-z-]+).html$")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

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
			// isImport := t.Attr["rel"] == "import"

			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, urlt, _ := getHref(t)
			if !ok {
				continue
			}

			urld, _ := url.Parse(string(rootpath) + string(urlt))

			pushpath, _ := filepath.Abs(urld.EscapedPath())
			pushpathd := strings.Replace(pushpath, dir+"/"+gC("publicdir"), "", -1)

			paths[pushpathd] = pushpath
			recurvPreload(pushpath, paths, level)
			*level++

		}
	}
}

var preload = func(ctx *context.Context) {
	level := 1
	path := gC("publicdir") + "/index.html"
	paths := make(map[string]string)
	recurvPreload(path, paths, &level)

	w := ctx.Output.Context.ResponseWriter.ResponseWriter
	for key, _ := range paths {

		if pusher, ok := w.(http.Pusher); ok {
			// Push is supported.
			if err := pusher.Push(key, nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
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

	beego.Run()
}
