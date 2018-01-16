package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func GetDB(c *gin.Context) *mgo.Database {
	return c.MustGet("db").(*mgo.Database)
}
