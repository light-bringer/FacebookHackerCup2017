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

type DiceRoll struct {
	Dice   int
	Faces  int
	Offset int
}

type inputCase struct {
	hp    int
	rolls []DiceRoll
}

func (c inputCase) maxProbabilityOfKill() float64 {
	var probability float64
	for _, roll := range c.rolls {
		m := NewMatrix(roll.Faces)
		rp := m.Prob(roll.Dice, c.hp-roll.Offset)
		if rp > probability {
			probability = rp
		}
	}
	return probability
}

type Matrix struct {
	Faces int
	M     map[int]map[int]float64
}

func NewMatrix(faces int) Matrix {
	return Matrix{
		Faces: faces,
		M:     make(map[int]map[int]float64),
	}
}

func (m *Matrix) Prob(dice, hp int) float64 {
	p, ok := m.M[dice][hp]
	if !ok {
		p = m.calc(dice, hp)
		if _, ok := m.M[dice]; !ok {
			m.M[dice] = make(map[int]float64)
		}
		m.M[dice][hp] = p
	}

	return p
}

func (m *Matrix) calc(dice, hp int) float64 {
	if hp <= 0 {
		return 1
	} else if dice < 1 {
		return 0
	}

	var probability float64
	factor := 1 / float64(m.Faces)
	for damage := 1; damage <= m.Faces; damage++ {
		probability += m.Prob(dice-1, hp-damage) * factor
	}

	return probability
}
