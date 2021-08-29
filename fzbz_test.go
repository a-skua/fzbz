package fzbz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
		{19, "19"},
		{20, "Buzz"},
	}

	for _, tt := range tests {
		actual := FizzBuzz(tt.num)

		if actual != tt.expected {
			t.Fatalf("wrong return value %d. want=%s, got=%s", tt.num, tt.expected, actual)
		}
	}
}
