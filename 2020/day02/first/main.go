package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Range    []int
	Letter   string
	Password string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	var inputList []Input

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		inputList = append(inputList, parseInput(line))

	}

	solution := countValidPasswords(inputList)
	fmt.Printf("total of valid passwords: %d\n", solution)
}

func parseInput(input string) Input {
	substrings := strings.Split(input, " ")

	minMax := strings.Split(substrings[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	letter := strings.Split(substrings[1], "")[0]
	return Input{[]int{min, max}, letter, substrings[2]}
}

func countValidPasswords(inputs []Input) int {
	var count int
	for _, input := range inputs {
		var charCount int
		for _, char := range input.Password {
			if string(char) == input.Letter {
				charCount++
			}
		}

		if charCount >= input.Range[0] && charCount <= input.Range[1] {
			count++
		}
	}

	return count
}
