package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getId(c *gin.Context) (id int, err error) {
	rawId := c.Params.ByName("id")
	if rawId == "" {
		return 0, errors.New("parameter error")
	}
	return strconv.Atoi(rawId)
}

func getPagination(c *gin.Context) (page int, limit int, err error) {
	pageRaw := c.DefaultQuery("page", "1")
	limitRaw := c.DefaultQuery("limit", "20")
	page, err = strconv.Atoi(pageRaw)
	limit, err = strconv.Atoi(limitRaw)
	if err != nil {
		return 0, 0, err
	} else if page < 1 || limit < 1 {
		return 0, 0, errors.New("pagination parameter error")
	}
	return page, limit, nil
}

const DBNotFound = "pg: no rows in result set"

func handleError(c *gin.Context, msg string, err error) {
	if err.Error() == DBNotFound {
		c.Status(http.StatusNotFound)
		return
	}
	res := gin.H{"msg": msg}
	if gin.Mode() != gin.ReleaseMode {
		res["error"] = err
	}
	c.JSON(http.StatusBadRequest, res)
}

func handleParameterError(c *gin.Context, err error) {
	handleError(c, "parameter error", err)
}
