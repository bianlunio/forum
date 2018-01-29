package routers

import (
	"github.com/gin-gonic/gin"
)

type MiddleWares []gin.HandlerFunc

type route struct {
	Method      string
	Path        string
	MiddleWares MiddleWares
	Handler     gin.HandlerFunc
}

type routes []route

func RouterMapper(g *gin.RouterGroup, rs routes) {
	for _, route := range rs {
		handlers := append(route.MiddleWares, route.Handler)
		g.Handle(route.Method, route.Path, handlers...)
	}
}

func bindRouter(g *gin.RouterGroup, relativePath string, rs routes) {
	RouterMapper(g.Group(relativePath), rs)
}

func SetRouter() *gin.Engine {
	r := gin.Default()

	gin.Recovery()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	apiRouter := r.Group("/api")

	bindRouter(apiRouter, "/topics", TopicRoutes)

	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//    "foo":  "bar",
	//    "manu": "123",
	//}))

	//authorized.POST("admin", func(c *gin.Context) {
	//    user := c.MustGet(gin.AuthUserKey).(string)
	//
	//    // Parse JSON
	//    var json struct {
	//        Value string `json:"value" binding:"required"`
	//    }
	//
	//    if c.Bind(&json) == nil {
	//        DB[user] = json.Value
	//        c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//    }
	//})

	return r
}
