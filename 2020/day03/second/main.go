package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := parseInput("input.txt")

	trees := 1
	trees = trees * walkAndCount(matrix, 1, 1)
	trees = trees * walkAndCount(matrix, 3, 1)
	trees = trees * walkAndCount(matrix, 5, 1)
	trees = trees * walkAndCount(matrix, 7, 1)
	trees = trees * walkAndCount(matrix, 1, 2)

	fmt.Printf("Number of trees: %d\n", trees)
}

func walkAndCount(matrix [][]string, right, down int) int {
	var trees int
	c := right
	for l := down; l < len(matrix); l = l + down {
		item := matrix[l][c]

		if item == "#" {
			trees++
		}

		c = (c + right) % len(matrix[l])
	}

	return trees
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
