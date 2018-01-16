package utils

import (
	"github.com/gin-gonic/gin"
)

type MiddleWares []gin.HandlerFunc

type Route struct {
	Method      string
	Path        string
	MiddleWares MiddleWares
	Handler     gin.HandlerFunc
}

type Routes []Route

func RouterMapper(g *gin.RouterGroup, routes Routes) {
	for _, route := range routes {
		handlers := append(route.MiddleWares, route.Handler)
		g.Handle(route.Method, route.Path, handlers...)
	}
}

func BindRouter(g *gin.RouterGroup, relativePath string, routes Routes) {
	RouterMapper(g.Group(relativePath), routes)
}
