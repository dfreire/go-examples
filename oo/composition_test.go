package oo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foo struct{}

func (self Foo) Ping() string {
	return "Pong"
}

func (self Foo) Hello() string {
	return "World"
}

type Bar struct {
	Foo
}

func (self Bar) Hello() string {
	return "World!"
}

func TestComposition(t *testing.T) {
	foo := Foo{}
	assert.Equal(t, "Pong", foo.Ping())
	assert.Equal(t, "World", foo.Hello())

	bar := Bar{}
	assert.Equal(t, "Pong", bar.Ping())
	assert.Equal(t, "World!", bar.Hello())
}
