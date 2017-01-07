package main

import (
	"fmt"
	"io"
	"os"
)

type inputCase struct {
	p int
	x int
	y int
}

func parseInput(r io.Reader) []inputCase {
	var amount int
	fmt.Fscanln(r, &amount)

	cases := []inputCase{}
	for i := 0; i < amount; i++ {
		c := inputCase{}
		fmt.Fscanln(r, &c.p, &c.x, &c.y)
		cases = append(cases, c)
	}

	return cases
}

func (c inputCase) colour() string {
	return "black"
}

func Run(r io.Reader) []string {
	results := []string{}
	for i, c := range parseInput(r) {
		result := fmt.Sprintf("Case #%d: %s", i+1, c.colour())
		results = append(results, result)
	}
	return results
}

func main() {
	results := Run(os.Stdin)
	for _, r := range results {
		fmt.Println(r)
	}
}
