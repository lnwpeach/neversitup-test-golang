package main

import (
	"regexp"
)

func CountSmilyFace(strs []string) int {
	// TODO : start your code here
	re, _ := regexp.Compile("^[:;][-~]?[)D]$")
	count := 0
	for _, v := range strs {
		match := re.MatchString(v)
		if match {
			count++
		}
	}

	return count
}
