package fizzbuzz_test

import (
	"go101/fizzbuzz"
	"testing"
)

func TestFizzBuzzGiven1(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}

}
