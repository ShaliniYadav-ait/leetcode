package main

func waysToSplitArray(nums []int) int {
	totalSum := getSum(nums)
	currSum := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		currSum += nums[i]
		sumRight := totalSum - currSum
		if currSum >= sumRight {
			count++
		}
	}
	return count
}

func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
