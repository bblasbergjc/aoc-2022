package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func scenicScore(grid [][]int, row int, col int) int {
	treeHeight := grid[row][col]

	upScore := 0
	// up
	for y := row - 1; y >= 0; y -= 1 {
		if grid[y][col] >= treeHeight {
			upScore += 1
			break
		}

		upScore += 1
	}

	downScore := 0
	// down
	for y := row + 1; y < len(grid); y += 1 {
		if grid[y][col] >= treeHeight {
			downScore += 1
			break
		}

		downScore += 1
	}

	leftScore := 0
	//left
	for x := col - 1; x >= 0; x -= 1 {
		if grid[row][x] >= treeHeight {
			leftScore += 1
			break
		}

		leftScore += 1
	}

	rightScore := 0
	//right
	for x := col + 1; x < len(grid[0]); x += 1 {
		if grid[row][x] >= treeHeight {
			rightScore += 1
			break
		}

		rightScore += 1
	}

	return upScore * downScore * leftScore * rightScore
}

// todo: find a way to shorten this function
func isVisibleFromAnyDirection(grid [][]int, row int, col int) bool {
	treeHeight := grid[row][col]
	visible := true

	//visible up
	for y := row - 1; y >= 0; y -= 1 {
		if grid[y][col] >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	//visible down
	for y := row + 1; y < len(grid); y += 1 {
		if grid[y][col] >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	visible = true
	//visible left
	for x := col - 1; x >= 0; x -= 1 {
		if grid[row][x] >= treeHeight {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	//visible right
	for x := col + 1; x < len(grid[0]); x += 1 {
		if grid[row][x] >= treeHeight {
			visible = false
			break
		}
	}

	return visible
}

func partOne(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// init our number of visible trees with all trees on the outer edges
	// minus 4 at the end to account for the corners which are part of both rows and cols
	numVisible := (rows * 2) + (cols * 2) - 4

	// start at row 1 col 1 because we already calculated the outer rows and cols
	for row := 1; row < len(grid)-1; row += 1 {
		for col := 1; col < len(grid[0])-1; col += 1 {
			if isVisibleFromAnyDirection(grid, row, col) {
				numVisible += 1
			}
		}
	}
	return numVisible
}

func partTwo(grid [][]int) int {
	maxScore := 0

	// start at row 1 col 1 because outer edges will contain a zero and
	// with thus always have a score of zero
	for row := 1; row < len(grid)-1; row += 1 {
		for col := 1; col < len(grid[0])-1; col += 1 {
			score := scenicScore(grid, row, col)

			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func main() {
	data, err := os.ReadFile("./day8.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")

	rows := make([][]int, 0)

	// parse the grid of trees
	for _, line := range lines {
		if line == "" {
			break
		}
		cols := make([]int, 0)

		for i := range line {
			height, err := strconv.Atoi(string(line[i]))
			checkErr(err)

			cols = append(cols, height)
		}

		rows = append(rows, cols)
	}

	fmt.Println("Part 1: ", partOne(rows))
	fmt.Println("Part 2: ", partTwo(rows))
}
