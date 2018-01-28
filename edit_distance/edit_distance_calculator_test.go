package edit_distance

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func makeRuneArrays(a, b string) ([]rune, []rune) {
	return []rune(a), []rune(b)
}

func testMinAndPosition(calculator editDistanceCalculator, firstParameter, secondParameter, thirdParameter, expectedMinimum, expectedPosition int) {
	minimum, position := calculator.Min3WithPosition(firstParameter, secondParameter, thirdParameter)
	So(minimum, ShouldEqual, expectedMinimum)
	So(position, ShouldEqual, expectedPosition)
}

func TestEditDistanceCalculator_Calculate(t *testing.T) {
	Convey("Test Calculator", t, func() {
		calculator := NewEditDistanceCalculator()

		So(calculator.Calculate(makeRuneArrays("", "")), ShouldEqual, 0)
		So(calculator.Calculate(makeRuneArrays("a", "")), ShouldEqual, 1)
		So(calculator.Calculate(makeRuneArrays("", "a")), ShouldEqual, 1)
		So(calculator.Calculate(makeRuneArrays("a", "b")), ShouldEqual, 1)
		So(calculator.Calculate(makeRuneArrays("aa", "abc")), ShouldEqual, 2)
		So(calculator.Calculate(makeRuneArrays("abc", "aa")), ShouldEqual, 2)
	})
}

func TestEditDistanceCalculator_Max3(t *testing.T) {
	Convey("Test Max3", t, func() {
		calculator := editDistanceCalculator{}

		So(calculator.Min3(1, 2, 3), ShouldEqual, 1)
		So(calculator.Min3(2, 1, 3), ShouldEqual, 1)
		So(calculator.Min3(1, 3, 2), ShouldEqual, 1)
		So(calculator.Min3(2, 3, 1), ShouldEqual, 1)
		So(calculator.Min3(3, 1, 2), ShouldEqual, 1)
		So(calculator.Min3(3, 2, 1), ShouldEqual, 1)
		So(calculator.Min3(1, 1, 3), ShouldEqual, 1)
	})
}

func TestEditDistanceCalculator_GetShortestequenceOfOperations(t *testing.T) {
	Convey("Test GetShortestequenceOfOperations", t, func() {
		calculator := NewEditDistanceCalculator()

		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("", "")), ShouldBeEmpty)
		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("a", "")), ShouldHaveLength, 1)
		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("", "a")), ShouldHaveLength, 1)
		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("a", "b")), ShouldHaveLength, 1)
		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("aa", "abc")), ShouldHaveLength, 2)
		So(calculator.GetShortestSequenceOfOperations(makeRuneArrays("abc", "aa")), ShouldHaveLength, 2)
	})
}

func TestEditDistanceCalculator_Max3WithPosition(t *testing.T) {
	Convey("Test Max3WithPosition", t, func() {
		calculator := editDistanceCalculator{}

		testMinAndPosition(calculator, 2, 1, 3, 1, 2)
		testMinAndPosition(calculator, 1, 2, 3, 1, 1)
		testMinAndPosition(calculator, 1, 3, 2, 1, 1)
		testMinAndPosition(calculator, 2, 3, 1, 1, 3)
		testMinAndPosition(calculator, 3, 1, 2, 1, 2)
		testMinAndPosition(calculator, 3, 2, 1, 1, 3)
		testMinAndPosition(calculator, 1, 1, 3, 1, 1)
	})
}
