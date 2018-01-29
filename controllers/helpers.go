package controllers

import (
	"net/http"
	"net/url"

	"forum/utils"

	"github.com/gin-gonic/gin"
)

func getPagination(c *gin.Context) url.Values {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	if utils.String2Int(page) < 1 || utils.String2Int(limit) < 1 {
		panic("Pagination Error")
	}
	values := url.Values{}
	values.Set("page", page)
	values.Set("limit", limit)
	return values
}

func handleParameterError(c *gin.Context, err error) {
	msg := "parameter error"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
}
