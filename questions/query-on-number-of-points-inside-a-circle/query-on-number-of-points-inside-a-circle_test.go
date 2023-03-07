package leetcode

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	t.Parallel()
	type args struct {
		points  [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{}
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
