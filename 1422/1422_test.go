package main

import "testing"

func Test_maxScore(t *testing.T) {
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
				s: "011101",
			},
			want: 5,
		},
		{
			name: "second",
			args: args{
				s: "00111",
			},
			want: 5,
		},
		{
			name: "third",
			args: args{
				s: "1111",
			},
			want: 3,
		},
		{
			name: "fourth",
			args: args{
				s: "00",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
