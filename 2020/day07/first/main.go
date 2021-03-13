package main

import (
	"fmt"
	"os"

	"github.com/ditointernet/advnetofcode/2020/day07/parser"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	rules := parser.Parse(file)

	graph := map[string][]string{}

	for _, rule := range rules {
		for k, _ := range rule.Contents {
			graph[k] = append(graph[k], rule.Bag)
		}
	}

	result := findBags(graph, "shiny gold")
	fmt.Printf("How many bag colors can eventually contain at least one shiny gold bag? R: %d\n", len(result))

}

func findBags(graph map[string][]string, bag string) map[string]bool {
	bags := map[string]bool{}

	for _, c := range graph[bag] {
		bags[c] = true
		for k, _ := range findBags(graph, c) {
			bags[k] = true
		}
	}
	return bags
}
