package main

import "testing"

func Test_quicksort(t *testing.T) {
	type args struct {
		array []int
		p     int
		r     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "q1",
			args: args{
				array: []int{2, 8, 7, 1, 3, 5, 6, 4},
				p:     0,
				r:     7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.array, tt.args.p, tt.args.r)
			t.Log(tt.args.array)
		})
	}
}
