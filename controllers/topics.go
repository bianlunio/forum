package controllers

import (
	"net/http"

	"forum/models"

	"github.com/gin-gonic/gin"
)

func TopicList(c *gin.Context) {
	page, limit, err := getPagination(c)
	if err != nil {
		handleParameterError(c, err)
		return
	}
	data, err := models.Topic{}.List(page, limit)
	if err != nil {
		handleError(c, "request fail", err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func TopicDetail(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		handleParameterError(c, err)
		return
	}
	topic, err := models.Topic{}.Detail(id)
	if err != nil {
		handleError(c, "request fail", err)
		return
	}
	c.JSON(http.StatusOK, topic)
}

func CreateTopic(c *gin.Context) {
	data := models.Topic{}
	if err := c.Bind(&data); err != nil {
		handleParameterError(c, err)
		return
	}
	topic, err := models.Topic{}.Create(data)
	if err != nil {
		handleError(c, "request fail", err)
		return
	}
	c.JSON(http.StatusCreated, topic)
}

func EditTopic(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		handleParameterError(c, err)
		return
	}
	data := models.Topic{}
	if err := c.Bind(&data); err != nil {
		handleParameterError(c, err)
		return
	}
	topic, err := models.Topic{}.Update(id, data)
	if err != nil {
		handleError(c, "request fail", err)
		return
	}
	c.JSON(http.StatusOK, topic)
}

func DeleteTopic(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		handleParameterError(c, err)
		return
	}
	//err := db.C("topics").RemoveId(id)
	rowsAffected, err := models.Topic{}.SoftDelete(id)
	if err != nil {
		handleError(c, "request fail", err)
		return
	} else if rowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}
