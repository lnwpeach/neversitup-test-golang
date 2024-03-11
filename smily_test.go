package main

import (
	"testing"
)

// Write your test here
func Test_CountSmilyFace(t *testing.T) {
	tests := []testSuite{
		{
			Label:  "test case 1",
			Input:  args{CountSmilyFace: []string{":)", ";(", ";}", ":-D"}},
			Expect: 2,
		},
		{
			Label:  "test case 2",
			Input:  args{CountSmilyFace: []string{";D", ":-(", ":-)", ";~)"}},
			Expect: 3,
		},
		{
			Label:  "test case 3",
			Input:  args{CountSmilyFace: []string{";]", ":[", ";*", ":$", ";-D"}},
			Expect: 1,
		},
		{
			Label:  "test case 4",
			Input:  args{CountSmilyFace: []string{":)", ";(", ";}", ":-D", ";D", ":-(", ":-)", ";~)", ";]", ":[", ";*", ":$", ";-D", ";~~D", ";~D"}},
			Expect: 7,
		},
		{
			Label:  "test case 5",
			Input:  args{CountSmilyFace: []string{}},
			Expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Label, func(t *testing.T) {
			if res := CountSmilyFace(tt.Input.CountSmilyFace); res != tt.Expect {
				t.Errorf("case %v failed -> expect: %v, actual: %v", tt.Label, tt.Expect, res)
			}
		})

	}
}
