package main

type ProductOfNumbers struct {
	Nums []int
	Zero map[int][]int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		Nums: []int{},
		Zero: make(map[int][]int),
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.Zero[0] = append(this.Zero[0], len(this.Nums))
		this.Nums = append(this.Nums, num)
		return
	}
	lastProd := 1
	if len(this.Nums) > 0 {
		lastProd = this.Nums[len(this.Nums)-1]
		if lastProd == 0 {
			lastProd = 1
		} 
	}
	currProd := lastProd * num
	this.Nums = append(this.Nums, currProd)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	last := len(this.Nums) - 1
	first := last - k + 1
	if val, ok := this.Zero[0]; ok {
		for i := 0; i < len(val); i++ {
			if val[i] >= this.Nums[first] && val[i] <= this.Nums[last] {
				return 0
			}
		}
	}
	var prodK int
	if first-1 >= 0 && this.Nums[first-1] > 0 {
		prodK = this.Nums[last] / this.Nums[first-1]
	}
  return prodK
}
