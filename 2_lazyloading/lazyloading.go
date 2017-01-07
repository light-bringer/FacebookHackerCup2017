package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	results := Run(os.Stdin)
	for _, r := range results {
		fmt.Println(r)
	}
}

func Run(r io.Reader) []string {
	results := []string{}
	for i, c := range parseInput(r) {
		result := fmt.Sprintf("Case #%d: %d", i+1, c.MinTrips())
		results = append(results, result)
	}
	return results
}

func parseInput(r io.Reader) []InputCase {
	var (
		t int
		n int
	)

	fmt.Fscanln(r, &t)

	cases := []InputCase{}
	for i := 0; i < t; i++ {
		fmt.Fscanln(r, &n)

		c := InputCase{}
		for j := 0; j < n; j++ {
			var w int
			fmt.Fscanln(r, &w)
			c = append(c, w)
		}

		cases = append(cases, c)
	}

	return cases
}

type InputCase []int

const minWeight = 50

func (c InputCase) MinTrips() int {
	sort.Ints(c)

	trips := 0
	t := trip{c.pop()}

	for len(c) > 0 {
		if t.value() >= minWeight {
			trips++
			t = trip{c.pop()}
		} else {
			t = append(t, c.shift())
		}
	}

	// If last trip is heavy enough, he can take an extra trip
	if t.value() >= minWeight {
		trips++
	}

	return trips
}

func (c *InputCase) shift() int {
	var w int
	w, *c = (*c)[0], (*c)[1:]
	return w
}

func (c *InputCase) pop() int {
	var w int
	w, *c = (*c)[len(*c) - 1], (*c)[:len(*c) - 1]
	return w
}

type trip []int

func (t trip) value() int {
	return t[0] * len(t)
}
