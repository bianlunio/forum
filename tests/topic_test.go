package tests

import (
	"strconv"
	"testing"

	"forum/models"
	"forum/routers"
	"forum/utils"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTopicRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routers.SetRouter()
	db := models.Connect()
	defer db.Close()

	Convey("TopicRouter", t, func() {
		baseUrl := "/api/topics"

		t1 := models.Topic{
			Title:   "t1",
			Content: "t1t1t1",
		}
		t2 := models.Topic{
			Title:   "t2",
			Content: "t2t2t2",
		}
		db.Model(&t1, &t2).Insert()

		Convey("list topics", func() {
			Convey("should success", func() {
				Convey("when default", func() {
					w := TestRequest(router, "GET", baseUrl, "")
					So(w.Code, ShouldEqual, 200)
					data := ListResponse2Dict(w.Body.Bytes())

					pagination := data.Pagination
					So(pagination["page"], ShouldEqual, 1)
					So(pagination["total"], ShouldEqual, 2)
					So(pagination["limit"], ShouldEqual, 20)

					list := data.List
					So(list, ShouldHaveLength, 2)
					So(list[0]["title"], ShouldEqual, "t1")
					So(list[0]["content"], ShouldEqual, "t1t1t1")
					So(list[1]["title"], ShouldEqual, "t2")
					So(list[1]["content"], ShouldEqual, "t2t2t2")
				})

				Convey("when limit = 1", func() {
					query := utils.Query{"limit": "1"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 200)
					data := ListResponse2Dict(w.Body.Bytes())

					pagination := data.Pagination
					So(pagination["page"], ShouldEqual, 1)
					So(pagination["total"], ShouldEqual, 2)
					So(pagination["limit"], ShouldEqual, 1)

					list := data.List
					So(list, ShouldHaveLength, 1)
					So(list[0]["title"], ShouldEqual, "t1")
					So(list[0]["content"], ShouldEqual, "t1t1t1")
				})
			})

			Convey("should fail", func() {
				Convey("with invalid page number", func() {
					query := utils.Query{"page": "0"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 400)
					data := Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "parameter error")
				})

				Convey("with invalid limit number", func() {
					query := utils.Query{"limit": "0"}
					url := utils.ParseQueryUrl(baseUrl, query)
					w := TestRequest(router, "GET", url, "")
					So(w.Code, ShouldEqual, 400)
					data := Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "parameter error")
				})
			})
		})

		Convey("get topic detail", func() {
			Convey("should success", func() {
				url := utils.JoinUrl(baseUrl, strconv.Itoa(t1.Id))
				w := TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 200)

				data := Response2Dict(w.Body.Bytes())
				So(data["id"], ShouldEqual, t1.Id)
				So(data["title"], ShouldEqual, "t1")
				So(data["content"], ShouldEqual, "t1t1t1")
			})

			Convey("should fail with wrong id", func() {
				dummyId := utils.JoinStrings("123", "12")
				url := utils.JoinUrl(baseUrl, dummyId)
				w := TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 404)
			})

			Convey("should fail with invalid id", func() {
				url := utils.JoinUrl(baseUrl, "someErrorId")
				w := TestRequest(router, "GET", url, "")
				So(w.Code, ShouldEqual, 400)
				data := Response2Dict(w.Body.Bytes())
				So(data["msg"], ShouldEqual, "parameter error")
			})
		})

		Convey("create topic", func() {
			Convey("should success", func() {
				body := `{"title": "t3", "content": "t3t3t3"}`
				w := TestRequest(router, "POST", baseUrl, body)
				So(w.Code, ShouldEqual, 201)
				data := Response2Dict(w.Body.Bytes())
				So(data["id"], ShouldNotBeNil)
				So(data["title"], ShouldEqual, "t3")
				So(data["content"], ShouldEqual, "t3t3t3")

				id := int(data["id"].(float64))
				topic, err := models.Topic{}.Detail(id)
				So(err, ShouldBeNil)
				So(topic.Id, ShouldEqual, 3)
				So(topic.Title, ShouldEqual, "t3")
				So(topic.Content, ShouldEqual, "t3t3t3")
				So(topic.Deleted, ShouldBeFalse)
				So(topic.DeletedAt.IsZero(), ShouldBeTrue)
			})

			Convey("should fail", func() {
				body := `{}`
				w := TestRequest(router, "POST", baseUrl, body)
				So(w.Code, ShouldEqual, 400)
				data := Response2Dict(w.Body.Bytes())
				So(data["msg"], ShouldEqual, "parameter error")
			})
		})

		Convey("edit topic", func() {
			url := utils.JoinUrl(baseUrl, strconv.Itoa(t1.Id))

			Convey("should success", func() {
				body := `{"title": "t111", "content": "t111t111t111"}`
				w := TestRequest(router, "PUT", url, body)
				So(w.Code, ShouldEqual, 200)

				topic, err := models.Topic{}.Detail(t1.Id)
				So(err, ShouldBeNil)
				So(topic.Title, ShouldEqual, "t111")
				So(topic.Content, ShouldEqual, "t111t111t111")
				So(topic.UpdatedAt, ShouldNotBeNil)
				So(topic.UpdatedAt.After(topic.CreatedAt), ShouldBeTrue)
				So(topic.Deleted, ShouldBeFalse)
				So(topic.DeletedAt.IsZero(), ShouldBeTrue)
			})

			Convey("should fail", func() {
				Convey("with empty body", func() {
					body := `{}`
					w := TestRequest(router, "PUT", url, body)
					So(w.Code, ShouldEqual, 400)
					data := Response2Dict(w.Body.Bytes())
					So(data["msg"], ShouldEqual, "parameter error")
				})

				Convey("with wrong id", func() {
					dummyId := utils.JoinStrings("123", "13")
					url = utils.JoinUrl(baseUrl, dummyId)
					body := `{"title": "t111", "content": "t111t111t111"}`
					w := TestRequest(router, "PUT", url, body)
					So(w.Code, ShouldEqual, 404)
				})
			})
		})

		Convey("delete topic", func() {
			Convey("should success", func() {
				url := utils.JoinUrl(baseUrl, strconv.Itoa(t1.Id))
				w := TestRequest(router, "DELETE", url, "")
				So(w.Code, ShouldEqual, 204)

				topic := models.Topic{Id: t1.Id}
				err := db.Select(&topic)
				So(err, ShouldBeNil)
				So(topic.Deleted, ShouldBeTrue)
				So(topic.DeletedAt.IsZero(), ShouldBeFalse)
			})

			Convey("should fail with wrong id", func() {
				dummyId := utils.JoinStrings("123", "1231")
				url := utils.JoinUrl(baseUrl, dummyId)
				w := TestRequest(router, "DELETE", url, "")
				So(w.Code, ShouldEqual, 404)
			})
		})

		Reset(func() {
			models.ClearAll()
		})
	})
}
