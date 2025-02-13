package main

func maxScore(s string) int {
	totalOne := getTotal(s)
	maxScore, currScore := 0, 0
	leftZero, leftOne := 0, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			leftZero++
		} else {
			leftOne++
		}
		currScore = leftZero + totalOne - leftOne
		maxScore = max(currScore, maxScore)
	}
	return maxScore
}

func getTotal(s string) int {
	totalOne := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			totalOne++
		}
	}
	return totalOne
}
