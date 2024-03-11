package main

import (
	"reflect"
	"testing"
)

// Write your test here
func Test_Permutations(t *testing.T) {
	tests := []testSuite{
		{
			Label:  "test case 1",
			Input:  args{permutations: "a"},
			Expect: []string{"a"},
		},
		{
			Label:  "test case 2",
			Input:  args{permutations: "ab"},
			Expect: []string{"ab", "ba"},
		},
		{
			Label:  "test case 3",
			Input:  args{permutations: "abc"},
			Expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			Label:  "test case 4",
			Input:  args{permutations: "aabb"},
			Expect: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			Label: "test case 5",
			Input: args{permutations: "abcd"},
			Expect: []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad",
				"cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
		},
		{
			Label:  "test case 6",
			Input:  args{permutations: ""},
			Expect: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Label, func(t *testing.T) {
			if res := Permutations(tt.Input.permutations); !reflect.DeepEqual(res, tt.Expect) {
				t.Errorf("case %v failed -> expect: %v, actual: %v", tt.Label, tt.Expect, res)
			}
		})

	}
}
