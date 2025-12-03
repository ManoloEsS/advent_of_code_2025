package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNumRepeat(t *testing.T) {
	var tests = []struct {
		name      string
		pattern   int
		numString int
		expected  bool
	}{
		{name: "single", pattern: 1, numString: 1, expected: false},
		{name: "single two digit", pattern: 1, numString: 11, expected: true},
		{name: "single three digit", pattern: 1, numString: 111, expected: false},
		{name: "single one digit false", pattern: 1, numString: 2, expected: false},
		{name: "single two digit false", pattern: 1, numString: 21, expected: false},
		{name: "single three digit false", pattern: 1, numString: 212, expected: false},
		{name: "two, two digit true", pattern: 11, numString: 11, expected: false},
		{name: "two, four digit true", pattern: 11, numString: 1111, expected: true},
		{name: "two, six digit true", pattern: 111, numString: 111111, expected: true},
		{name: "two, two digit false", pattern: 11, numString: 12, expected: false},
		{name: "two, four digit false", pattern: 11, numString: 1112, expected: false},
		{name: "two, six digit false", pattern: 11, numString: 111112, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repeats := checkNumRepeatsTwice(strconv.Itoa(tt.pattern), strconv.Itoa(tt.numString))
			assert.Equal(t, tt.expected, repeats)
		})
	}
}
