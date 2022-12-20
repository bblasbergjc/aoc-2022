package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

type Point struct {
	X int
	Y int
	Z int
}

type Coords struct {
	m map[int]map[int]map[int]struct{}
}

func (c *Coords) Add(x, y, z int) {
	_, hasX := c.m[x]
	if !hasX {
		c.m[x] = make(map[int]map[int]struct{})
	}

	_, hasY := c.m[x][y]
	if !hasY {
		c.m[x][y] = make(map[int]struct{})
	}

	_, hasZ := c.m[x][y][z]
	if !hasZ {
		c.m[x][y][z] = struct{}{}
	}
}

func (c *Coords) Contains(x, y, z int) bool {
	_, hasX := c.m[x]
	if !hasX {
		return false
	}

	_, hasY := c.m[x][y]
	if !hasY {
		return false
	}

	_, hasZ := c.m[x][y][z]
	if !hasZ {
		return false
	}

	return true
}

func parseInput(lines []string) Coords {
	coords := Coords{
		m: make(map[int]map[int]map[int]struct{}),
	}
	for _, line := range lines {
		rawNums := strings.Split(line, ",")

		x, err := strconv.Atoi(rawNums[0])
		CheckErr(err)

		y, err := strconv.Atoi(rawNums[1])
		CheckErr(err)

		z, err := strconv.Atoi(rawNums[2])
		CheckErr(err)

		coords.Add(x, y, z)
	}
	return coords
}

func Contains(lst []int, x int) bool {
	for i := range lst {
		if lst[i] == x {
			return true
		}
	}
	return false
}

func (c *Coords) CountOpenSides(x int, y int, z int) int {
	openSides := 0

	// right
	if !c.Contains(x+1, y, z) {
		openSides += 1
	}

	// left
	if !c.Contains(x-1, y, z) {
		openSides += 1
	}

	// above
	if !c.Contains(x, y+1, z) {
		openSides += 1
	}

	// below
	if !c.Contains(x, y-1, z) {
		openSides += 1
	}

	// in front
	if !c.Contains(x, y, z-1) {
		openSides += 1
	}

	// behind
	if !c.Contains(x, y, z+1) {
		openSides += 1
	}

	return openSides
}

func partOne(coords Coords) int {
	openSides := 0
	for x := range coords.m {
		for y := range coords.m[x] {
			for z := range coords.m[x][y] {
				openSides += coords.CountOpenSides(x, y, z)
			}
		}
	}
	return openSides
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./day18.txt")
	coords := parseInput(lines)
	fmt.Println("Part 1:", partOne(coords))
}
