package main

import (
	"testing"

	"github.com/Pizzu/pokedexcli/common"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := common.CleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Slices are different! \n Expected: %v \n Received: %v", len(c.expected), len(actual))
			t.Fail()
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words are not matching! \n Expected: %s \n Received: %s", expectedWord, word)
				t.Fail()
			}
		}
	}

}
