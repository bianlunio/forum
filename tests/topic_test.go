package tests

import (
	"testing"

	"github.com/bianlunio/forum/models"
	"github.com/bianlunio/forum/routers"
	"github.com/bianlunio/forum/utils"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTopicRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routers.SetRouter()

	var db *mgo.Database

	Convey("TopicRouter", t, func() {
		baseUrl := "/api/topics"

		session := models.Session.Clone()
		db = utils.GetTestDB(session)
		var t1 models.Topic
		topic1 := t1.Create("t1", "t1t1t1", "")
		var t2 models.Topic
		topic2 := t2.Create("t2", "t2t2t2", "")
		if err := db.C("topics").Insert(&topic1, &topic2); err != nil {
			panic(err)
		}

		Convey("list topics", func() {
			Convey("should success", func() {
				w := utils.TestRequest(router, "GET", baseUrl, "")
				So(w.Code, ShouldEqual, 200)
				data := utils.ListResponse2Dict(w.Body.Bytes())

				So(data.Page, ShouldEqual, 1)
				So(data.Total, ShouldEqual, 2)
				So(data.TotalPage, ShouldEqual, 1)
				So(data.Size, ShouldEqual, 20)

				list := data.List
				So(list[0]["title"], ShouldEqual, "t1")
				So(list[0]["content"], ShouldEqual, "t1t1t1")
				So(list[1]["title"], ShouldEqual, "t2")
				So(list[1]["content"], ShouldEqual, "t2t2t2")

				Convey("when size = 1", func() {
					query := utils.Query{"size": "1"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := utils.TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 200)
					data := utils.ListResponse2Dict(w.Body.Bytes())
					So(data.Page, ShouldEqual, 1)
					So(data.Total, ShouldEqual, 2)
					So(data.TotalPage, ShouldEqual, 2)
					So(data.Size, ShouldEqual, 1)

					list := data.List
					So(list[0]["title"], ShouldEqual, "t1")
					So(list[0]["content"], ShouldEqual, "t1t1t1")
				})
			})

			Convey("should fail", func() {
				Convey("with invalid page number", func() {
					query := utils.Query{"page": "0"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := utils.TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 400)
					data := utils.Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "pagination error")
				})

				Convey("with invalid size number", func() {
					query := utils.Query{"size": "0"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := utils.TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 400)
					data := utils.Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "pagination error")
				})
			})
		})

		Convey("get topic detail", func() {
			Convey("should success", func() {
				url := utils.JoinUrl(baseUrl, t1.Id.Hex())
				w := utils.TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 200)

				data := utils.Response2Dict(w.Body.Bytes())
				So(data["id"], ShouldEqual, t1.Id.Hex())
				So(data["title"], ShouldEqual, "t1")
				So(data["content"], ShouldEqual, "t1t1t1")
			})

			Convey("should fail with wrong id", func() {
				dummyId := utils.JoinStrings("123", t1.Id.Hex()[3:])
				url := utils.JoinUrl(baseUrl, dummyId)
				w := utils.TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 404)
			})

			Convey("should fail with invalid id", func() {
				url := utils.JoinUrl(baseUrl, "someErrorId")
				w := utils.TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 400)
				data := utils.Response2Dict(w.Body.Bytes())
				So(data["msg"], ShouldEqual, "invalid id")
			})
		})

		Convey("create topic", func() {
			Convey("should success", func() {
				body := `{"title": "t3", "content": "t3t3t3"}`
				w := utils.TestRequest(router, "POST", baseUrl, body)
				So(w.Code, ShouldEqual, 201)

				var topic models.Topic
				err := db.C("topics").Find(bson.M{"title": "t3"}).One(&topic)
				So(err, ShouldBeNil)
				So(topic.Id.Hex(), ShouldNotBeEmpty)
				So(topic.Title, ShouldEqual, "t3")
				So(topic.Content, ShouldEqual, "t3t3t3")
				So(topic.Deleted, ShouldBeFalse)
				So(topic.DeleteAt, ShouldBeNil)
			})

			Convey("should fail", func() {
				body := `{}`
				w := utils.TestRequest(router, "POST", baseUrl, body)
				So(w.Code, ShouldEqual, 400)
				data := utils.Response2Dict(w.Body.Bytes())
				So(data["msg"], ShouldEqual, "parameter error")
			})
		})

		Convey("edit topic", func() {
			url := utils.JoinUrl(baseUrl, t1.Id.Hex())

			Convey("should success", func() {
				body := `{"title": "t111", "content": "t111t111t111"}`
				w := utils.TestRequest(router, "PUT", url, body)
				So(w.Code, ShouldEqual, 200)

				var topic models.Topic
				err := db.C("topics").FindId(t1.Id).One(&topic)
				So(err, ShouldBeNil)
				So(topic.Title, ShouldEqual, "t111")
				So(topic.Content, ShouldEqual, "t111t111t111")
				So(topic.UpdateAt, ShouldNotBeNil)
				So(topic.UpdateAt.After(*topic.CreateAt), ShouldBeTrue)
				So(topic.Deleted, ShouldBeFalse)
				So(topic.DeleteAt, ShouldBeNil)
			})

			Convey("should fail", func() {
				Convey("with empty body", func() {
					body := `{}`
					w := utils.TestRequest(router, "PUT", url, body)
					So(w.Code, ShouldEqual, 400)
					data := utils.Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "parameter error")
				})

				Convey("with wrong id", func() {
					dummyId := utils.JoinStrings("123", t1.Id.Hex()[3:])
					url = utils.JoinUrl(baseUrl, dummyId)
					body := `{"title": "t111", "content": "t111t111t111"}`
					w := utils.TestRequest(router, "PUT", url, body)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("delete topic", func() {
			Convey("should success", func() {
				url := utils.JoinUrl(baseUrl, t1.Id.Hex())
				w := utils.TestRequest(router, "DELETE", url, "")
				So(w.Code, ShouldEqual, 204)
				count, err := db.C("topics").Find(bson.M{"_id": t1.Id, "deleted": false}).Count()
				So(err, ShouldBeNil)
				So(count, ShouldEqual, 0)
			})

			Convey("should fail with wrong id", func() {
				dummyId := utils.JoinStrings("123", t1.Id.Hex()[3:])
				url := utils.JoinUrl(baseUrl, dummyId)
				w := utils.TestRequest(router, "DELETE", url, "")
				So(w.Code, ShouldEqual, 404)
			})
		})

		Reset(func() {
			utils.ClearAll(session)
		})
	})
}
