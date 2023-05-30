package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	//There are 2 ways to initialize an array:
	//way 1:
	// numbers := [5]int{1, 2, 3, 4, 5}

	//way 2:
	//numbers := [...]int{1, 2, 3, 4, 5}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers) //%v is a specifier for the default format which works well for Arrays
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		//This is a slice
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) { //Best way to compare 2 slices is with reflect. However, DeepEqual will compile even if the 2 things you're comparing aren't of the same type
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	//You can assign a function to a variable much like you can assign a string or int to it
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) { // a cool thing with passing the checkSum responsability to a function is that you can now delegate the type checking responsibility to the func.
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
