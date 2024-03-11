package main

func Permutations(str string) []string {
	// TODO : start your code here
	if len(str) == 1 {
		return []string{str}
	}

	dupMap := make(map[string]struct{})
	permutations := []string{}
	for i, c := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := Permutations(remaining)

		for _, p := range subPermutations {
			if _, ok := dupMap[string(string(c)+p)]; !ok {
				permutations = append(permutations, string(c)+p)
				dupMap[string(c)+p] = struct{}{}
			}
		}
	}

	return permutations
}
