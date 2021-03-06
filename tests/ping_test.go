package tests

import (
	"testing"

	"forum/routers"
	"forum/utils"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPingRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routers.SetRouter()

	Convey("PingRouter", t, func() {
		Convey("ping", func() {
			w := utils.TestRequest(router, "GET", "/ping", "")
			So(w.Code, ShouldEqual, 200)
			So(w.Body.String(), ShouldEqual, "pong")
		})
	})
}
