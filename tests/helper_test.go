package tests

import (
	"testing"

	"forum/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelper(t *testing.T) {
	Convey("Helper", t, func() {
		Convey("String2Int", func() {
			n := utils.String2Int("121")
			So(n, ShouldEqual, 121)
		})

		Convey("JoinStrings", func() {
			s := utils.JoinStrings("ab", "Cd", "121")
			So(s, ShouldEqual, "abCd121")
		})
	})
}
