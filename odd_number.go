package main

func FindOddNumber(ints []int) int {
	// TODO : start your code here
	if len(ints) == 0 {
		return 0
	}

	countMap := make(map[int]int)
	orderArr := []int{}
	for _, v := range ints {
		if _, ok := countMap[v]; !ok {
			orderArr = append(orderArr, v)
		}
		countMap[v]++
	}
	for _, num := range orderArr {
		count := countMap[num]
		if count%2 == 1 {
			return num
		}
	}

	return 0
}
