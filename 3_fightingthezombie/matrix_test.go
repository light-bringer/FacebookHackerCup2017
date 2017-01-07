package main_test

import (
	"testing"

	matrix "github.com/rogerclotet/FacebookHackerCup2017/3_fightingthezombie"
	"github.com/stretchr/testify/assert"
)

func TestMatrixFor2Faces(t *testing.T) {
	m := matrix.NewMatrix(2)

	prob := m.Prob(2, 4)
	assert.InDelta(t, 0.25, prob, 1e-6)
}
