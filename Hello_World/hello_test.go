package main

import "testing"

//Testing Function Semantics
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
func TestHello(t *testing.T) { //t of type *testing.T is our door into the testing package which gives us
	//access to all sorts of testing functions like FAIL

	//Here is how you do subtests - subtests should be designed to test some specific aspect of our code
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "") //variables are declared by := but are changed with =
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	//passing the *testing.T object to t.Run, you enable the subtest to communicate its failures, log additional information,
	//and control its execution. This helps in providing accurate test results and facilitating the debugging process.
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean-Marie", "French")
		want := "Bonjour, Jean-Marie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // t being of type testing.TB means that it is able to call both Testing methods and Benchmark methods
	t.Helper() //Signfities that this is a helper function to our compiler
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
