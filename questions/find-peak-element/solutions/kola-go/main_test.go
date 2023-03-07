package main

import (
	"reflect"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
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
				n: []int{1, 2, 3, 4, 2, 1},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				n: []int{1,2,1,3,5,6,4},
			},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{
				n: []int{ 3, 4, 5, 1, 2 },
			},
			want: 2,

		},
		//{
		//	name: "Example 3",
		//	args: args{
		//		n: []int{ 1, 2, 1, 1, 1 },
		//	},
		//	want: 1,
		//
		//},
		{
			name: "Example 4",
			args: args{
				n: []int{ 1, 2, 4, 6, 6, 6, 7 },
			},
			want: 6,

		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findPeakElement(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
