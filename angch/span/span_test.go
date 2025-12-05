package span

import (
	"log"
	"reflect"
	"testing"
)

func TestAddInt(t *testing.T) {
	tests := []struct {
		name string
		args Span[int]
		want Spans[int]
	}{
		{
			name: "Reversed",
			args: Span[int]{From: 10, To: 5, Content: 2},
			want: Spans[int]{{From: 5, To: 10, Content: 2}},
		},
		{
			name: "Expand left",
			args: Span[int]{From: 10, To: 4, Content: 2},
			want: Spans[int]{{From: 4, To: 10, Content: 2}},
		},
		{
			name: "Expand right",
			args: Span[int]{From: 6, To: 12, Content: 2},
			want: Spans[int]{{From: 5, To: 12, Content: 2}},
		},
		{
			name: "No change",
			args: Span[int]{From: 6, To: 8, Content: 2},
			want: Spans[int]{{From: 5, To: 10, Content: 2}},
		},
		{
			name: "Expand both",
			args: Span[int]{From: 4, To: 12, Content: 2},
			want: Spans[int]{{From: 4, To: 12, Content: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpans[int]()
			s = s.Add(5, 10, 2)
			got := s.AddCompress(tt.args.From, tt.args.To, tt.args.Content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpans_Contains(t *testing.T) {

	// Test spans based on AoC 2025-05
	spans := NewSpans[bool]()
	spans = spans.AddCompress(3, 5, true)
	spans = spans.AddCompress(10, 14, true)
	spans = spans.AddCompress(16, 20, true)
	spans = spans.AddCompress(12, 18, true)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i    int
		want bool
	}{
		{"1", 1, false},
		{"5", 5, true},
		{"8", 8, false},
		{"11", 11, true},
		{"17", 17, true},
		{"32", 32, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(spans)
			got := spans.Contains(tt.i)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
