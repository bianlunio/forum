package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func IDValidator(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.Next()
	} else if !bson.IsObjectIdHex(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "invalid id"})
	} else {
		c.Set("_id", bson.ObjectIdHex(id))
		c.Next()
	}
}
