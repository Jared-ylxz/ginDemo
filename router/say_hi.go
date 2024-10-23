package router

import (
	"github.com/gin-gonic/gin"
	"ginDemo/service"
)

func init() {
	anonymousRouterRole = append(anonymousRouterRole, registerAgentRoute)
}

func registerAgentRoute(rg *gin.RouterGroup) {
	rg.GET("/sayhi", service.Sayhi)
}