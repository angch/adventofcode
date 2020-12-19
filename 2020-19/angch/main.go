package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Subrules [][]int
	Literal  string
}

func match(l string, rules []Rule, ii []int) bool {
	for _, i := range ii {
		r := rules[i]

		if r.Literal != "" {
			if l[0] == r.Literal[0] {
				l = l[1:]
				continue
			} else {
				return false
			}
		}

		// for subrule := r.Subrules {

		// }

	}
	return true
}

var rules_ = map[int]Rule{}
var rulesMap = map[int]string{}

func compile(i int) string {
	// regex := ""
	r, ok := rulesMap[i]
	if ok {
		return r
	}

	if rules_[i].Literal != "" {
		rulesMap[i] = rules_[i].Literal
		return rules_[i].Literal
	}
	subrules := make([]string, 0)

	hasPlus := false

a:
	for _, v := range rules_[i].Subrules {
		subsubrules := make([]string, 0)
		for _, v2 := range v {
			if v2 == i {
				hasPlus = true
				break a
			}
			subsubrule := compile(v2)
			subsubrules = append(subsubrules, subsubrule)
		}
		subrule := "(" + strings.Join(subsubrules, "") + ")"
		subrules = append(subrules, subrule)
	}
	r = "(" + strings.Join(subrules, "|") + ")"

	if hasPlus {
		// subrules = make([]string, 0)
		// for k, v := range rules_[i].Subrules {
		// 	if k == len(rules_[i].Subrules)-1 {
		// 		break
		// 	}
		// 	for _, v2 := range v {
		// 		subrule := "(" + compile(v2) + ")+"
		// 		subrules = append(subrules, subrule)
		// 	}
		// }
		// r = "(" + strings.Join(subrules, "") + ")"
		if i == 8 {
			r = "(" + compile(42) + "+" + ")"
		}
		if i == 11 {
			r42 := compile(42)
			r31 := compile(31)

			r2 := "(" + r42 + "" + r31 + ")"
			// r2 := ""
			for k := 2; k < 10; k++ {
				r2 = fmt.Sprintf("(%s|(%s{%d}%s{%d}))", r2, r42, k, r31, k)
			}
			r = r2
		}
	}

	rulesMap[i] = r
	return r
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[int]Rule, 0)
	// input := make([]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		rule := strings.Split(l, ":")
		r := rule[1]
		ruleNStr := rule[0]
		ruleN, _ := strconv.Atoi(ruleNStr)
		r = r[1:]

		myrule := Rule{}
		if r[0] == '"' {
			myrule.Literal = r[1 : len(r)-1]
		} else {
			subs := strings.Split(r, "|")
			r0 := make([][]int, 0)
			for _, sub := range subs {
				r2 := strings.Split(sub, " ")

				r3 := make([]int, 0)
				for _, r4 := range r2 {
					a, err := strconv.Atoi(r4)
					if err != nil {
						continue
					}
					r3 = append(r3, a)
				}
				r0 = append(r0, r3)
			}
			myrule.Subrules = r0
		}
		rules[ruleN] = myrule
		// rules = append(rules, myrule)
	}
	log.Printf("%+v\n", rules)
	rules_ = rules
	rule := "^" + compile(0) + "$"
	log.Println(rule)
	re := regexp.MustCompile(rule)

	// input := make([]string, 0)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		m := re.MatchString(l)
		log.Println(l, m)
		// input = append(input, l)
		if m {
			count++
		}
	}
	ret1 = count
	//
	// log.Println(input)

	return ret1, ret2
}

func do2(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[int]Rule, 0)
	// input := make([]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		rule := strings.Split(l, ":")
		r := rule[1]
		ruleNStr := rule[0]
		ruleN, _ := strconv.Atoi(ruleNStr)
		r = r[1:]

		myrule := Rule{}
		if r[0] == '"' {
			myrule.Literal = r[1 : len(r)-1]
		} else {
			subs := strings.Split(r, "|")
			r0 := make([][]int, 0)
			for _, sub := range subs {
				r2 := strings.Split(sub, " ")

				r3 := make([]int, 0)
				for _, r4 := range r2 {
					a, err := strconv.Atoi(r4)
					if err != nil {
						continue
					}
					r3 = append(r3, a)
				}
				r0 = append(r0, r3)
			}
			myrule.Subrules = r0
		}
		rules[ruleN] = myrule
		// rules = append(rules, myrule)
	}

	rules[8] = Rule{
		Subrules: [][]int{{42}, {42, 8}},
	}
	rules[11] = Rule{
		Subrules: [][]int{{42, 31}, {42, 11, 31}},
	}
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31

	log.Printf("%+v\n", rules)
	rules_ = rules
	rule := "^" + compile(0) + "$"
	log.Println(rule)
	re := regexp.MustCompile(rule)

	// input := make([]string, 0)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		m := re.MatchString(l)
		log.Println(l, m)
		// input = append(input, l)
		if m {
			count++
		}
	}
	ret1 = count
	//
	// log.Println(input)

	return ret1, ret2
}

func main() {
	// log.Println(do("test.txt"))
	// log.Println(do2("test2.txt"))
	// log.Println(do2("test2.txt"))
	// log.Println(do2("test3.txt"))
	// log.Println(do("test2.txt"))
	// log.Println(do2("test2.txt"))
	// log.Println(do("input.txt"))
	// log.Println(do2("input.txt"))
	// log.Println(do2("test3.txt"))

	// 371 too high
	// 266 too low
	// 273
	// 274

	// 357
	log.Println(do2("input.txt"))
	// log.Println(do("input.txt"))
}
