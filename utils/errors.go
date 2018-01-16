package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleParameterError(c *gin.Context, err error) {
	msg := "parameter error"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
}

func HandleDBError(c *gin.Context, err error) {
	if err.Error() == "not found" {
		c.Status(http.StatusNotFound)
	} else {
		panic(err.Error())
	}
}
