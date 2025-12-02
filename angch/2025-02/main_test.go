package main

import "testing"

func Test_dupe(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    string
		want bool
	}{
		{"11", "11", true},
		{"12", "12", false},
		{"22", "22", true},
		{"1010", "1010", true},
		{"1011", "1011", false},
		{"1188511885", "1188511885", true},
		{"446447", "446447", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dupe(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("dupe(%s) = %v, want %v", tt.a, got, tt.want)
			}
		})
	}
}

func Test_dupe2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    string
		want bool
	}{
		{"11", "11", true},
		{"12", "12", false},
		{"22", "22", true},
		{"111", "111", true},
		{"1010", "1010", true},
		{"1011", "1011", false},
		{"1188511885", "1188511885", true},
		{"123123123", "123123123", true},
		{"2121212121", "2121212121", true},
		{"446447", "446447", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dupe2(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("dupe(%s) = %v, want %v", tt.a, got, tt.want)
			}
		})
	}
}
