package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_redNosedReports(t *testing.T) {
	testCases := []struct {
		name    string
		reports string
		numSafe int
	}{
		{
			name: "example",
			reports: "7 6 4 2 1\n" +
				"1 2 7 8 9\n" +
				"9 7 6 2 1\n" +
				"1 3 2 4 5\n" +
				"8 6 4 4 1\n" +
				"1 3 6 7 9",
			numSafe: 2,
		},
		{
			name:    "no reports",
			reports: "",
			numSafe: 0,
		},
		{
			name:    "not enough levels to determine safety",
			reports: "1\n2\n3\n4",
			numSafe: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := redNosedReports(tc.reports)
			assert.Equal(t, tc.numSafe, result)
		})
	}
}
