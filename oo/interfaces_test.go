package oo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type englishSpeaker interface {
	Hello() string
}

type englishSpeakerImpl struct{}

func (self englishSpeakerImpl) Hello() string {
	return "Hello"
}

type portugueseSpeaker interface {
	Ola() string
}

type portugueseSpeakerImpl struct{}

func (self portugueseSpeakerImpl) Ola() string {
	return "Olá"
}

type polyglotSpeaker struct {
	englishSpeaker
	portugueseSpeaker
}

func TestInterfaces(t *testing.T) {
	speaker := polyglotSpeaker{
		englishSpeaker:    englishSpeakerImpl{},
		portugueseSpeaker: portugueseSpeakerImpl{},
	}
	assert.Equal(t, "Hello", speaker.Hello())
	assert.Equal(t, "Olá", speaker.Ola())
}
