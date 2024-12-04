package span

import (
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
