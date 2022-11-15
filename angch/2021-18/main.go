package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// https://adventofcode.com/2021/day/18

type SNum struct {
	L, R   int
	Lp, Rp *SNum
}

func AddSNum(s1 *SNum, s2 *SNum) *SNum {
	return &SNum{
		0, 0, s1, s2,
	}
}

func (p *SNum) Reduce() *SNum {
	action := true

	for action {
		action, _, _ = p.Explode(0)
		if action {
			continue
		}
		p, action = p.Split()
	}

	return p
}

func Split(i int) *SNum {
	if i >= 10 {
		return &SNum{
			i / 2, i/2 + i%2, nil, nil,
		}
	}
	return nil
}

func (s *SNum) Split() (*SNum, bool) {
	action := false
	if s.Lp == nil {
		if s.L >= 10 {
			s.Lp = Split(s.L)
			s.L = 0
			action = true
		}
	} else {
		var x *SNum
		x, action = s.Lp.Split()
		if action {
			s.Lp = x
			return s, action
		}
	}
	if action {
		return s, action
	}

	if s.Rp == nil {
		if s.R >= 10 {
			s.Rp = Split(s.R)
			s.R = 0
			action = true
		}
	} else {
		var x *SNum
		x, action = s.Rp.Split()
		if action {
			s.Rp = x
			return s, action
		}
	}
	return s, action
}

func (p *SNum) AddExplodeRight(i int) {
	if p.Rp == nil {
		p.R += i
	} else {
		p.Rp.AddExplodeRight(i)
	}
}

func (p *SNum) AddExplodeLeft(i int) {
	if p.Lp == nil {
		p.L += i
	} else {
		p.Lp.AddExplodeLeft(i)
	}
}

func (p *SNum) Explode(level int) (bool, int, int) {
	if level == 4 {
		return true, p.L, p.R
	}

	if p.Lp != nil {
		// Left explosion
		ex, l, r := p.Lp.Explode(level + 1)
		if ex {
			if r >= 0 && l >= 0 {
				p.Lp = nil
				if p.Rp == nil {
					p.R += r
				} else {
					p.Rp.AddExplodeLeft(r)
				}
			} else if r > 0 {
				if p.Rp == nil {
					p.R += r
				} else {
					p.Rp.AddExplodeLeft(r)
				}
			}
			return true, l, -1
		}
	}

	if p.Rp != nil {
		// Right explosion, so left is added
		ex, l, r := p.Rp.Explode(level + 1)
		if ex {
			if l >= 0 && r >= 0 {
				p.Rp = nil
				p.R = 0
				if p.Lp == nil {
					p.L += l
				} else {
					p.Lp.AddExplodeRight(l)
				}
			} else if l > 0 {
				if p.Lp == nil {
					p.L += l
				} else {
					p.Lp.AddExplodeRight(l)
				}
			}
			return true, -1, r
		}
	}
	return false, 0, 0
}

func ParseSNum(line string) (*SNum, int) {
	p := &SNum{}
	last := 0
	count := 0
a:
	for ; last < len(line) && (count == 0 || last != 0); last++ {
		switch line[last] {
		case '[':
			count++
		case ']':
			count--
			if count == 0 {
				break a
			}
		}
	}

	if line[0] == '[' && line[last] == ']' {
		l := line[1:last]
		last2 := 0
		if l[0] == '[' {
			p.Lp, last2 = ParseSNum(l)
			last2++
		} else {
			for last2 = 0; last2 < len(l); last2++ {
				if l[last2] == ',' {
					break
				}
			}
			if last2 == len(l) {
			} else {
				var err error
				p.L, err = strconv.Atoi(l[:last2])
				if err != nil {
					log.Fatal(err)
				}
				for ; last2 < len(l); last2++ {
					if l[last2] != ']' {
						break
					}
				}
			}
		}
		if last2 < len(l) {
			l = l[last2+1:]
			if l[0] == '[' {
				// last3 := 0
				p.Rp, _ = ParseSNum(l)
				// last = last2 + last3
			} else {
				fmt.Sscanf(l, "%d", &p.R)
			}
		}
	} else {
		log.Println("Fatal")
	}

	return p, last
}

func testparsesnum(a string) string {
	p, _ := ParseSNum(a)
	out := p.Print()
	fmt.Printf("%s = %s", a, out)
	if out != a {
		fmt.Println(" FAIL")
	} else {
		fmt.Println(" PASS")
	}
	return out
}

func testaddsum(a, b string, expect string) string {
	a1, _ := ParseSNum(a)
	b1, _ := ParseSNum(b)
	c := AddSNum(a1, b1)
	out := c.Print()
	fmt.Printf("%s + %s = %s", a, b, out)
	if out != expect {
		fmt.Println(" FAIL")
	} else {
		fmt.Println(" PASS")
	}
	return out
}

func testaddreduce(a, b string, expect string) string {
	a1, _ := ParseSNum(a)
	b1, _ := ParseSNum(b)
	c := AddSNum(a1, b1)
	c = c.Reduce()
	out := c.Print()
	fmt.Printf("%s + %s = %s", a, b, out)
	if out != expect {
		fmt.Println(" FAIL")
	} else {
		fmt.Println(" PASS")
	}
	return out
}

