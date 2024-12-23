package day1

import (
	"aoc2024/lib"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func SolutionB() {
	start := time.Now()
	listOne := []int{}
	listTwo := []int{}

	lib.ReadFileByLine("day1/input.txt", func(line string) {
		parts := strings.Fields(line)
		one, _ := strconv.Atoi(parts[0])
		two, _ := strconv.Atoi(parts[1])
		listOne = append(listOne, one)
		listTwo = append(listTwo, two)
	})

	sum := 0
	for _, num := range listOne {
		occurance := 0
		for _, numTwo := range listTwo {
			if num == numTwo {
				occurance++
			}
		}
		sum += occurance * num
		occurance = 0
	}

	elapsed := time.Since(start)
	fmt.Printf("Answer day 1b: %d in %s\n", sum, elapsed)
}
