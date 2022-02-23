package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		input  string
		want   string
		length int
	}{
		{
			input:  "Freddie Mercury",
			want:   "Queen",
			length: 1,
		},
		{
			input:  "My Chemical Romance",
			want:   "",
			length: 0,
		},
	}

	for _, tc := range testCases {
		result := Search(tc.input)
		if len(result) != tc.length {
			t.Logf("tc.input = %s", tc.input)
			t.Fatalf("len(result) is %d, expected a length of %d", len(result), tc.length)
		}
		if len(result) > 0 {
			if result[0].Name != tc.want {
				t.Fatalf("result[0].Name failed, wanted %s, got %v", tc.want, result[0].Name)
			}
		}
	}
}

// func TestHandle500(t *testing.T) {
// }

// func TestHandle400(t *testing.T) {
// }