func testexplode(input string, expect string) string {
	a, _ := ParseSNum(input)
	fmt.Print(a.Print(), " -> ")
	exp, _, _ := a.Explode(0)
	out := a.Print()
	fmt.Print(out)

	if out != expect {
		fmt.Println(" FAIL", exp)
	} else {
		fmt.Println(" PASS", exp)
	}

	return out
}

func testsplit(input string, expect string) string {
	a, _ := ParseSNum(input)
	fmt.Print(a.Print(), " -> ")
	_, exp := a.Split()
	out := a.Print()
	fmt.Print(out)

	if out != expect {
		fmt.Println(" FAIL", exp)
	} else {
		fmt.Println(" PASS", exp)
	}

	return out
}

func testmagnitude(input string, expect int) {
	a, _ := ParseSNum(input)
	out := a.Magnitude()
	if out != expect {
		fmt.Println(" FAIL")
	} else {
		fmt.Println(" PASS")
	}
}

func (s *SNum) Magnitude() int {
	l, r := 0, 0
	if s.Lp != nil {
		l = s.Lp.Magnitude()
	} else {
		l = s.L
	}
	if s.Rp != nil {
		r = s.Rp.Magnitude()
	} else {
		r = s.R
	}
	return 3*l + 2*r
}

func (s *SNum) DeepClone() *SNum {
	s1 := &SNum{
		s.L,
		s.R,
		nil,
		nil,
	}

	if s.Lp != nil {
		s1.Lp = s.Lp.DeepClone()
	}
	if s.Rp != nil {
		s1.Rp = s.Rp.DeepClone()
	}

	return s1
}

func day18test(filepath string) {
	if true {
		testparsesnum("[1,2]")
		testparsesnum("[[[[4,3],4],4],[7,[[8,4],9]]]")
		testparsesnum("[[[[0,7],4],[15,[0,13]]],[1,1]]")
	}
	if true {
		testaddsum("[1,2]", "[[3,4],5]", "[[1,2],[[3,4],5]]")
		testaddsum("[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	}
	if true {
		testexplode("[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]")
		testexplode("[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]")
		testexplode("[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]")
		testexplode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
		testexplode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
		testexplode("[[[[0,9],2],3],4]", "[[[[0,9],2],3],4]")
	}
	if true {
		a := "[[[[4,3],4],4],[7,[[8,4],9]]]"
		b := "[1,1]"
		a1, _ := ParseSNum(a)
		b1, _ := ParseSNum(b)
		c := AddSNum(a1, b1)

		testexplode(c.Print(), "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]")
		c.Explode(0)
		testexplode(c.Print(), "[[[[0,7],4],[15,[0,13]]],[1,1]]")
		fmt.Println(c.Print(), "z")
		c.Explode(0)
		fmt.Println(c.Print(), "x")
		testexplode(c.Print(), "[[[[0,7],4],[15,[0,13]]],[1,1]]")

		testsplit(c.Print(), "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")

		testsplit("[0,13]", "[0,[6,7]]")
		testsplit("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
		testsplit("[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")

		testsplit("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
		testexplode("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
		// c.Explode(0)

	}

	if true {
		testaddreduce("[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	}

	if true {
		testmagnitude("[9,1]", 29)
		testmagnitude("[[9,1],[1,9]]", 129)
	}

	return
}

func day18(filepath string) {
	if true {
		file, err := os.Open(filepath)
		if err != nil {
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		snums := make([]*SNum, 0)
		scanner.Scan()
		t := scanner.Text()
		p, _ := ParseSNum(t)
		snums = append(snums, p.DeepClone())
		verbose := false

		for scanner.Scan() {
			t := scanner.Text()
			p1, _ := ParseSNum(t)
			snums = append(snums, p1.DeepClone())

			if verbose {
				fmt.Printf(" %s\n+ %s\n", p.Print(), p1.Print())
			}
			p = AddSNum(p, p1)
			p = p.Reduce()
			if verbose {
				fmt.Printf("= %s\n\n", p.Print())
			}
		}
		fmt.Println("Part 1", p.Magnitude(), p.Print())

		// Part 2, len of snums is 100, so we're doing 10000 - 100 evaluations
		largest, large1, large2 := 0, 0, 0
		for i := 0; i < len(snums); i++ {
			for j := 0; j < len(snums); j++ {
				if i == j {
					continue
				}
				p = AddSNum(snums[i].DeepClone(), snums[j].DeepClone())
				p = p.Reduce()
				m := p.Magnitude()
				if m > largest {
					largest, large1, large2 = m, i, j
				}
			}
		}
		fmt.Println("Part 2", largest, large1, large2)
	}
}

func (p *SNum) Print() string {
	a := "["
	if p.Lp != nil {
		a += p.Lp.Print()
	} else {
		a += fmt.Sprintf("%d", p.L)
	}
	a += ","
	if p.Rp != nil {
		a += p.Rp.Print()
	} else {
		a += fmt.Sprintf("%d", p.R)
	}
	a += "]"
	return a
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// For the purpose of running on machines without `time` aka Windows
	t1 := time.Now()
	// day18("test.txt")
	// day18("test2.txt")
	day18("input.txt")
	d1 := time.Since(t1)
	fmt.Println("Duration", d1)
}
