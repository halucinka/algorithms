package edit_distance

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func makeRuneArrays(a, b string) ([]rune, []rune){
    return []rune(a), []rune(b)
}


func TestEditDistanceCalculator_Calculate(t *testing.T) {
    Convey("Test Calculator", t, func() {
        calculator := NewEditDistanceCalculator()

        So(calculator.Calculate(makeRuneArrays("","")), ShouldEqual, 0)
        So(calculator.Calculate(makeRuneArrays("a","")), ShouldEqual, 1)
        So(calculator.Calculate(makeRuneArrays("","a")), ShouldEqual, 1)
        So(calculator.Calculate(makeRuneArrays("a","b")), ShouldEqual, 1)
        So(calculator.Calculate(makeRuneArrays("aa","abc")), ShouldEqual, 2)
        So(calculator.Calculate(makeRuneArrays("abc","aa")), ShouldEqual, 2)
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
