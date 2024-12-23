package day2

import (
	"aoc2024/lib"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func SolutionB() {
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

	for _, numList := range list {
		if possibleSafe(numList) {
			sum++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Answer day 2b: %d in %s\n", sum, elapsed)
}

func possibleSafe(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	for i := range nums {
		new := make([]int, len(nums))
		copy(new, nums)

		if isSafe(append(new[:i], new[i+1:]...)) {
			return true
		}
	}

	return false
}

func isSafe(nums []int) bool {
	direction := ""
	safeValues := []int{-3, -2, -1, 1, 2, 3}
	for i, num := range nums {
		if i == 0 {
			continue
		}
		value := nums[i-1] - num
		safe := slices.Contains(safeValues, value)
		if !safe {
			return false
		}
		if i == 1 && value > 0 {
			direction = "desc"
		}
		if i == 1 && value < 0 {
			direction = "asc"
		}
		if direction == "desc" && value < 0 || direction == "asc" && value > 0 {
			return false
		}
	}
	return true
}
