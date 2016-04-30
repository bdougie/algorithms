package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeIsInitialized(t *testing.T) {
	list := []string{"C", "B", "F", "E", "G", "A", "D"}
	b := quickSort(list)
	assert.Equal(t, b, []string{"A", "B", "C", "D", "E", "F", "G"})
}
