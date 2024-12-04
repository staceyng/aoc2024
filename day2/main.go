package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("[DEBUG]: %s \n", line)
		split := strings.Split(line, " ")
		// fmt.Printf("[DEBUG]: split=%v \n", split)
		data = append(data, convertStrSliceToIntSlice(split))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Printf("[DEBUG] data=%v", data)

	// part1
	var safeReports int
	for _, v := range data {
		if checkIsSafe(v) {
			safeReports += 1
		}
	}

	fmt.Printf("Number of Safe Reports: %d\n", safeReports)

	// part2
	var trueSafe int
	var modSafe int
	for _, v := range data {
		if checkIsSafe(v) {
			trueSafe += 1
		} else {
			for i := 0; i < len(v); i++ {
				var modlevels []int
				modlevels = append(modlevels, v[:i]...)
				modlevels = append(modlevels, v[i+1:]...)
				isModSafe := checkIsSafe(modlevels)

				if isModSafe {
					modSafe += 1
					break
				}
			}

		}
	}
	fmt.Printf("Number of TrueSafe Reports: %d\n", trueSafe)
	fmt.Printf("Number of modSafe Reports: %d\n", modSafe)
	fmt.Printf("Number of Safe+ModSafe Reports: %d\n", trueSafe+modSafe)

}

func convertStrSliceToIntSlice(s []string) []int {
	slice := make([]int, len(s))

	for i, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		slice[i] = num
	}
	return slice
}

func checkIsSafe(slice []int) bool {
	// slice is safe only if
	// 1. all increasing or all decreasing
	// 2. difference between 2 elements are at least 1 and at most 3
	if len(slice) <= 1 {
		return true
	}

	isIncreasing := slice[1] > slice[0]
	// use 2 pointers to determine if inc/desc and diff
	lp := 0
	rp := 1

	for rp < len(slice) {
		left := slice[lp]
		right := slice[rp]

		if (right > left && !isIncreasing) || (right < left && isIncreasing) {
			return false
		}

		diff := int(math.Abs(float64(right - left)))
		if diff < 1 || diff > 3 {
			return false
		}
		lp++
		rp++
	}
	return true
}
