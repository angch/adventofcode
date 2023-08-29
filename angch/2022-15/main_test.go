package main

import (
	"reflect"
	"testing"
)

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

func TestSpans_Compress(t *testing.T) {
	tests := []struct {
		name string
		s    Spans
		want Spans
	}{
		{"1", Spans{{0, 0}}, Spans{{0, 0}}},
		{"2", Spans{{0, 10}, {9, 12}}, Spans{{0, 12}}},
		{"3", Spans{{0, 20}, {9, 12}, {14, 15}}, Spans{{0, 20}}},
		{"4", Spans{{0, 20}, {9, 12}, {14, 15}, {23, 24}}, Spans{{0, 20}, {23, 24}}},
		{"5", Spans{{0, 20}, {9, 12}, {14, 15}, {19, 24}}, Spans{{0, 24}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Compress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spans.Compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
