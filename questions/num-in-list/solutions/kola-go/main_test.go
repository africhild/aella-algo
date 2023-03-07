package main

import "testing"

func TestNumInList(t *testing.T) {
	type args struct {
		list []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				list: []int{2, 7, 9, 11, 15},
				num:  9,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				list: []int{3, 2, 6, 4},
				num:  6,
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				list: []int{3,3},
				num:  6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NumInList(tt.args.list, tt.args.num); got != tt.want {
				t.Errorf("NumInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
