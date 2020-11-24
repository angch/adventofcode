package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// mostCommon returns the top 5 most common rune, alphabetically
type kv struct {
	Key   rune
	Count int
}

func mostCommon(a string) []string {
	counts := make([]kv, 256)
	for _, v := range a {
		if v == '-' || v == '[' || v == ']' {
			continue
		}
		counts[int(v)].Key = v
		counts[int(v)].Count++
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count ||
			(counts[i].Count == counts[j].Count && counts[i].Key < counts[j].Key)
	})

	top5 := make([]string, 5)
	for i := range top5 {
		top5[i] = string(counts[i].Key)
	}
	return top5
}

type Component struct {
	Name     string
	SectorId int
	Checksum string
}

// splitToComponents decomposes "aaaaa-bbb-z-y-x-123[abxyz]"
// into Component{"aaaaa-bbb-z-y-x", 123, "abxyz"}
func splitToComponents(a string) Component {
	c := Component{}

	dashes := strings.Split(a, "-")
	if len(dashes) < 2 {
		log.Fatal(dashes)
	}
	rightMost := dashes[len(dashes)-1]
	c.Name = strings.Join(dashes[:len(dashes)-1], "-")
	secCheck := strings.Split(rightMost, "[")
	if len(secCheck) < 2 {
		log.Fatal(secCheck)
	}
	var err error
	c.SectorId, err = strconv.Atoi(secCheck[0])
	if err != nil {
		log.Fatal(err)
	}
	checksum := secCheck[1]
	c.Checksum = checksum[:len(checksum)-1]

	return c
}

func splitToComponents2(a string) Component {
	c := Component{}
	fragments := make([]string, 5)
	// Ok this didn't work. Revert to splitToComponents
	fmt.Sscanf(a, "%s-%s-%s-%s-%s-%d[%s]",
		&fragments[0], &fragments[1], &fragments[2], &fragments[3], &fragments[4],
		&c.SectorId, &c.Checksum,
	)
	c.Name = strings.Join(fragments, "-")
	return c
}

func isReal(c Component) bool {
	return c.Checksum == strings.Join(mostCommon(c.Name), "")
}

func decrypt(c Component) string {
	shift := c.SectorId % 26
	out := []byte(c.Name)
	for k, v := range c.Name {
		if v == '-' {
			out[k] = ' '
		} else {
			out[k] = byte(int(out[k]) + shift)
			if out[k] > 'z' {
				out[k] -= 26
			}
		}

	}
	return string(out)
}

func main() {
	inputs := []string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]",
		"qzmt-zixmtkozy-ivhz-343[zimth]",
	}
	for _, input := range inputs {
		c := splitToComponents(input)
		log.Printf("%+v\n", c)
		log.Println(mostCommon(c.Name))

		log.Printf("isReal: %v %s\n", isReal(c), decrypt(c))
		// log.Printf("%+v\n", splitToComponents2(input))
	}

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()
		c := splitToComponents(l)
		// log.Printf("%+v\n", c)
		// log.Println(mostCommon(c.Name), isReal(c))

		if isReal(c) {
			count += c.SectorId
			d := decrypt(c)
			if strings.HasPrefix(d, "northpole") {
				log.Println(d, c.SectorId)
			}
		}
	}
	log.Println(count) // too low
}
