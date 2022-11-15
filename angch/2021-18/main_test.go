package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParsePair(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *SNum
	}{
		{"[1,2]", args{"[1,2]"}, &SNum{1, 2, nil, nil}},
		{"[[1,2],3]", args{"[[1,2],3]"}, &SNum{0, 3, &SNum{1, 2, nil, nil}, nil}},
		{"[9,[8,7]]", args{"[9,[8,7]]"}, &SNum{9, 0, nil, &SNum{8, 7, nil, nil}}},
		{"[[1,9],[8,5]]", args{"[[1,9],[8,5]]"}, &SNum{0, 0, &SNum{1, 9, nil, nil}, &SNum{8, 5, nil, nil}}},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", args{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
			&SNum{0, 9, &SNum{0, 0, &SNum{0, 0, &SNum{1, 2, nil, nil}, &SNum{3, 4, nil, nil}}, &SNum{0, 0, &SNum{5, 6, nil, nil}, &SNum{7, 8, nil, nil}}}, nil}},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]", args{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
			&SNum{0, 0,
				&SNum{0, 0,
					&SNum{9, 0,
						nil,
						&SNum{3, 8, nil, nil},
					},
					&SNum{0, 6, &SNum{0, 9, nil, nil}, nil},
				},
				&SNum{0, 3,
					&SNum{0, 0,
						&SNum{3, 7, nil, nil},
						&SNum{4, 9, nil, nil},
					}, nil,
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ParseSNum(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePair() = \n got %+v\nwant %+v\n%s\n", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
