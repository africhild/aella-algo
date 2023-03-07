package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			main()
		})
	}
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				input: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				input: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				input: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				input: "",
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				input: "aab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := LongestSubstringWithoutRepeatingCharacters(tt.args.input); got != tt.want {
				t.Errorf("LongestSubstringWithoutRepeatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
