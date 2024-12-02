package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_historianHysteria(t *testing.T) {
	testCases := []struct {
		name     string
		listOne  []int
		listTwo  []int
		distance int
	}{
		{
			name:     "example",
			listOne:  []int{3, 4, 2, 1, 3, 3},
			listTwo:  []int{4, 3, 5, 3, 9, 3},
			distance: 11,
		},
		{
			name:     "left list larger than right",
			listOne:  []int{3, 4, 2, 1, 3, 3, 2, 7},
			listTwo:  []int{4, 3, 5, 3, 9, 3},
			distance: 24,
		},
		{
			name:     "right list larger than left",
			listOne:  []int{3, 4, 2, 1, 3, 3},
			listTwo:  []int{4, 3, 5, 3, 9, 3, 12, 7},
			distance: 30,
		},
		{
			name:     "both lists empty",
			listOne:  []int{},
			listTwo:  []int{},
			distance: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := historianHysteria(tc.listOne, tc.listTwo)
			assert.Equal(t, tc.distance, result)
		})
	}
}
