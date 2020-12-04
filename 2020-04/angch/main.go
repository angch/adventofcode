package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var eclMap = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func checkvalid(kv map[string]string) bool {
	hclre := regexp.MustCompile("#[0-9a-fA-F]{6}")
	pidre := regexp.MustCompile("^[0-9]{9}$")
	for k, v := range kv {
		valid := false
		switch k {
		case "byr":
			if len(v) == 4 && v >= "1920" && v <= "2002" {
				valid = true
			}
		case "iyr":
			if len(v) == 4 && v >= "2010" && v <= "2020" {
				valid = true
			}
		case "eyr":
			if len(v) == 4 && v >= "2020" && v <= "2030" {
				valid = true
			}
		case "hgt":
			if strings.HasSuffix(v, "cm") {
				i := 0
				fmt.Sscanf(v, "%dcm", &i)
				if i >= 150 && i <= 193 {
					valid = true
				}
			} else if strings.HasSuffix(v, "in") {
				i := 0
				fmt.Sscanf(v, "%din", &i)
				if i >= 59 && i <= 76 {
					valid = true
				}
			}
		case "hcl":
			if hclre.MatchString(v) {
				valid = true
			}
		case "ecl":
			if eclMap[v] {
				valid = true
			}
		case "pid":
			if pidre.MatchString(v) {
				valid = true
			}
		case "cid":
			valid = true

		}
		if !valid {
			return false
		}
	}
	return true
}

func main() {
	inputs2 := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
		"",
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
		"",
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
	}
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}
	if false {
		log.Println(fields)
		log.Println(inputs2)
	}

	inputs := make([]string, 0)
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		inputs = append(inputs, l)
	}
	if false {
		log.Println(inputs)
	}
	current := make(map[string]string)
	valid := 0
	// not 54
	for _, v := range inputs {
		if v == "" {
			isvalid := checkvalid(current)
			if isvalid && len(current) == 8 {
				valid++
			} else if isvalid && len(current) == 7 {
				_, ok := current["cid"]
				if !ok {
					valid++
				}

			}
			log.Println(current, isvalid)
			current = make(map[string]string)
		} else {
			pairs := strings.Split(v, " ")
			for _, v2 := range pairs {
				kv := strings.Split(v2, ":")
				current[kv[0]] = kv[1]
			}
		}
	}
	isvalid := checkvalid(current)
	if isvalid && len(current) == 8 {
		valid++
	} else if isvalid && len(current) == 7 {
		_, ok := current["cid"]
		if !ok {
			valid++
		}

	}
	log.Println(current, isvalid)
	// not 224
	log.Println(valid)

}
