package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	passports := parseInput("../input.txt")

	validPassports := 0
	for _, pass := range passports {
		if len(pass) == 8 {
			validPassports++
		} else if len(pass) == 7 && pass["cid"] == "" {
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
