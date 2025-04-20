package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "", expected: []string{}},
		{input: "HELLO", expected: []string{"hello"}},
		{input: "hello world", expected: []string{"hello", "world"}},
		{input: "hello world!", expected: []string{"hello", "world!"}},
		{input: "hello world!", expected: []string{"hello", "world!"}},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		for i, v := range actual {
			if i >= len(tc.expected) || v != tc.expected[i] {
				t.Errorf("cleanInput(%q) = %v, want %v", tc.input, actual, tc.expected)
				break
			}
		}
		if len(actual) != len(tc.expected) {
			t.Errorf("cleanInput(%q) = %v, want %v", tc.input, actual, tc.expected)
		}
	}
}
