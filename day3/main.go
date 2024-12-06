package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := readInput("day3/input.txt")
	part1 := findPatternAndCalculate(file)
	part2 := findPatternAndCalculateWDoAndDont(file)

	fmt.Println("[Part 1] Total mul value: ", part1)
	fmt.Println("[Part 2] Total mul value: ", part2)

}

func readInput(file string) string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func findPatternAndCalculate(line string) int {
	outerPattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`

	outerRe := regexp.MustCompile(outerPattern)
	outerMatches := outerRe.FindAllString(line, -1)

	var total int
	// Process each outer match
	for _, match := range outerMatches {
		// fmt.Printf("Outer match: %s\n", match)
		total += calculateMulValue(match)
		// fmt.Printf("Mul outer: %v\n", outer)
	}

	return total
}

func findPatternAndCalculateWDoAndDont(line string) int {
	re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	operations := re.FindAllString(line, -1)

	var result int
	mode := "enabled"

	for _, operation := range operations {
		switch operation {
		case "do()":
			mode = "enabled"
		case "don't()":
			mode = "disabled"
		default:
			if mode == "enabled" {
				result += calculateMulValue(operation)
			}
		}
	}

	return result
}

func calculateMulValue(s string) int {
	s = s[4 : len(s)-1]
	values := strings.Split(s, ",")
	a := convertStrToInt(values[0])
	b := convertStrToInt(values[1])
	return a * b
}

func convertStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
