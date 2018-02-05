package naive

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testMatching(text, pattern string, expectedMatchingResponse bool, expectedPosition int) {
	matchingResponse, position := MatchString(text, pattern)
	So(matchingResponse, ShouldEqual, expectedMatchingResponse)
	So(position, ShouldEqual, expectedPosition)
}

func Test_StringMatching_NaiveImplementation(t *testing.T) {
	Convey("Test StringMatching_NaiveImplementation", t, func() {
		testMatching("textsomething", "text", true, 0)
		testMatching("", "text", false, -1)
		testMatching("somethingtexttext", "text", true, 9)
		testMatching("textsomething", "", true, 0)
	})
}
