package main

import (
	"fmt"
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	type testcase struct {
		array []int
		key int
		expected int
	}

	testcases := []testcase{
		{[]int{1,2,3,4,5,6,7,8,9,10}, 7, 6},
		{[]int{1,2,3,4,5,6,7,8,9,10}, 3, 2},
		{[]int{1, 2, 10, 21, 44, 56, 102}, 21, 3},
		{[]int{1,3,7,10,14,15,16,18,20,21,22,23,25,33,35,42,45,47,50,52}, 33, 13},
		{[]int{1,3,7,10,14,15,16,18,20,21,22,23,25,33,35,42,45,47,50,52}, 50, 18},
		{[]int{1,3,7,10,14,15,16,18,20,21,22,23,25,33,35,42,45,47,50,52}, 18, 7},
	}

	for i, v := range testcases{
		aim := InterpolationSearch(v.array, v.key)
		if aim == v.expected{
			fmt.Println("Test case num", i ," passed successfully with values: Expected: ", v.expected, " Got: ", aim)
		} else{
			t.Fatal("Test case num", i ," failed with values: Expected: ", v.expected, " Got: ", aim)
		}
	}
}
