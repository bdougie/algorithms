package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeIsInitialized(t *testing.T) {
	b := fib(10)
	assert.Equal(t, b, 55)
}
