package main_test

import (
	"strings"
	"testing"

	zombie "github.com/rogerclotet/FacebookHackerCup2017/3_fightingthezombie"
	"github.com/stretchr/testify/assert"
)

func TestFightingTheZombie(t *testing.T) {
	input := `5
2 2
2d4 1d8
10 2
10d6-10 1d6+1
8 3
1d4+4 2d4 3d4-4
40 3
10d4 5d8 2d20
10 4
1d10 1d10+1 1d10+2 1d10+3`

	results := zombie.Run(strings.NewReader(input))

	expectations := []string{
		"Case #1: 1.000000",
		"Case #2: 0.998520",
		"Case #3: 0.250000",
		"Case #4: 0.002500",
		"Case #5: 0.400000",
	}

	if len(results) != len(expectations) {
		t.Errorf("invalid results amount: %d", len(results))
	}

	for i, r := range results {
		assert.Equal(t, expectations[i], r)
	}
}
