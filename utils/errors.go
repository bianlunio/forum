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
