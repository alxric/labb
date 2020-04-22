package main

import "testing"

func Test_hello(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "Alex",
			Output: "Hello ALEX!",
		},
		{
			Input:  "simon",
			Output: "Good bye SIMON!",
		},
		{
			Input:  "viktor",
			Output: "",
		},
	}

	for _, tc := range tt {
		output := helloOrGoodBye(tc.Input)
		if output != tc.Output {
			t.Fatalf("Expected: %s, got: %s", tc.Output, output)
		}
	}
}
