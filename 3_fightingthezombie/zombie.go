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
		result := fmt.Sprintf("Case #%d: %.06f", i+1, c.maxProbabilityOfKill())
		results = append(results, result)
	}
	return results
}

func parseInput(r io.Reader) []inputCase {
	var (
		t int
		s int
	)

	fmt.Fscanln(r, &t)

	cases := []inputCase{}
	for i := 0; i < t; i++ {
		c := inputCase{
			rolls: []DiceRoll{},
		}
		fmt.Fscan(r, &c.hp, &s)

		for j := 0; j < s; j++ {
			roll := DiceRoll{}
			var rollString string
			fmt.Fscan(r, &rollString)
			fmt.Sscanf(rollString, "%dd%d%d", &roll.Dice, &roll.Faces, &roll.Offset)
			c.rolls = append(c.rolls, roll)
		}

		cases = append(cases, c)
	}

	return cases
}

type inputCase struct {
	hp    int
	rolls []DiceRoll
}

func (c inputCase) maxProbabilityOfKill() float64 {
	var probability float64
	for _, roll := range c.rolls {
		rp := roll.ProbabilityOfKill(c.hp - roll.Offset)
		if rp > probability {
			probability = rp
		}
	}
	return probability
}

type DiceRoll struct {
	Dice   int
	Faces  int
	Offset int
}

func (d DiceRoll) ProbabilityOfKill(zombieHP int) float64 {
	if zombieHP <= 0 {
		return 1
	} else if d.Dice < 1 {
		return 0
	}

	var probability float64

	d.Dice--
	factor := 1 / float64(d.Faces)
	for damage := 1; damage <= d.Faces; damage++ {
		probability += d.ProbabilityOfKill(zombieHP - damage) * factor
	}

	return probability
}
