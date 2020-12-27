package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := parseInput("../input.txt")

	validPassports := 0
	for _, pass := range passports {
		if isValidPassport(pass) {
			validPassports++
		}
	}
	fmt.Printf("Valids passports: %d\n", validPassports)
}

func parseInput(filename string) []map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}

	var fields []string
	var passports []map[string]string

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		if line != "" {
			fields = append(fields, strings.Split(line, " ")...)
		} else {
			pass := make(map[string]string)
			for _, f := range fields {
				keyValueField := strings.Split(f, ":")
				pass[keyValueField[0]] = keyValueField[1]
			}
			passports = append(passports, pass)
			fields = nil
		}
	}

	pass := make(map[string]string)
	for _, f := range fields {
		keyValueField := strings.Split(f, ":")
		pass[keyValueField[0]] = keyValueField[1]
	}
	passports = append(passports, pass)

	return passports
}

func isValidPassport(pass map[string]string) bool {
	if len(pass) < 7 {
		return false
	}

	if pass["byr"] == "" {
		return false
	}

	byr, _ := strconv.Atoi(pass["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	if pass["iyr"] == "" {
		return false
	}

	iyr, _ := strconv.Atoi(pass["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	if pass["eyr"] == "" {
		return false
	}

	eyr, _ := strconv.Atoi(pass["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	if pass["hgt"] == "" {
		return false
	}

	hgt, _ := strconv.Atoi(pass["hgt"][:len(pass["hgt"])-2])
	if strings.Contains(pass["hgt"], "cm") {
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.Contains(pass["hgt"], "in") {
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	if pass["hcl"] == "" {
		return false
	}

	hclRegex := regexp.MustCompile("([0-9]|[a-f])")
	hclString := hclRegex.FindAllString(pass["hcl"][1:], -1)
	if pass["hcl"][:1] != "#" {
		return false
	}

	if len(hclString) != len(pass["hcl"][1:]) {
		return false
	}

	if pass["ecl"] == "" {
		return false
	}

	eclRegex := regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")
	if !eclRegex.MatchString(pass["ecl"]) {
		return false
	}

	if pass["pid"] == "" {
		return false
	}

	if len(pass["pid"]) != 9 {
		return false
	}

	return true

}
