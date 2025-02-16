package main

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				s: "aabca",
			},
			want: 3,
		},
		{
			name: "second",
			args: args{
				s: "adc",
			},
			want: 0,
		},
		{
			name: "third",
			args: args{
				s: "bbcbaba",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

// {
// 	name : ,
// 	args : args{
// 		s : ,
// 	},
// 	want : ,
// 	},
