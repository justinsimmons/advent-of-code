package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mullItOver(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedTotal float64
		expectedError error
	}{

		{
			name: "example",
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]" +
				"then(mul(11,8)mul(8,5))",
			expectedTotal: 161,
			expectedError: nil,
		},
		{
			name:          "empty string",
			input:         "",
			expectedTotal: 0,
			expectedError: nil,
		},
		{
			name:          "should handle negative numbers",
			input:         "mul(-1,100)",
			expectedTotal: -100,
			expectedError: nil,
		},
		{
			name:          "should strip extraneous white space",
			input:         "mul(1, 8)",
			expectedTotal: 8,
			expectedError: nil,
		},
		{
			name:          "should handle garbage input",
			input:         "aslkd;ghnnfnbas;hfjgda;sdhgiha( askjdfd\nags\t124",
			expectedTotal: 0,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mullItOver(tc.input)

			assert.Equal(t, tc.expectedTotal, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
