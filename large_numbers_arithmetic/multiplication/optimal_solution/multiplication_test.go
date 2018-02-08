package optimal_solution

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Multiply(t *testing.T) {
	Convey("Test Multiply", t, func() {
		So(Multiply("100", "100"), ShouldEqual, "10000")
		So(Multiply("1", "100"), ShouldEqual, "100")
		So(Multiply("3", "13"), ShouldEqual, "39")
		So(Multiply("3", "14"), ShouldEqual, "42")
		So(Multiply("12", "12"), ShouldEqual, "144")
		So(Multiply("0", "120"), ShouldEqual, "0")
	})
}
