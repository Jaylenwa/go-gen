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
	TempSvcNameCaseCamelHandler adapterDriver.RouterInterface
}

func (s *Server) Start() {
	go func() {
		engine := gin.New()
		routerGroup := engine.Group("/api/v1")
		routerGroup.Use(middleware.ErrorHandlerMiddleware())
		s.TempSvcNameCaseCamelHandler.RegisterRouterPublic(routerGroup)
		url := fmt.Sprintf("%s:%d", global.GConfig.Project.Host, global.GConfig.Project.Port)
		_ = engine.Run(url)
	}()
}

func main() {

	s := &Server{
		TempSvcNameCaseCamelHandler: adapterDriver.NewTempSvcNameCamelLowerHandler(),
	}
	s.Start()

	select {}
}


`
