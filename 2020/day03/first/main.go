package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := parseInput("input.txt")
	var trees int

	c := 3
	for l := 1; l < len(matrix); l++ {
		item := matrix[l][c]

		if item == "#" {
			trees++
		}

		c = (c + 3) % len(matrix[l])
	}

	fmt.Printf("Number of trees: %d\n", trees)
}

func parseInput(filename string) [][]string {
	var matrix [][]string

	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)

	}

	return matrix
}
