package main

import (

	"fmt"
	"testing"
)
var zeroMap = make(map[int][]int)
func TestProductOfNumbers_Add(t *testing.T) {
	type fields struct { 
		Nums []int
		Zero map[int][]int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{
			name: "add number 3",
			fields: fields{
				Nums: []int{1, 2},
				Zero: zeroMap,
			},
			args: args{
				num: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{
				Nums: tt.fields.Nums,
			}
			this.Add(tt.args.num)
			fmt.Println(tt.fields.Nums)
		})
	}
}

func TestProductOfNumbers_GetProduct(t *testing.T) {
	zeroMap[0] = append(zeroMap[0], 1)
	type fields struct {
		Nums []int
		Zero map[int][]int
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "get product",
			fields: fields{
				Nums: []int{1, 2, 6},
				Zero : zeroMap,
			},
			args: args{
				k: 2,
			},
			want: 6,
		},
		{
			name: "get product 2nd",
			fields: fields{
				Nums: []int{3,0,2,10,40,320},
				Zero : zeroMap,
			},
			args: args{
				k: 3,
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{
				Nums: tt.fields.Nums,
			}
			if got := this.GetProduct(tt.args.k); got != tt.want {
				t.Errorf("ProductOfNumbers.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}