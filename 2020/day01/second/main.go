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

	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Printf("Solution: %d\n", numbers[i]*numbers[j]*numbers[k])
					return
				}
			}
		}
	}

	fmt.Println("Solution not found")
}
