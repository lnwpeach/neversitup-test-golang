package main

import (
	"testing"
)

// Write your test here
func Test_FindOddNumber(t *testing.T) {
	tests := []testSuite{
		{
			Label:  "test case 1",
			Input:  args{FindOddNumber: []int{7}},
			Expect: 7,
		},
		{
			Label:  "test case 2",
			Input:  args{FindOddNumber: []int{0}},
			Expect: 0,
		},
		{
			Label:  "test case 3",
			Input:  args{FindOddNumber: []int{1, 1, 2}},
			Expect: 2,
		},
		{
			Label:  "test case 4",
			Input:  args{FindOddNumber: []int{0, 1, 0, 1, 0}},
			Expect: 0,
		},
		{
			Label:  "test case 5",
			Input:  args{FindOddNumber: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}},
			Expect: 4,
		},
		{
			Label:  "test case 6",
			Input:  args{FindOddNumber: []int{1, 2, 2, 3, 3, 8, 3, 4, 3, 3, 3, 2, 8, 2, 1, 8}},
			Expect: 8,
		},
		{
			Label:  "test case 7",
			Input:  args{FindOddNumber: []int{}},
			Expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Label, func(t *testing.T) {
			if res := FindOddNumber(tt.Input.FindOddNumber); res != tt.Expect {
				t.Errorf("case %v failed -> expect: %v, actual: %v", tt.Label, tt.Expect, res)
			}
		})

	}
}
