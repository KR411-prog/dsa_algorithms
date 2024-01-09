package main

import "testing"
func Test_binarysearch(t *testing.T) {
	input := []int{2, 3, 4, 5, 6,7,9}
	target := 5

	if binarysearch(input,target) != 3 {
		t.Errorf("expected not same as output. expected=%d, input=%d", 3, binarysearch(input,target))
	}
}