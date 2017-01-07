package main_test

import (
	"strings"
	"testing"

	lazyLoading "github.com/rogerclotet/FacebookHackerCup2017/2_lazyloading"
	"github.com/stretchr/testify/assert"
)

func TestLazyLoading(t *testing.T) {
	input := `5
4
30
30
1
1
3
20
20
20
11
1
2
3
4
5
6
7
8
9
10
11
6
9
19
29
39
49
59
10
32
56
76
8
44
60
47
85
71
91`

	results := lazyLoading.Run(strings.NewReader(input))

	expectations := []string{
		"Case #1: 2",
		"Case #2: 1",
		"Case #3: 2",
		"Case #4: 3",
		"Case #5: 8",
	}

	assert.Len(t, results, len(expectations))

	for i, r := range results {
		assert.Equal(t, expectations[i], r)
	}
}

func TestCase1(t *testing.T) {
	c := lazyLoading.InputCase{30, 30, 1, 1}

	assert.Equal(t, 2, c.MinTrips())
}

func TestCase3(t *testing.T) {
	c := lazyLoading.InputCase{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	assert.Equal(t, 2, c.MinTrips())
}

func TestCase5(t *testing.T) {
	c := lazyLoading.InputCase{
		32,
		56,
		76,
		8,
		44,
		60,
		47,
		85,
		71,
		91,
	}

	assert.Equal(t, 8, c.MinTrips())
}
