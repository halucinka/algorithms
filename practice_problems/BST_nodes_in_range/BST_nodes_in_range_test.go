package BST_nodes_in_range

// https://practice.geeksforgeeks.org/problems/count-bst-nodes-that-lie-in-a-given-range/1

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetNumbersInRange_recursive(t *testing.T) {
	Convey("Test GetNumbersInRange_recursive", t, func() {
		root := BuildBinarySearchTree([]int{10, 5, 50, 1, 40, 100})

		So(GetNumbersInRange_recursive(0, 10, root), ShouldResemble, []int{1, 5, 10})
		So(GetNumbersInRange_recursive(5, 45, root), ShouldResemble, []int{5, 10, 40})
	})
	Convey("Test GetNumbersInRange_recursive - with repeating numbers", t, func() {
		root := BuildBinarySearchTree([]int{10, 10, 10, 10, 10})

		So(GetNumbersInRange_recursive(0, 10, root), ShouldResemble, []int{10, 10, 10, 10, 10})
		So(GetNumbersInRange_recursive(9, 11, root), ShouldResemble, []int{10, 10, 10, 10, 10})
		So(GetNumbersInRange_recursive(10, 11, root), ShouldResemble, []int{10, 10, 10, 10, 10})
		So(GetNumbersInRange_recursive(10, 10, root), ShouldResemble, []int{10, 10, 10, 10, 10})
		So(GetNumbersInRange_recursive(17, 47, root), ShouldResemble, []int{})
	})
}
