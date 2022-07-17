package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"

	"LGIS/api/internal/config"
	"LGIS/api/internal/handler"
	"LGIS/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mgis-api.yaml", "the config file")

//go:embed dist
var content embed.FS

func main() {
	flag.Parse()

	var c config.Config
	// conf.MustLoad(*configFile, &c)
	c.Name = "mgis-api"
	c.Host = "0.0.0.0"
	c.Port = 8888
	c.Db = "xx:xxxx@tcp(xxx:xxx)/xx?charset=utf8mb4&parseTime=True&loc=Local"

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fsys, _ := fs.Sub(content, "dist")
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: http.StripPrefix("/", http.FileServer(http.FS(fsys))).ServeHTTP,
	})

	asys,_ := fs.Sub(content, "dist/assets")
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/assets/:file",
		Handler: http.StripPrefix("/assets/", http.FileServer(http.FS(asys))).ServeHTTP,
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
