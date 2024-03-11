package main

type testSuite struct {
	Label  string
	Input  args
	Expect any
}

type args struct {
	permutations   string
	FindOddNumber  []int
	CountSmilyFace []string
}
