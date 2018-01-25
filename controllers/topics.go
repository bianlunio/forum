package controllers

import (
	"net/http"

	"forum/models"
	"forum/utils"

	"github.com/gin-gonic/gin"
)

func TopicList(c *gin.Context) {
	data := models.Topic{}.List(1, 20)
	c.JSON(http.StatusOK, data)
	//db := utils.GetDB(c)
	//rawQuery := db.C("topics").Find(bson.M{"deleted": false}).Sort("-createdAt")
	//if queryBuilder, responseBuilder, err := utils.Pagination(c, rawQuery); err == nil {
	//	topics := Topics{}
	//	query := queryBuilder()
	//	if err := query.All(&topics); err == nil {
	//		data := responseBuilder(topics)
	//	} else {
	//		utils.HandleDBError(c, err)
	//	}
	//} else {
	//	utils.HandleParameterError(c, err)
	//}
}

//func TopicDetail(c *gin.Context) {
//	id := utils.GetId(c)
//	db := utils.GetDB(c)
//	topic := Topic{}
//	err := db.C("topics").Find(bson.M{"_id": id, "deleted": false}).One(&topic)
//	if err != nil {
//		utils.HandleDBError(c, err)
//	} else {
//		c.JSON(http.StatusOK, topic)
//	}
//}
//
//type topicData struct {
//	Title   string `json:"title" binding:"required"`
//	Content string `json:"content" binding:"required"`
//}

func CreateTopic(c *gin.Context) {
	var data models.Topic
	if err := c.Bind(&data); err == nil {
		topic := data.Create()
		c.JSON(http.StatusCreated, topic)

		//t := Topic{}
		//topic := t.Create(data.Title, data.Content, "")
		//err := db.C("topics").Insert(&topic)
		//if err != nil {
		//	utils.HandleDBError(c, err)
		//} else {
		//	c.Status(http.StatusCreated)
		//}
	} else {
		utils.HandleParameterError(c, nil)
	}
}

//func EditTopic(c *gin.Context) {
//	id := utils.GetId(c)
//	var data topicData
//	if err := c.Bind(&data); err == nil {
//		db := utils.GetDB(c)
//
//		err := db.C("topics").UpdateId(id, bson.M{
//			"$set": bson.M{
//				"title":    data.Title,
//				"content":  data.Content,
//				"updatedAt": time.Now(),
//			},
//		})
//		if err != nil {
//			utils.HandleDBError(c, err)
//		} else {
//			c.Status(http.StatusOK)
//		}
//	} else {
//		utils.HandleParameterError(c, nil)
//	}
//}
//
//func DeleteTopic(c *gin.Context) {
//	id := utils.GetId(c)
//	db := utils.GetDB(c)
//	//err := db.C("topics").RemoveId(id)
//	err := db.C("topics").UpdateId(id, bson.M{
//		"$set": bson.M{
//			"deleted":  true,
//			"deletedAt": time.Now(),
//		},
//	})
//	if err != nil {
//		utils.HandleDBError(c, err)
//	} else {
//		c.Status(http.StatusNoContent)
//	}
//}
