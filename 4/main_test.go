package main

import "testing"

func Test_hello(t *testing.T) {
	tt := []struct {
		Number1 int
		Number2 int
		Largest int
	}{
		{
			Number1: 1,
			Number2: 2,
			Largest: 2,
		},
		{
			Number1: 1,
			Number2: 0,
			Largest: 1,
		},
		{
			Number1: 10,
			Number2: -10,
			Largest: 10,
		},
		{
			Number1: 20,
			Number2: 20,
			Largest: 20,
		},
	}

	for _, tc := range tt {
		largest := largestNumber(tc.Number1, tc.Number2)
		if largest != tc.Largest {
			t.Fatalf("Expected: %d, got: %d", tc.Largest, largest)
		}
	}
}
