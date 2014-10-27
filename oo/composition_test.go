package oo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Shape interface {
	numOfCorners() int
	equilateral() bool
}

type HasFourCorners struct{}

func (self HasFourCorners) numOfCorners() int {
	return 4
}

type HasThreeCorners struct{}

func (self HasThreeCorners) numOfCorners() int {
	return 3
}

type IsEquilateral struct{}

func (self IsEquilateral) equilateral() bool {
	return true
}

type IsntEquilateral struct{}

func (self IsntEquilateral) equilateral() bool {
	return false
}

type Square struct {
	HasFourCorners
	IsEquilateral
}

type Rectangle struct {
	HasFourCorners
	IsntEquilateral
}

type RightTriangle struct {
	HasThreeCorners
	IsntEquilateral
}

func TestComposition(t *testing.T) {
	var square Shape = Square{}
	assert.Equal(t, 4, square.numOfCorners())
	assert.Equal(t, true, square.equilateral())

	var rectangle Shape = Rectangle{}
	assert.Equal(t, 4, rectangle.numOfCorners())
	assert.Equal(t, false, rectangle.equilateral())

	var triangle Shape = RightTriangle{}
	assert.Equal(t, 3, triangle.numOfCorners())
	assert.Equal(t, false, triangle.equilateral())
}
