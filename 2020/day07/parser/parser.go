package parser

import (
	"bufio"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Bag      string
	Contents map[string]int
}

func (r1 Rule) Equal(r2 Rule) bool {
	if r1.Bag != r2.Bag {
		return false
	}

	return reflect.DeepEqual(r1.Contents, r2.Contents)
}

func Parse(input io.Reader) []Rule {
	var rules []Rule
	s := bufio.NewScanner(input)

	re := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

	for s.Scan() {
		parts := strings.Split(s.Text(), "bags contain")
		bag := strings.TrimSpace(parts[0])

		if parts[1] == " no other bags." {
			rules = append(rules, Rule{bag, map[string]int{}})
			continue
		}

		matches := re.FindAllStringSubmatch(parts[1], -1)
		contents := map[string]int{}

		for _, m := range matches {
			count, _ := strconv.Atoi(m[1])
			contents[m[2]] = count
		}

		rules = append(rules, Rule{bag, contents})

	}
	return rules
}
