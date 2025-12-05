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

func (s Spans[T]) Add(from, to int, content T) Spans[T] {
	from, to = min(from, to), max(from, to)
	newspan := append(s, Span[T]{from, to, content})
	return newspan
}

func (s Spans[T]) AddCompress(from, to int, content T) Spans[T] {
	// FIXME, this func hasn't been fixed to make sure content is the same before merging
	from, to = min(from, to), max(from, to)

	if len(s) == 0 {
		return append(s, Span[T]{from, to, content})
	}

	var start int
	// Insertion sort
	for i := 0; i < len(s); i++ {
		if s[i].From > from {
			// Check for special case where we are within the previous span
			if i > 0 && s[i-1].To >= to {
				// It is within the previous span, let's just go home
				return s
			}
			if i > 0 && s[i-1].To >= from && to < s[i].From {
				// New span only extends to left span, not touching the next span's left
				s[i-1].To = to
				return s
			}

			s = append(s, Span[T]{})
			copy(s[i+1:], s[i:])
			s[i] = Span[T]{from, to, content}

			// Slower, as it forces an alloc
			// The above avoids an alloc
			// *s = append((*s)[:i], append([]Span{{l, r}}, (*s)[i:]...)...)
			start = i + 1 // FIXME: We might have an off by one error
			goto compress
		}
	}

	start = len(s)
	s = append(s, Span[T]{from, to, content})

compress:
	for i := start; i > 0; i-- {
		j := i - 1
		a, b := s[j], s[i]
		if a.To >= b.From-1 {
			a.To = max(b.To, a.To)

			c := 1
			i2 := i + 1
			for ; i2 < len(s); i2, c = i2+1, c+1 {
				if a.To < s[i2].To {
					if a.To >= s[i2].From {
						a.To = s[i2].To
						continue
					}
					break
				}
			}
			s[j] = a

			copy(s[i:], s[i2:])
			s = s[:len(s)-c]
		}
	}
	return s
}

func (s Spans[T]) Contains(i int) bool {
	// binary search
	l, r := 0, len(s)
	for l < r {
		mid := l + (r-l)/2

		if mid >= len(s) {
			return false
		}

		diff := s[mid].From - i
		if diff <= 0 && s[mid].To >= i {
			return true
		}
		if diff <= 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

// Initial conversion from 2022-15 code to type param'd Spans
