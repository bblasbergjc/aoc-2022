package main

import (
	"fmt"
	"os"
	"strings"
)

const debug = false

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Point struct {
	Row int
	Col int
	Val rune
}

// builds the grid and returns the starting point and all 'a' points
func buildGrid(lines []string) ([][]rune, Point, []Point) {
	var start Point
	var aPoints []Point
	grid := make([][]rune, len(lines))
	for row, line := range lines {
		heights := make([]rune, len(line))
		for col, ch := range line {
			heights[col] = ch

			if ch == 'S' {
				start = Point{row, col, ch}
				aPoints = append(aPoints, start)
			} else if ch == 'a' {
				aPoints = append(aPoints, Point{row, col, ch})
			}
		}
		grid[row] = heights
	}

	return grid, start, aPoints
}

func getPoint(grid [][]rune, row int, col int) *Point {
	if row >= len(grid) || row < 0 {
		return nil
	}

	if col >= len(grid[0]) || col < 0 {
		return nil
	}
	return &Point{row, col, grid[row][col]}
}

func isValidMoveFunc(grid [][]rune, start Point, visited [][]bool) func(int, int) (Point, bool) {
	return func(row, col int) (Point, bool) {
		next := getPoint(grid, row, col)

		if next == nil {
			return Point{}, false
		}

		nextVal := next.Val
		if nextVal == 'E' {
			nextVal = 'z'
		}

		startVal := start.Val
		if startVal == 'S' {
			startVal = 'a'
		}

		valid := !visited[row][col] && nextVal-startVal <= 1

		return Point{row, col, grid[row][col]}, valid
	}
}

func validMoves(grid [][]rune, start Point, visited [][]bool) []Point {
	validMoves := make([]Point, 0)

	isValidMove := isValidMoveFunc(grid, start, visited)

	// up
	if next, ok := isValidMove(start.Row-1, start.Col); ok {
		validMoves = append(validMoves, next)
	}

	// down
	if next, ok := isValidMove(start.Row+1, start.Col); ok {
		validMoves = append(validMoves, next)
	}

	// left
	if next, ok := isValidMove(start.Row, start.Col-1); ok {
		validMoves = append(validMoves, next)
	}

	// right
	if next, ok := isValidMove(start.Row, start.Col+1); ok {
		validMoves = append(validMoves, next)
	}

	if debug {
		fmt.Println(start, "Valid moves:", len(validMoves))
	}

	return validMoves
}

func bfs(grid [][]rune, starts []Point) int {
	queue := make([]Point, 0)
	visited := make([][]bool, len(grid))
	for row := 0; row < len(grid); row += 1 {
		visited[row] = make([]bool, len(grid[0]))
		for col := 0; col < len(grid[0]); col += 1 {
			visited[row][col] = false
		}
	}

	queueConents := make(map[Point]bool)
	queue = append(queue, starts...)
	depth := 0
	for len(queue) > 0 {
		childNodesSize := len(queue)

		if debug {
			fmt.Println("======= Depth", depth, "========")
		}

		// process all children at this depth
		for childNodesSize > 0 {
			// pop first node off the queue
			node := queue[0]
			queue = queue[1:]

			if debug {
				fmt.Printf("%c", node.Val)
			}

			visited[node.Row][node.Col] = true

			if node.Val == 'E' {
				if debug {
					fmt.Println()
				}
				return depth
			}

			validMoves := validMoves(grid, node, visited)
			removedDups := make([]Point, 0)

			for _, move := range validMoves {
				if _, ok := queueConents[move]; !ok {
					removedDups = append(removedDups, move)
					queueConents[move] = true
				}
			}

			// add the child nodes we can move to
			queue = append(queue, removedDups...)

			childNodesSize -= 1
		}

		depth += 1
	}

	return depth
}

func partOne(grid [][]rune, start Point) int {
	return bfs(grid, []Point{start})
}

func partTwo(grid [][]rune, starts []Point) int {
	return bfs(grid, starts)
}

func main() {
	data, err := os.ReadFile("./day12.txt")
	checkErr(err)
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1] // trim empty last line

	grid, start, aPoints := buildGrid(lines)

	fmt.Println("Part 1:", partOne(grid, start))
	fmt.Println("Part 2:", partTwo(grid, aPoints))

}
