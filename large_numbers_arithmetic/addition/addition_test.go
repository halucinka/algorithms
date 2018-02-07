package addition

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Add(t *testing.T) {
	Convey("Test Add", t, func() {
		So(Add("100", "100"), ShouldEqual, "200")
		So(Add("1", "100"), ShouldEqual, "101")
		So(Add("999", "999"), ShouldEqual, "1998")
		So(Add("10", "1"), ShouldEqual, "11")
		So(Add("999", "1"), ShouldEqual, "1000")
	})
}
