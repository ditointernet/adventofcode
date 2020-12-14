package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("failed to open file")
	}

	var numbers []int

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				fmt.Printf("Solution: %d\n", numbers[i]*numbers[j])
				return
			}
		}
	}

	fmt.Println("Solution not found")
}
