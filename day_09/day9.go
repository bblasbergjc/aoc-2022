package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

const debug = false

type Position struct {
	X int
	Y int
}

func newPosition() Position {
	return Position{0, 0}
}

func parseLine(line string) (string, int) {
	parts := strings.Split(line, " ")
	moveNum, err := strconv.Atoi(parts[1])
	CheckErr(err)

	return parts[0], moveNum
}

func moveByOne(pos Position, direction string) Position {
	if direction == "R" {
		pos.X += 1
	} else if direction == "U" {
		pos.Y += 1
	} else if direction == "L" {
		pos.X -= 1
	} else {
		pos.Y -= 1
	}

	return pos
}

// is there a way to do this more consicely?
func moveTailKnot(head, tail Position) Position {
	if head.X-tail.X == 2 { //head right of tail
		tail.X += 1

		if head.Y-tail.Y > 0 { // tail is diagonally below
			tail.Y += 1
		} else if tail.Y-head.Y > 0 { // tail is diagonally above
			tail.Y -= 1
		}
	} else if tail.X-head.X == 2 { //head left of tail
		tail.X -= 1

		if head.Y-tail.Y > 0 { // tail is diagonally below
			tail.Y += 1
		} else if tail.Y-head.Y > 0 { // tail is diagonally above
			tail.Y -= 1
		}
	} else if head.Y-tail.Y == 2 { //head above tail
		tail.Y += 1

		if head.X-tail.X > 0 { // tail is diagonally left
			tail.X += 1
		} else if tail.X-head.X > 0 { // tail is diagonally right
			tail.X -= 1
		}
	} else if tail.Y-head.Y == 2 { //head below tail
		tail.Y -= 1

		if head.X-tail.X > 0 { // tail is diagonally left
			tail.X += 1
		} else if tail.X-head.X > 0 { // tail is diagonally right
			tail.X -= 1
		}
	}

	return tail
}

// How many positions does the tail of the rope visit at least once?
func partOne(lines []string) int {
	head := newPosition()
	tail := newPosition()

	// set of coordinates the tail has visited
	positionsVisited := make(map[int]map[int]struct{})

	uniqueTailPositions := 0 // don't count the starting position
	for _, line := range lines {
		if line == "" {
			break
		}

		direction, movement := parseLine(line)

		for i := 0; i < movement; i += 1 {
			head = moveByOne(head, direction)

			if debug {
				fmt.Println("-----------")
				fmt.Println("head X:", head.X, "Y:", head.Y)
				fmt.Println("tail X:", tail.X, "Y:", tail.Y)
			}

			tail = moveTailKnot(head, tail)

			_, hasX := positionsVisited[tail.X]
			if !hasX {
				positionsVisited[tail.X] = make(map[int]struct{})
			}

			_, hasY := positionsVisited[tail.X][tail.Y]
			if !hasY {
				uniqueTailPositions += 1
			}

			positionsVisited[tail.X][tail.Y] = struct{}{}
		}
	}

	return uniqueTailPositions
}

// how many unique positions does the tail visit if we have 9 knots?
func partTwo(lines []string) int {
	head := newPosition()

	knots := [9]Position{}
	for i := 0; i < len(knots); i += 1 {
		knots[i] = newPosition()
	}

	// set of coordinates the tail has visited
	positionsVisited := make(map[int]map[int]struct{})

	uniqueTailPositions := 0 // don't count the starting position
	for _, line := range lines {
		if line == "" {
			break
		}

		direction, movement := parseLine(line)

		for i := 0; i < movement; i += 1 {
			head = moveByOne(head, direction)

			// move all our tail knots using the preceding knot as the "head"
			knots[0] = moveTailKnot(head, knots[0])
			for i := 1; i < len(knots); i += 1 {
				knots[i] = moveTailKnot(knots[i-1], knots[i])
			}

			tail := knots[len(knots)-1]
			_, hasX := positionsVisited[tail.X]
			if !hasX {
				positionsVisited[tail.X] = make(map[int]struct{})
			}

			_, hasY := positionsVisited[tail.X][tail.Y]
			if !hasY {
				uniqueTailPositions += 1
			}

			positionsVisited[tail.X][tail.Y] = struct{}{}
		}
	}

	return uniqueTailPositions

}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day9.txt")

	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))
}
