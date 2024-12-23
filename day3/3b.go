package day3

import (
	"aoc2024/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SolutionB() {
	start := time.Now()
	r, _ := regexp.Compile("(mul\\([0-9]{1,3},[0-9]{1,3}\\))|(do\\(\\))|(don't\\(\\))")
	replacer := strings.NewReplacer("mul(", "", ")", "")
	list := [][]int{}

	lib.ReadFileByLine("day3/input.txt", func(line string) {
		matches := r.FindAllString(line, -1)
		include := true
		for _, match := range matches {
			if match == "do()" {
				include = true
				continue
			}
			if match == "don't()" {
				include = false
				continue
			}
			if !include {
				continue
			}
			newString := replacer.Replace(match)
			parts := strings.Split(newString, ",")
			partOne, _ := strconv.Atoi(parts[0])
			partTwo, _ := strconv.Atoi(parts[1])
			ints := []int{partOne, partTwo}
			list = append(list, ints)
		}
	})

	sum := 0

	for _, item := range list {
		total := item[0] * item[1]
		sum += total
	}
	elapsed := time.Since(start)
	fmt.Printf("Answer day 3b: %d in %s\n", sum, elapsed)
}