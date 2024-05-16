package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXXX(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	var pos int = 5
	v1 := v[:pos]
	v1 = append(v1, v[pos+1:]...)
	assert.Equal(t, []int{1, 2, 3, 5}, v1)
}
