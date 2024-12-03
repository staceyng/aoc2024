package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList []int
	var rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("[DEBUG]: %s \n", line)
		split := strings.Split(line, "   ")
		// fmt.Printf("[DEBUG]: %v \n", split)
		leftList = append(leftList, convertStrToInt(split[0]))
		rightList = append(rightList, convertStrToInt(split[1]))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Printf("[DEBUG]: leftlist %v \n", leftList)
	// fmt.Printf("[DEBUG]: rightlist %v \n", rightList)
	copyLeftList := deepCopySlice(leftList)
	copyRightList := deepCopySlice(rightList)

	newLeft, newRight := pairList(leftList, rightList)
	// fmt.Printf("[DEBUG]: newleftlist: %v \n", newLeft)
	// fmt.Printf("[DEBUG]: newrightlist: %v \n", newRight)

	distance := getTotalDistance(newLeft, newRight)
	fmt.Println("Total Distance:", distance)
	similarity := getSimilarityScore(copyLeftList, copyRightList)
	fmt.Println("Similarity Score:", similarity)
}

func getSimilarityScore(left []int, right []int) int {
	// part2
	count := len(left)
	var similarity int

	for i := 0; i < count; i++ {
		val := left[i]
		v := val * countOccurrences(right, val)
		similarity += int(v)
	}

	return similarity
}

func getTotalDistance(left []int, right []int) int {
	// part1
	count := len(left)
	var distance int

	for i := 0; i < count; i++ {
		v := math.Abs(float64(left[i] - right[i]))
		distance += int(v)
	}

	return distance
}

func pairList(left []int, right []int) ([]int, []int) {
	var newLeft []int
	var newRight []int
	count := len(left)
	fmt.Printf("[DEBUG]: count=%v \n", count)

	for i := 0; i < count; i++ {
		minLeft := slices.Min(left)
		minRight := slices.Min(right)
		// fmt.Printf("[DEBUG]: minleft=%v, minright=%v \n", minLeft, minRight)
		newLeft = append(newLeft, minLeft)
		newRight = append(newRight, minRight)
		left = removeFirstOccurenceFromSlice(left, minLeft)
		right = removeFirstOccurenceFromSlice(right, minRight)
	}

	return newLeft, newRight
}

func countOccurrences(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func removeFirstOccurenceFromSlice(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func convertStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

// deep copy slice for primitive types
func deepCopySlice(original []int) []int {
	copied := make([]int, len(original))
	copy(copied, original)
	return copied
}
