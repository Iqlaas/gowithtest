package main

import "testing"

func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func(t *testing.T) {

	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got '%d' but want '%d', '%v'", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' but '%d' given, '%v'", got, want, numbers)
		}
	})

}
