package dayXX_test

import (
	"aoc/dayXX"
	"testing"

	A "github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	a := A.New(t)
	input := dayXX.ReadInput()
	a.Equal("", dayXX.Step1(input))
	a.Equal("", dayXX.Step2(input))
}
