package main

import (
	"fmt"
	"io"
	"os"
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
		result := fmt.Sprintf("Case #%d: %.06f", i+1, c.probabilityOfDefeat())
		results = append(results, result)
	}
	return results
}

func parseInput(r io.Reader) []inputCase3 {
	var (
		t int
		s int
	)

	fmt.Fscanln(r, &t)

	cases := []inputCase3{}
	for i := 0; i < t; i++ {
		c := inputCase3{
			rolls: []diceRoll{},
		}
		fmt.Fscan(r, &c.h, &s)

		for j := 0; j < s; j++ {
			roll := diceRoll{}
			var rollString string
			fmt.Fscan(r, &rollString)
			fmt.Sscanf(rollString, "%dd%d%d", &roll.dice, &roll.sides, &roll.offset)
			c.rolls = append(c.rolls, roll)
		}

		cases = append(cases, c)
	}

	return cases
}

type diceRoll struct {
	dice   int
	sides  int
	offset int
}

type inputCase3 struct {
	h     int
	rolls []diceRoll
}

func (c inputCase3) probabilityOfDefeat() float64 {
	return 0
}
