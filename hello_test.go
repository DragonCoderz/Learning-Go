package main

import "testing"

//Testing Function Semantics
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
func TestHello(t *testing.T) { //t of type *testing.T is our door into the testing package which gives us
	//access to all sorts of testing functions like FAIL

	//variables are declared by :=
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
