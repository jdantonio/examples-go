package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanString(t *testing.T) {
	original := "abcdef"
	expected := "abcdef"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}

func TestOneSlashAndN(t *testing.T) {
	original := "abc\ndef"
	expected := "abc\ndef"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}

func TestTwoSlashes(t *testing.T) {
	original := "abc\\def"
	expected := "abc\\def"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}

func TestWrongOrder(t *testing.T) {
	original := "abc\n\\def"
	expected := "abc\n\\def"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}

func TestOneStrippedOut(t *testing.T) {
	original := "abc\\\ndef"
	expected := "abcdef"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}

func TestTwoStrippedOut(t *testing.T) {
	original := "ab\\\ncd\\\nef"
	expected := "abcdef"
	actual := Unsplice(original)
	assert.Equal(t, expected, actual)
}
