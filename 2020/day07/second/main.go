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

	graph := map[string]map[string]int{}

	for _, rule := range rules {
		graph[rule.Bag] = rule.Contents
	}

	result := countBags(graph, "shiny gold")
	fmt.Printf("How many individual bags are required inside your single shiny gold bag? R: %d\n", result)

}

func countBags(graph map[string]map[string]int, bag string) int {
	var total int

	for k, v := range graph[bag] {
		total = total + v + v*countBags(graph, k)
	}

	return total
}
