package rabin_karp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testMatching(text, pattern string, expectedMatchingResponse bool, expectedPosition int) {
	matchingResponse, position := MatchString(text, pattern)
	So(matchingResponse, ShouldEqual, expectedMatchingResponse)
	So(position, ShouldEqual, expectedPosition)
}

func Test_StringMatching_RabinKarp(t *testing.T) {
	Convey("Test RabinKarp", t, func() {
		testMatching("textsomething", "text", true, 0)
		testMatching("", "text", false, -1)
		testMatching("textsomething", "", true, 0)
		testMatching("text", "text", true, 0)
		testMatching("somethingtext", "text", true, 9)
		testMatching("somethingtexttext", "text", true, 9)
	})
}
