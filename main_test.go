package main

import (
	"testing"
)

func TestfindKthLargest(t *testing.T) {
	out := findKthLargest([]int{99, 1, 2, 2, 3, 5, 7, 99, 2, 100}, 3)
	want := 99
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
