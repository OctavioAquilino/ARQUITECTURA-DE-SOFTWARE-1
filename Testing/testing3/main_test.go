package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		numerador   int
		denominador int
		expected    int
	}{
		{4, 1, 4},
		{-1, 0, 0},
		{0, 2, 0},
	}
	for _, test := range tests {
		div, err := division(test.numerador, test.denominador)

		if test.denominador == 0 {
			assert.Error(err, "No puedo dividir por 0")
			assert.Equal(test.expected, 0)
		} else {
			assert.Equal(div, test.expected)
			assert.Equal(err, nil)
		}
	}
}
