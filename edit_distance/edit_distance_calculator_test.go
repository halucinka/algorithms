package edit_distance

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func makeRuneArrays(a, b string) ([]rune, []rune) {
	return []rune(a), []rune(b)
}

func makeRuneArray(a string) []rune {
	return []rune(a)
}

func testMinAndPosition(calculator editDistanceCalculator, firstParameter, secondParameter, thirdParameter, expectedMinimum, expectedPosition int) {
	minimum, position := calculator.Min3WithPosition(firstParameter, secondParameter, thirdParameter)
	So(minimum, ShouldEqual, expectedMinimum)
	So(position, ShouldEqual, expectedPosition)
}

func testSequenceOfOperations(calculator EditDistanceCalculator, a, b []rune, expectedFirstWord, expectedSecondWord string) {
	firstWord, secondWord := calculator.GetOperations(a, b)
	So(firstWord, ShouldEqual, expectedFirstWord)
	So(secondWord, ShouldEqual, expectedSecondWord)
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

func TestEditDistanceCalculator_GetOperations(t *testing.T) {
	Convey("Test GetOperations", t, func() {
		calculator := NewEditDistanceCalculator()

		testSequenceOfOperations(calculator, makeRuneArray(""), makeRuneArray(""), "", "")
		testSequenceOfOperations(calculator, makeRuneArray("a"), makeRuneArray(""), "a", "-")
		testSequenceOfOperations(calculator, makeRuneArray(""), makeRuneArray("a"), "-", "a")
		testSequenceOfOperations(calculator, makeRuneArray("a"), makeRuneArray("b"), "a", "b")
		testSequenceOfOperations(calculator, makeRuneArray("aa"), makeRuneArray("abc"), "aa-", "abc")
		testSequenceOfOperations(calculator, makeRuneArray("abc"), makeRuneArray("aa"), "abc", "aa-")
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
