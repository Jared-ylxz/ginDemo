package router

import (
	"github.com/gin-gonic/gin"
)

var (
	anonymousRouterRole = make([]func(*gin.RouterGroup), 0)
)

func Init(r *gin.Engine) {
	registerRoute(r)
}

func registerRoute(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	for _, f := range anonymousRouterRole {
		f(v1)
	}
}