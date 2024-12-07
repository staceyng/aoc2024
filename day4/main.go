package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := readInput("day4/input.txt")
	// fmt.Println("File Contents: ", file)
	lines := strings.Split(file, "\n")
	// fmt.Println("line split:", lines)

	grid := make([]string, len(lines))

	for i, line := range lines {
		grid[i] = line
	}
	// fmt.Println("grid:", grid)

	// part 1
	ans := findXmas(grid, "XMAS")
	fmt.Println("[Part 1] ans:", len(ans))

}

func readInput(file string) string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func findXmas(grid []string, word string) []string {
	directions := []struct {
		dx, dy int
	}{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	foundLocations := []string{}
	wordLen := len(word)
	gridRows := len(grid)
	gridCols := len(grid[0])

	for row := 0; row < gridRows; row++ {
		for col := 0; col < gridCols; col++ {
			for _, dir := range directions {
				matched := true
				for k := 0; k < wordLen; k++ {
					nr := row + dir.dx*k
					nc := col + dir.dy*k
					if nr < 0 || nr >= gridRows || nc < 0 || nc >= gridCols || grid[nr][nc] != word[k] {
						matched = false
						break
					}
				}
				if matched {
					foundLocations = append(foundLocations, fmt.Sprintf("Start: (%d, %d), Direction: (%d, %d)", row, col, dir.dx, dir.dy))
				}
			}
		}
	}
	return foundLocations
}
