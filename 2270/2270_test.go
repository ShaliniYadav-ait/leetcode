package main

import "testing"

func Test_waysToSplitArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				nums: []int{10, 4, -8, 7},
			},
			want: 2,
		},
		{
			name: "second",
			args: args{
				nums: []int{2,3,1,0},
			},
			want: 2,
		},

		{
			name: "third",
			args: args{
				nums: []int{0,0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// {
// 	name : ,
// args : args{
// 	nums : []int{},
// },
// want : ,
// }
