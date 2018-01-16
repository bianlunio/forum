package utils

import (
	"errors"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func GetId(c *gin.Context) bson.ObjectId {
	id := c.Params.ByName("id")
	return bson.ObjectIdHex(id)
}

func Pagination(c *gin.Context, query *mgo.Query) (func() *mgo.Query, func(interface{}) gin.H, error) {
	rawPage := c.DefaultQuery("page", "1")
	rawSize := c.DefaultQuery("size", "20")
	page := String2Int(rawPage)
	size := String2Int(rawSize)
	if page < 1 || size < 1 {
		return nil, nil, errors.New("pagination error")
	} else {
		// get a copy, so the original query is not tampered with
		q := *query
		skip := (page - 1) * size
		queryBuilder := func() *mgo.Query {
			return q.Skip(skip).Limit(size)
		}
		responseBuilder := func(data interface{}) gin.H {
			total, _ := query.Count()
			return gin.H{
				"page":      page,
				"size":      size,
				"total":     total,
				"totalPage": math.Ceil(float64(total) / float64(size)),
				"list":      data,
			}
		}
		return queryBuilder, responseBuilder, nil
	}
}
