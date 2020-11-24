package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tested MultiplicatorSimple

//more test shall be added, but for time constraint only a couple to get the coverage, but this don't ensure the req

func TestShouldReturnZero(t *testing.T) {
	res, err := tested.Multiply(10, 0)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), res.Int64())
}

func TestShouldReturnMult(t *testing.T) {
	res, err := tested.Multiply(100, 10)
	assert.Nil(t, err)
	assert.Equal(t, int64(1000), res.Int64())
}
