package gen

const TempMain = `
package main

import (
	"fmt"
	"log"
	"TempImportPkg/adapter/driver/handler"
	"TempImportPkg/adapter/driver/router"
	"TempImportPkg/global"
	"TempImportPkg/infra/middleware"

	_ "TempImportPkg/init"

	"github.com/gin-gonic/gin"
)

type Server struct {
	defaultGroup []router.CommonRouter
	// 可以根据需求新建多个组 用于应对不同的路由组前缀
	// defaultGroup1 []router.CommonRouter
	// defaultGroup2 []router.CommonRouter
}

var server = &Server{
	// handler注册到路由组
	defaultGroup: []router.CommonRouter{
		handler.NewHttpUserHandler(),
	},
}

func (s *Server) Start() {
	go func() {
		cfg := global.Config
		engine := gin.New()
		routerV1 := engine.Group("/v1")
		routerV1.Use(middleware.ErrorHandlerMiddleware())
		for _, router := range s.defaultGroup {
			router.InitRouter(routerV1)
		}
		url := fmt.Sprintf("%s:%d", cfg.Project.Host, cfg.Project.Port)
		err := engine.Run(url)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func main() {
	server.Start()
	select {}
}
`
