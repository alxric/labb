package main

import "testing"

func Test_hello(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "alex",
			Output: "Hello alex!",
		},
		{
			Input:  "Simon",
			Output: "Hello Simon!",
		},
	}

	for _, tc := range tt {
		output := hello(tc.Input)
		if output != tc.Output {
			t.Fatalf("Expected: %s, got: %s", tc.Output, output)
		}
	}
}
