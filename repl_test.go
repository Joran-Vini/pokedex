package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    "  hello world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "foo   bar baz",
		expected: []string{"foo", "bar", "baz"},
	},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("for input '%s': expected length %d, but got length %d", 
        	c.input, len(c.expected), len(actual))
    		continue
			
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("for input '%s': expected word '%s' at index %d, but got '%s'", 
        		c.input, expectedWord, i, word)
				continue
			}
		}
	}
}
