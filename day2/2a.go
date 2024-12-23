package day2

import (
	"aoc2024/lib"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func SolutionA() {
	start := time.Now()
	list := [][]int{}

	lib.ReadFileByLine("day2/input.txt", func(line string) {
		parts := strings.Fields(line)
		numbers := []int{}
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			numbers = append(numbers, number)
		}
		list = append(list, numbers)
	})

	sum := 0
	safeValues := []int{-3, -2, -1, 1, 2, 3}

	for _, numList := range list {
		direction := ""
		for i, num := range numList {
			if i == 0 {
				continue
			}
			value := numList[i-1] - num
			safe := slices.Contains(safeValues, value)
			if !safe {
				break
			}
			if i == 1 && value > 0 {
				direction = "desc"
			}
			if i == 1 && value < 0 {
				direction = "asc"
			}
			if direction == "desc" && value < 0 || direction == "asc" && value > 0 {
				direction = ""
				break
			}
			if len(numList)-1 == i {
				sum++
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Answer day 2a: %d in %s\n", sum, elapsed)
}
