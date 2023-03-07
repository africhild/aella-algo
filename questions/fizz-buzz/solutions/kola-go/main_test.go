package main

import (
	"reflect"
	"testing"
)


func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				n: 15,
			},
			want: []string{"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"},
		},
		{
			name: "Example 2",
			args: args{
				n: 5,
			},
			want: []string{"1","2","Fizz","4","Buzz"},
		},
		{
			name: "Example 3",
			args: args{
				n: 3,
			},
			want: []string{"1","2","Fizz"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := FizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
