package gen

const TempMain = `
package main

import (
	"fmt"
	adapterDriver "TempImportPkg/adapter/driver"
	_ "TempImportPkg/boot"
	"TempImportPkg/global"
	"TempImportPkg/infra/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpTempSvcNameCaseCamelHandler adapterDriver.HttpRouterInterface
}

func (s *Server) Start() {
	go func() {
		engine := gin.New()
		// 注册路由 - 健康检查
		routerHealth := engine.Group("/api/v1")
		routerHealth.Use(middleware.ErrorHandlerMiddleware())
		s.httpTempSvcNameCaseCamelHandler.RegisterRouterPublic(routerHealth)
		url := fmt.Sprintf("%s:%d", global.GConfig.Project.Host, global.GConfig.Project.Port)
		_ = engine.Run(url)
	}()
}

func main() {

	s := &Server{
		httpTempSvcNameCaseCamelHandler: adapterDriver.NewHttpTempSvcNameCaseCamelHandler(),
	}
	s.Start()

	select {}
}


`
