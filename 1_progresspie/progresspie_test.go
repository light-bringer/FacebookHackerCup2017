package main_test

import (
	"strings"
	"testing"

	progressPie "github.com/rogerclotet/FacebookHackerCup2017/1_progresspie"
)

func TestProgressPie(t *testing.T) {
	input := `7
0 55 55
12 55 55
13 55 55
99 99 99
87 20 40
0 50 50
1 50 50`

	results := progressPie.Run(strings.NewReader(input))

	expectations := []string{
		"Case #1: white",
		"Case #2: white",
		"Case #3: black",
		"Case #4: white",
		"Case #5: black",
		"Case #6: white",
		"Case #7: black",
	}

	if len(results) != len(expectations) {
		t.Errorf("invalid results amount: %d", len(results))
	}

	for i, r := range results {
		if r != expectations[i] {
			t.Errorf("\"%s\" does not match expectation \"%s\"", r, expectations[i])
		}
	}
}
