package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bblasbergjc/aoc-2022/util"
)

var startingPoint = Point{500, 0}

type Point struct {
	x int
	y int
}

type Range struct {
	low  int
	high int
}

func (r Range) InRange(num int) bool {
	return num >= r.low && num <= r.high
}

type Cave struct {
	x map[int][]Range
	y map[int][]Range

	sand  map[int]map[int]struct{}
	maxY  int
	edges Range
}

func (c *Cave) Print() {
	for y := 0; y <= c.maxY+2; y += 1 {
		for x := c.edges.low - 1; x <= c.edges.high+1; x += 1 {
			if c.HasRock(Point{x, y}) {
				fmt.Print("#")
			} else if c.HasSand(Point{x, y}) {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (c *Cave) AddSand(p Point) {
	if y, ok := c.sand[p.x]; ok {
		y[p.y] = struct{}{}
	} else {
		y := make(map[int]struct{})
		y[p.y] = struct{}{}
		c.sand[p.x] = y
	}

	if p.x < c.edges.low {
		c.edges.low = p.x
	}

	if p.x > c.edges.high {
		c.edges.high = p.x
	}
}

func (c *Cave) HasSand(p Point) bool {
	if y, ok := c.sand[p.x]; ok {
		_, ok := y[p.y]
		return ok
	}

	return false
}

func (c *Cave) HasRock(p Point) bool {
	if p.y == c.maxY+2 {
		return true
	}

	if ranges, ok := c.x[p.x]; ok {
		for _, r := range ranges {
			if r.InRange(p.y) {
				return true
			}
		}
	}

	if ranges, ok := c.y[p.y]; ok {
		for _, r := range ranges {
			if r.InRange(p.x) {
				return true
			}
		}
	}

	return false
}

func (c *Cave) IsAvailable(p Point) bool {
	return !c.HasRock(p) && !c.HasSand(p)
}

func (c *Cave) IsInAbyss(p Point) bool {
	return p.y > c.maxY
}

func parsePoint(raw string) Point {
	coords := strings.Split(raw, ",")

	x, err := strconv.Atoi(coords[0])
	util.CheckErr(err)

	y, err := strconv.Atoi(coords[1])
	util.CheckErr(err)

	return Point{x, y}
}

func parseCave(lines []string) Cave {
	cave := Cave{
		make(map[int][]Range),
		make(map[int][]Range),
		make(map[int]map[int]struct{}),
		0,
		Range{math.MaxInt, 0},
	}

	maxY := 0
	for _, line := range lines {
		var prev *Point
		rawPoints := strings.Split(line, " -> ")

		for _, rawPoint := range rawPoints {
			p := parsePoint(rawPoint)
			if p.y > maxY {
				maxY = p.y
			}
			if p.x < cave.edges.low {
				cave.edges.low = p.x
			}
			if p.x > cave.edges.high {
				cave.edges.high = p.x
			}
			if prev != nil {
				if prev.x == p.x {
					low := p.y
					high := prev.y
					if low > high {
						low, high = high, low
					}

					cave.x[p.x] = append(cave.x[p.x], Range{low, high})
				} else if prev.y == p.y {
					low := p.x
					high := prev.x
					if low > high {
						low, high = high, low
					}

					cave.y[p.y] = append(cave.y[p.y], Range{low, high})
				}

			}

			prev = &p
		}
	}

	cave.maxY = maxY

	return cave
}

func nextMove(p Point, r Cave) (Point, bool) {
	//down
	next := Point{p.x, p.y + 1}
	if r.IsAvailable(next) {
		return next, true
	}

	//diagonal down left
	next = Point{p.x - 1, p.y + 1}
	if r.IsAvailable(next) {
		return next, true
	}

	//diagonal right
	next = Point{p.x + 1, p.y + 1}
	if r.IsAvailable(next) {
		return next, true
	}

	return Point{}, false
}

func partOne(cave Cave) int {
	sandUnits := 0

	fmt.Println("MaxY", cave.maxY, "minx", cave.edges.low, "maxx", cave.edges.high)
	for {
		sandUnits += 1
		prevPoint := startingPoint
		for {
			currentPoint, ok := nextMove(prevPoint, cave)

			// if this point is lower (higher y value) than the lowest rock, its in the abyss
			if cave.IsInAbyss(currentPoint) {
				return sandUnits - 1
			}

			// if we are out of moves, add the sand to the cave
			// and get the next piece of sand
			if !ok {
				cave.AddSand(prevPoint)
				prevPoint = startingPoint
				break
			}

			prevPoint = currentPoint
		}
	}
}

func partTwo(cave Cave) int {
	sandUnits := 0

	fmt.Println("MaxY", cave.maxY, "minx", cave.edges.low, "maxx", cave.edges.high)
	for {
		sandUnits += 1
		prevPoint := startingPoint
		for {
			currentPoint, ok := nextMove(prevPoint, cave)

			// if we are out of moves, add the sand to the cave
			// and get the next piece of sand
			if !ok {
				cave.AddSand(prevPoint)
				if prevPoint == startingPoint {
					return sandUnits
				}
				break
			}

			prevPoint = currentPoint
		}
	}
}

func main() {
	lines := util.ParseLinesWithoutEndNewLine("./day14.txt")
	cave := parseCave(lines)
	answer := partOne(cave)

	cave = parseCave(lines)
	answer2 := partTwo(cave)
	cave.Print()

	fmt.Println("part 1:", answer)
	fmt.Println("part 2:", answer2)

}
