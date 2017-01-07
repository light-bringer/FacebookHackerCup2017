package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type inputCase struct {
	p int
	x int
	y int
}

func main() {
	results := Run(os.Stdin)
	for _, r := range results {
		fmt.Println(r)
	}
}

func Run(r io.Reader) []string {
	results := []string{}
	for i, c := range parseInput(r) {
		colour := "white"
		if c.coloured() {
			colour = "black"
		}

		result := fmt.Sprintf("Case #%d: %s", i+1, colour)
		results = append(results, result)
	}
	return results
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

const (
	size   = 100
	center = size / 2
)

// TODO punt central
func (c inputCase) coloured() bool {
	if c.x == center && c.y == center {
		return c.p > 0
	}

	if distanceFromCenter(c.x, c.y) > size/2 {
		return false
	}

	angle := c.angle()
	filledAngle := 360 * c.p / 100
	if angle > filledAngle {
		return false
	}

	return true
}

func (c inputCase) angle() int {
	x := float64(c.x - center)
	y := float64(c.y - center)
	dist := math.Sqrt(x*x + y*y)

	normX := x / dist
	degrees := int(math.Acos(normX) * 180 / math.Pi)

	var angle int
	if y >= 0 {
		angle = 90 - degrees
	} else {
		angle = 90 + degrees
	}

	if angle < 0 {
		angle += 360
	}

	return angle
}

func distanceFromCenter(x, y int) int {
	a := float64(center - x)
	b := float64(center - y)
	return int(math.Sqrt(a*a + b*b))
}
