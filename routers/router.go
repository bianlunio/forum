package routers

import (
	"github.com/bianlunio/forum/middlewares"
	"github.com/bianlunio/forum/utils"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	apiRouter := r.Group("/api", middlewares.BindDBSession)

	utils.BindRouter(apiRouter, "/topics", TopicRoutes)

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
