package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fourth",
			args: args{
				nums: []int{97,73,5,78},
				k:    98,
			},
			want: 3,
		},
		{
			name: "first",
			args: args{
				nums: []int{2, 11, 10, 1, 3},
				k:    10,
			},
			want: 2,
		},
		{
			name: "second",
			args: args{
				nums: []int{1, 1, 2, 4, 9},
				k:    20,
			},
			want: 4,
		},
		{
			name: "third",
			args: args{
				nums: []int{999999999,999999999,999999999},
				k:    1000000000,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

// {
// 	name : ,
// args : args{
// 	nums : ,
// 	k : ,
// },
// want : ,
// 	},
