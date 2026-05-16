package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sum(a, b int) int {
	return a + b
}

func Test_sum(t *testing.T) {
	assert.Equal(t, 4, sum(2, 2))
}
