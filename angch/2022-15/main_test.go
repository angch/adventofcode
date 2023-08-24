package main

import "testing"

func TestCompare(t *testing.T) {
	type args struct {
		a Span
		b Span
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{Span{0, 0}, Span{0, 0}}, 0},
		{"1", args{Span{1, 0}, Span{0, 0}}, 1},
		{"-1", args{Span{2, 14}, Span{12, 12}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
