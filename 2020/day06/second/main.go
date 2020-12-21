package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Group struct {
	Answers         map[string]int
	NumOfPassengers int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	var groups []Group
	var group = Group{make(map[string]int), 0}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()

		if line == "" {
			groups = append(groups, group)

			group = Group{make(map[string]int), 0}
			continue
		}

		group.NumOfPassengers++
		questions := strings.Split(line, "")
		for _, c := range questions {
			group.Answers[c]++
		}

	}

	if len(group.Answers) > 0 {
		groups = append(groups, group)
	}

	var total int
	for _, g := range groups {
		for _, v := range g.Answers {
			if v >= g.NumOfPassengers {
				total++
			}
		}
	}

	fmt.Printf("sum of counts: %d\n", total)
}
