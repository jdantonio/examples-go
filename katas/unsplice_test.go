package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_life_the_universe_and_everything(t *testing.T) {
	assert.Equal(t, 42, answer(), "No wonder the universe is so screwed up!")
}