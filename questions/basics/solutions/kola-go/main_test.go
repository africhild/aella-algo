package main

import (
	"testing"
)

func TestSumAll(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "Example 2",
			args: args{
				n: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 55,
		},
		{
			name: "Example 3",
			args: args{
				n: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			want: 210,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SumAll(tt.args.n); got != tt.want {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "Example 2",
			args: args{
				s: "Howdy",
			},
			want: "ydwoH",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
