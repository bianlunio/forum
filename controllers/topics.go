package controllers

import (
	"net/http"
	"time"

	. "forum/models"
	"forum/utils"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func TopicList(c *gin.Context) {
	db := utils.GetDB(c)
	rawQuery := db.C("topics").Find(bson.M{"deleted": false}).Sort("-createAt")
	if queryBuilder, responseBuilder, err := utils.Pagination(c, rawQuery); err == nil {
		topics := Topics{}
		query := queryBuilder()
		if err := query.All(&topics); err == nil {
			data := responseBuilder(topics)
			c.JSON(http.StatusOK, data)
		} else {
			utils.HandleDBError(c, err)
		}
	} else {
		utils.HandleParameterError(c, err)
	}
}

func TopicDetail(c *gin.Context) {
	id := utils.GetId(c)
	db := utils.GetDB(c)
	topic := Topic{}
	err := db.C("topics").Find(bson.M{"_id": id, "deleted": false}).One(&topic)
	if err != nil {
		utils.HandleDBError(c, err)
	} else {
		c.JSON(http.StatusOK, topic)
	}
}

type topicData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateTopic(c *gin.Context) {
	var data topicData
	if err := c.Bind(&data); err == nil {
		db := utils.GetDB(c)

		t := Topic{}
		topic := t.Create(data.Title, data.Content, "")
		err := db.C("topics").Insert(&topic)
		if err != nil {
			utils.HandleDBError(c, err)
		} else {
			c.Status(http.StatusCreated)
		}
	} else {
		utils.HandleParameterError(c, nil)
	}
}

func EditTopic(c *gin.Context) {
	id := utils.GetId(c)
	var data topicData
	if err := c.Bind(&data); err == nil {
		db := utils.GetDB(c)

		err := db.C("topics").UpdateId(id, bson.M{
			"$set": bson.M{
				"title":    data.Title,
				"content":  data.Content,
				"updateAt": time.Now(),
			},
		})
		if err != nil {
			utils.HandleDBError(c, err)
		} else {
			c.Status(http.StatusOK)
		}
	} else {
		utils.HandleParameterError(c, nil)
	}
}

func DeleteTopic(c *gin.Context) {
	id := utils.GetId(c)
	db := utils.GetDB(c)
	//err := db.C("topics").RemoveId(id)
	err := db.C("topics").UpdateId(id, bson.M{
		"$set": bson.M{
			"deleted":  true,
			"deleteAt": time.Now(),
		},
	})
	if err != nil {
		utils.HandleDBError(c, err)
	} else {
		c.Status(http.StatusNoContent)
	}
}
