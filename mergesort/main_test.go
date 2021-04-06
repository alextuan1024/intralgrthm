package mergesort

import "testing"

func Test_mergesort(t *testing.T) {
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
			name: "m1",
			args: args{
				array: []int{3, 41, 52, 26, 38, 57, 9, 49},
				p:     0,
				r:     7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergesort(tt.args.array, tt.args.p, tt.args.r)
			t.Log(tt.args.array)
		})
	}
}
