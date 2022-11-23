package main

import (
	"testing"
)

func Test_countLit(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{`""`}, 2, 0},
		{"2", args{`"abc"`}, 5, 3},
		{"3", args{`"aaa\"aaa"`}, 10, 7},
		{"4", args{`"\x27"`}, 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countLit(tt.args.a)
			if got != tt.want {
				t.Errorf("countLit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countLit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_countEnc(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{`""`}, 2, 6},
		{"2", args{`"abc"`}, 5, 9},
		{"3", args{`"aaa\"aaa"`}, 10, 16},
		{"4", args{`"\x27"`}, 6, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countEnc(tt.args.a)
			if got != tt.want {
				t.Errorf("countEnc() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countEnc() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
