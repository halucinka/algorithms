package jumping_numbers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetJumpingNumbers(t *testing.T) {

	Convey("Test GetJumpingNumbers", t, func() {
		So(GetJumpingNumbers(0), ShouldResemble, []int{0})
		So(GetJumpingNumbers(1), ShouldResemble, []int{0, 1})
		So(GetJumpingNumbers(22), ShouldResemble, []int{0, 1, 10, 12, 2, 21, 3, 4, 5, 6, 7, 8, 9})
		So(GetJumpingNumbers(10), ShouldResemble, []int{0, 1, 10, 2, 3, 4, 5, 6, 7, 8, 9})
		So(GetJumpingNumbers(50), ShouldResemble, []int{0, 1, 10, 12, 2, 21, 23, 3, 32, 34, 4, 43, 45, 5, 6, 7, 8, 9})
	})
}
