package main

import (
	"reflect"
	"testing"
)

func Test_hasInc(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"hijklmmn"}, true},
		{"2", args{"abbceffg"}, false},
		{"3", args{"abbcegjk"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasInc(tt.args.s); got != tt.want {
				t.Errorf("hasInc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"hijklmmn"}, false},
		{"2", args{"abbceffg"}, false},
		{"3", args{"abbcegjk"}, false},
		{"4", args{"abcdffaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incPassword(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte("abcd")}, []byte("abce")},
		{"2", args{[]byte("abcz")}, []byte("abda")},
		// {"2", args{[]byte("abcdefgh")}, []byte("abcdffaa")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incPassword(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abcdefgh"}, "abcdffaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incValid(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
