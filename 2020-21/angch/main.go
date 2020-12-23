package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gnboorse/centipede"
)

type Rule struct {
	Ingredient []string
	Allergen   []string
}

func intersect(a, b []string) []string {
	log.Println("intersect", a, b)
	a_ := make(map[string]bool)
	c := make([]string, 0)
	for _, v := range a {
		a_[v] = true
	}
	for _, v := range b {
		if a_[v] {
			c = append(c, v)
		}
	}
	sort.Strings(c)
	return c
}

func inslice(a string, b []string) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func do(fileName string) (ret1 int, ret2 int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	possibleAllergen := make(map[string]map[string]bool)
	allergenInIngredient := make(map[string]map[string]bool)
	appearances := make(map[string]int)
	rules := make([]Rule, 0)

	allergen := make([]string, 0)
	ingredient := make([]string, 0)
	allergenId := make(map[string]int)
	ingredientId := make(map[string]int)

	for scanner.Scan() {
		l := scanner.Text()
		_ = l

		if l == "" {
			break
		}
		rule := strings.Split(l, "(")
		contains := strings.Split(rule[1], ",")
		for k, v := range contains {
			contains[k] = strings.TrimPrefix(v, "contains ")
			contains[k] = strings.TrimSuffix(contains[k], ")")
			contains[k] = strings.TrimSpace(contains[k])
		}

		ingredients := strings.Split(strings.TrimSpace(rule[0]), " ")

		log.Println(ingredients, contains)

		for _, v := range ingredients {
			e, ok := possibleAllergen[v]
			if !ok {
				possibleAllergen[v] = make(map[string]bool)
				e = possibleAllergen[v]

				if _, ok = ingredientId[v]; !ok {
					ingredientId[v] = len(ingredient)
					ingredient = append(ingredient, v)
				}
			}

			for _, v2 := range contains {
				e[v2] = true

				f, ok := allergenInIngredient[v2]
				if !ok {
					allergenInIngredient[v2] = make(map[string]bool)
					f = allergenInIngredient[v2]
				}
				f[v] = true

				if _, ok = allergenId[v2]; !ok {
					allergenId[v2] = len(allergen)
					allergen = append(allergen, v2)
				}
			}

			// log.Println("incr", v)
			appearances[v]++
		}
		myrule := Rule{
			Ingredient: ingredients,
			Allergen:   contains,
		}
		rules = append(rules, myrule)
	}
	// log.Println("possibleAllergen", possibleAllergen)
	// log.Println("allergenInIngredient", allergenInIngredient)
	// log.Println(rules)
	// log.Println("appearances", appearances)

	{
		allergen := make(map[string][]string)
		counts := make(map[string]int)
		for k, v := range rules {
			log.Println(k, v)
			for _, all := range v.Allergen {
				if _, ok := allergen[all]; !ok {
					allergen[all] = make([]string, len(v.Ingredient))
					copy(allergen[all], v.Ingredient)
				} else {
					allergen[all] = intersect(allergen[all], v.Ingredient)
				}
			}
			for _, v := range v.Ingredient {
				counts[v]++
			}
		}

		ret1 = 0
	a:
		for k, v := range counts {
			for _, v2 := range allergen {
				if inslice(k, v2) {
					continue a
				}
			}
			ret1 += v
		}
		log.Println("out", allergen, ret1)
		confirmed := make(map[string]string)
	b:
		for k, v := range allergen {
			if len(v) == 1 {
				confirmed[k] = v[0]
				delete(allergen, k)

				for k2, v2 := range allergen {
					for k3, v3 := range v2 {
						if v3 == v[0] {
							// log.Println("xfound", v3, k3)
							allergen[k2] = append(v2[:k3], v2[k3+1:]...)
							break
						}
					}
				}
				// log.Println("removed", v[0], "become", allergen)
				goto b
			}
		}

		keys := make([]string, 0)
		for k, _ := range confirmed {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		outdanger := make([]string, 0)
		for _, v := range keys {
			outdanger = append(outdanger, confirmed[v])
		}
		log.Println("out2", confirmed, strings.Join(outdanger, ","))
	}

	if false {
		doesnotContain := make(map[string]map[string]bool)
		for rulen, rule := range rules {
			al := make(map[string]bool)
			notal := make([]string, 0)
			for _, al2 := range rule.Allergen {
				al[al2] = true
			}
			for al2 := range allergenInIngredient {
				if !al[al2] {
					notal = append(notal, al2)
				}
			}

			for _, ingredient := range rule.Ingredient {
				log.Println("Rule #", rulen+1, ingredient, "does not contain", notal)

				c, ok := doesnotContain[ingredient]
				if !ok {
					doesnotContain[ingredient] = make(map[string]bool)
					c = doesnotContain[ingredient]
				}
				for _, al2 := range notal {
					c[al2] = true
				}
			}
		}
		log.Println(doesnotContain)
	}
	log.Println("allergenid", allergenId)

	if false {
		// Ah fuggit.
		vars := make(centipede.Variables, 0)
		for i := 0; i < len(ingredient); i++ {
			varName := centipede.VariableName(ingredient[i])
			cvar := centipede.NewVariable(varName, centipede.IntRange(1, len(allergen)+1))
			vars = append(vars, cvar)
		}

		constraints := make([]centipede.Constraint, 0)
		allergenIds := make([][]int, 0)
		for ruleN, rule := range rules {
			varnames := make([]centipede.VariableName, 0)
			vv := make(map[string]bool)
			for _, id := range rule.Ingredient {
				varnames = append(varnames, centipede.VariableName(id))
				vv[id] = true
			}

			allergenIds = append(allergenIds, make([]int, 0))
			for _, id := range rule.Allergen {
				allergenIds[ruleN] = append(allergenIds[ruleN], allergenId[id])
			}

			ConstraintFunction := foo(ruleN, allergenIds, vv)
			log.Println("rule c added", ruleN, allergenIds, varnames)

			constraints = append(constraints, centipede.Constraint{Vars: varnames, ConstraintFunction: ConstraintFunction})
		}
		log.Println(vars, constraints)
		//constraints = append(constraints, centipede.Constraint{Vars: varnames, ConstraintFunction: ConstraintFunction})

		solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
		begin := time.Now()
		success := solver.Solve() // run the solution
		elapsed := time.Since(begin)
		if success {
			for _, variable := range solver.State.Vars {
				fmt.Printf("Variable %v = %v\n", variable.Name, variable.Value)
			}
		} else {
			log.Println("not found")
		}
		log.Println(elapsed)

		// 	centipede.NewVariable("A", centipede.IntRange(1, 10)),
		// 	centipede.NewVariable("B", centipede.IntRange(1, 10)),
		// 	centipede.NewVariable("C", centipede.IntRange(1, 10)),
		// 	centipede.NewVariable("D", centipede.IntRange(1, 10)),
		// 	centipede.NewVariable("E", centipede.IntRangeStep(0, 20, 2)), // even numbers < 20
		// }
	}
	// count := 0
	// for k, v := range doesnotContain {
	// 	if len(v) == len(allergenInIngredient) {
	// 		log.Println(k, "is ok", appearances[k])
	// 		count += appearances[k]
	// 	}
	// }
	// ret1 = count

	return ret1, ret2
}

func foo(n int, allergenIds [][]int, vv map[string]bool) func(variables *centipede.Variables) bool {
	return func(variables *centipede.Variables) bool {
		found := 0
		nn := make([]string, 0)
		for _, v := range *variables {
			nn = append(nn, string(v.Name))
		}
		log.Println("Rule ", n, nn)

		for _, id := range allergenIds[n] {
			for _, v := range *variables {
				if !vv[string(v.Name)] {
					continue
				}
				if !v.Empty {
					log.Println("xx rule", n, v.Name, v.Value, id+1)
					i, ok := v.Value.(int)
					if ok && i == id+1 {
						found++
						// log.Println("found")
						break
					}
				}
			}
		}
		log.Println("found", found, len(allergenIds[n]), found == len(allergenIds[n]))
		// return found > 0
		return found == len(allergenIds[n])
	}
}

func main() {
	// log.Println(do("test.txt"))
	log.Println(do("input.txt"))
}
