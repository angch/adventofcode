package main

import "testing"

type tcase struct {
	input  string
	expect bool
}

var tcases = []tcase{
	tcase{"ugknbfddgicrmopn", true},
	tcase{"aaa", true},
	tcase{"jchzalrnumimnmhp", false},
	tcase{"haegwjzuvuyypxyu", false},
	tcase{"dvszwmarrgswjxmb", false},
}
var tcases2 = []tcase{
	tcase{"qjhvhtzxzqqjkmpb", true},
	tcase{"xxyxx", true},
	tcase{"uurcxstgmygtbstg", false},
	tcase{"ieodomkazucvgmuy", false},
}

func TestOne(t *testing.T) {
	for _, test := range tcases {
		a := advent05(test.input)
		if a != test.expect {
			t.Error(test.input, "should be", test.expect, "but got", a)
		}
	}
}

func TestTwo(t *testing.T) {
	for _, test := range tcases2 {
		a := advent05b(test.input)
		if a != test.expect {
			t.Error(test.input, "should be", test.expect, "but got", a)
		}
	}
}
