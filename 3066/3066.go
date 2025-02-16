package main

import "container/heap"

type MinHeap struct {
	Nums []int
	K    int
}

func (h MinHeap) Len() int {
	return len(h.Nums)
}
 
func (h MinHeap) Less(i,j int) bool{
	return h.Nums[i] < h.Nums[j]
}

func (h *MinHeap) Push(a interface{}) {
	h.Nums = append(h.Nums, a.(int))
}

func (h *MinHeap) Pop() interface{} {
	temp := h.Nums[len(h.Nums)-1]
	h.Nums = h.Nums[:len(h.Nums)-1]
	return temp
}

func (h MinHeap) Swap(a, b int){
	h.Nums[a], h.Nums[b] = h.Nums[b], h.Nums[a]
}

func minOperations(nums []int, k int) int {
	var minHeap MinHeap
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			heap.Push(&minHeap, nums[i])
		}
	}

	var count int
	for minHeap.Len() > 0 {
		a := heap.Pop(&minHeap).(int)
		if minHeap.Len() == 0 {
			count++
			return count
		}
		b := heap.Pop(&minHeap).(int)
		newNum := (a * 2) + b
		if newNum < k {
			heap.Push(&minHeap, newNum)
		}
		count++
	}
	return count
}
