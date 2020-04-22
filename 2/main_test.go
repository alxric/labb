package main

import "testing"

func Test_helloCaps(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "ALEX",
			Output: "HELLO ALEX!",
		},
		{
			Input:  "simon",
			Output: "HELLO SIMON!",
		},
		{
			Input:  "vIkToR",
			Output: "HELLO VIKTOR!",
		},
	}

	for _, tc := range tt {
		output := helloCaps(tc.Input)
		if output != tc.Output {
			t.Fatalf("Expected: %s, got: %s", tc.Output, output)
		}
	}
}
