package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	var groups []map[string]int
	var group = make(map[string]int)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()

		if line == "" {
			groups = append(groups, group)

			group = make(map[string]int)
			continue
		}

		questions := strings.Split(line, "")
		for _, c := range questions {
			group[c]++
		}

	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	var total int
	for _, g := range groups {
		fmt.Printf("%+v => %d\n", g, len(g))
		total = total + len(g)
	}

	fmt.Printf("sum of counts: %d\n", total)
}
