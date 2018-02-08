package longset_valid_parentheses

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testStartEnd(parentheses string, expectedStart, expectedEnd int) {
	start, end := GetLongestSubString(parentheses)
	So(start, ShouldEqual, expectedStart)
	So(end, ShouldEqual, expectedEnd)
}

func Test_GetLongestSubstring(t *testing.T) {

	Convey("Test GetLongestSubstring", t, func() {
		testStartEnd("", 0, 0)
		testStartEnd("()", 0, 1)
		testStartEnd(")", 0, 0)
		testStartEnd("(", 0, 0)
		testStartEnd("((()))", 0, 5)
		testStartEnd("()()", 0, 3)
		testStartEnd(")()", 1, 2)
		testStartEnd("()(", 0, 1)
		testStartEnd("((()())(())(", 1, 10)
	})
}
