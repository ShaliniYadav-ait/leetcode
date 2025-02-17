package main

import "testing"

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "first",
			args : args{
				tiles : "AAB",
			},
			want : 8,
		},
		{
			name : "second",
			args : args{
				tiles : "AAABBC",
			},
			want : 188,
		},
		{
			name : "third",
			args : args{
				tiles : "V",
			},
			want : 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
