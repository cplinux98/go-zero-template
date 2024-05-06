package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"net/http"
	"strings"
	"test/internal/config"
	"test/internal/handler"
	"test/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/test.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 关闭定时打印的状态信息
	logx.DisableStat()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		//swagger
		registerHandlers(server, "/static/", "./static/")
		fmt.Printf("%s %s %s\n", strings.Repeat("=", 30), service.DevMode, strings.Repeat("=", 30))
		server.PrintRoutes()
		fmt.Printf("%s %s %s\n", strings.Repeat("=", 30), service.DevMode, strings.Repeat("=", 30))
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func dirHandler(prefix, fileDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		h := http.StripPrefix(prefix, http.FileServer(http.Dir(fileDir)))
		h.ServeHTTP(w, req)
	}
}

func registerHandlers(engine *rest.Server, prefix, dirPath string) {
	// Set up the dir level
	dirLevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	for i := 1; i < len(dirLevel); i++ {
		path := prefix + strings.Join(dirLevel[:i], "/")
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirHandler(prefix, dirPath),
			})
	}
}
