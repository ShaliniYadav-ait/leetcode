package main

func constructDistancedSequence(n int) []int {
	size := n*2 - 1
	result := make([]int, size)
	isUsed := make([]bool, n+1)

	backtracking(0, result, isUsed, n)
	return result
}

func backtracking(index int, result []int, isUsed []bool, n int) bool {
	if index == len(result) {
		return true
	}

	for num := n; num >= 1; num-- {
		if isUsed[num] || (num > 1 && (num+index >= len(result)) || (result[num+index] > 0)) {
			continue
		}

		isUsed[num] = true
		result[index] = num
		if num > 1 {
			result[index+num] = num
		}
		j := index + 1
		for j < len(result) && result[j] > 0 {
			j++
		}
		if backtracking(j, result, isUsed, n) {
			return true
		}

		isUsed[num] = false
		result[index] = 0
		if num > 1 {
			result[index+num] = 0
		}
	} 
	return false 
}
