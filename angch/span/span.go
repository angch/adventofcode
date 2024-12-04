package span

type Span[T any] struct {
	From    int
	To      int
	Content T
}

type Spans[T any] []Span[T]

func NewSpans[T any]() Spans[T] {
	return make(Spans[T], 0, 8)
}

func (s *Spans[T]) Add(from, to int, content T) Spans[T] {
	from, to = min(from, to), max(from, to)
	return append(*s, Span[T]{from, to, content})
}

// Initial conversion from 2022-15 code to type param'd Spans
