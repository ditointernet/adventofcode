package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BoardingPass struct {
	FrontBack, RightLeft []string
}

var Halves = map[string]int{
	"F": 1,
	"B": 0,
	"L": 1,
	"R": 0,
}

func main() {
	boardingPasses := parseInput("input.txt")

	var maxSeatID int

	for _, bp := range boardingPasses {
		seatID := binarySpacePartitioning(bp.FrontBack, 128)*8 + binarySpacePartitioning(bp.RightLeft, 8)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Printf("Max SeatID: %d\n", maxSeatID)
}

func binarySpacePartitioning(input []string, size int) int {
	bounds := []int{0, size - 1}
	for _, c := range input {
		diff := (bounds[1] - bounds[0] + 1) / 2
		op := Halves[c]

		if op == 1 {
			diff = diff * -1
		}

		bounds[op] = bounds[op] + diff
	}

	if bounds[0] != bounds[1] {
		panic("failed to calculte bounds")
	}

	return bounds[0]
}

func parseInput(filename string) []BoardingPass {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed to open file")
	}

	var boardingPasses []BoardingPass
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")

		pass := BoardingPass{}
		for i := 0; i < len(chars); i++ {
			if i < 7 {
				pass.FrontBack = append(pass.FrontBack, chars[i])
			} else {
				pass.RightLeft = append(pass.RightLeft, chars[i])
			}
		}

		boardingPasses = append(boardingPasses, pass)

	}

	return boardingPasses
}
