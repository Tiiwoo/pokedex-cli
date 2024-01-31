package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			"hello world",
			[]string{"hello", "world"},
		},
		{
			"HELLO world",
			[]string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal, got=%d, expected=%d", len(actual), len(cs.expected))
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}
}
