package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var list = []string{"A", "B", "C", "D", "E", "F", "G"}

func TestLinearSearchReturnsCount(t *testing.T) {
	c := LinearSearchIterations(&list, "G")
	assert.Equal(t, c, len(list), "iterations equals the length of list")
}

func TestBinarySearchReturnsCount(t *testing.T) {
	c := BinarySearchIterations(&list, "G")
	assert.Equal(t, c, 3, "iteerations equal half the length of list")
}

func TestBinarySearchReturnsMedian(t *testing.T) {
	c := BinarySearchIterations(&list, "D")
	assert.Equal(t, c, 1, "iteerations equal half the length of list")
}

func TestLinearSearchReturnsWorksForMedian(t *testing.T) {
	c := LinearSearchIterations(&list, "D")
	assert.Equal(t, c, 4)
}

func TestBinarySearchReturnsB(t *testing.T) {
	c := BinarySearchIterations(&list, "B")
	assert.Equal(t, c, 2)
}
