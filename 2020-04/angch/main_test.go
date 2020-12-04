package main

import "testing"

func Test_checkvalid(t *testing.T) {
	tests := []struct {
		name string
		k    string
		v    string
		want bool
	}{
		// {"a", "byr", "2002", true},
		// {"a", "byr", "2003", false},
		// {"a", "hgt", "60in", true},
		// {"a", "hgt", "190cm", true},
		// {"a", "hgt", "190in", false},
		// {"a", "hgt", "190", false},
		// {"a", "hcl", "#123abc", true},
		// {"a", "hcl", "#123abz", false},
		// {"a", "hcl", "123abc", false},
		// {"a", "ecl", "brn", true},
		// {"a", "ecl", "wat", false},
		// {"a", "pid", "000000001", true},
		{"a", "pid", "0123456789", false},
		{"b", "byr", "1989", true},
		{"b", "cid", "129", true},

		//  cid:129 ecl:blu eyr:2029 hcl:#a97842 hgt:165cm iyr:2014 pid:896056539]

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := make(map[string]string)
			kv[tt.k] = tt.v
			if got := checkvalid(kv); got != tt.want {
				t.Errorf("checkvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}
