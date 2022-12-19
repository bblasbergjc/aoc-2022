package main

import (
	"fmt"
	"math"

	. "github.com/bblasbergjc/aoc-2022/util"
)

const (
	left = iota
	right
	down
)

type movement int

type Point struct {
	X int
	Y int
}

type Shape struct {
	Rocks     map[int][]int //y, [x]
	Leftmost  int
	Rightmost int
	Lowest    int
	Highest   int
}

func Contains(lst []int, x int) bool {
	for _, item := range lst {
		if item == x {
			return true
		}
	}

	return false
}

func (s *Shape) HasRock(y, x int) bool {
	if level, hasY := s.Rocks[y]; hasY {
		return Contains(level, x)
	}

	return false
}

func (s *Shape) Clone() Shape {
	position := make(map[int][]int)
	for y, x := range s.Rocks {
		var xs []int
		for i := range x {
			xs = append(xs, x[i])
		}
		position[y] = xs
	}

	return Shape{
		Rocks:     position,
		Leftmost:  s.Leftmost,
		Rightmost: s.Rightmost,
		Lowest:    s.Lowest,
		Highest:   s.Highest,
	}
}

func (s *Shape) AddToMap(allRocks map[int][]int) {
	for y, xs := range s.Rocks {
		_, hasY := allRocks[y]
		if hasY {
			allRocks[y] = append(allRocks[y], xs...)
		} else {
			newXs := make([]int, 0)
			newXs = append(newXs, xs...)
			allRocks[y] = newXs
		}
	}
}

func (s *Shape) CanMove(dy int, dx int, allRocks map[int][]int) bool {
	for y, xs := range s.Rocks {
		for _, x := range xs {
			newX := x + dx
			newY := y + dy

			if newX < 0 || newX > 6 || newY < 0 {
				return false
			}

			if cols, hasY := allRocks[newY]; hasY {
				if Contains(cols, newX) {
					return false
				}
			}
		}
	}

	return true
}

func (s *Shape) SetStartingPosition(height int) {
	s.UpdatePosition(0, 0)
	s.UpdatePosition(height, s.Leftmost+2)
}

func (s *Shape) UpdatePosition(dy, dx int) {
	positions := make(map[int][]int)

	s.Rightmost = 0
	s.Highest = 0
	s.Leftmost = math.MaxInt
	s.Lowest = math.MaxInt
	for y, xs := range s.Rocks {
		var newXs []int
		for _, x := range xs {
			newX := x + dx
			newXs = append(newXs, newX)

			if newX > s.Rightmost {
				s.Rightmost = newX
			}

			if newX < s.Leftmost {
				s.Leftmost = newX
			}
		}

		newY := y + dy
		positions[newY] = newXs

		if newY < s.Lowest {
			s.Lowest = newY
		}

		if newY > s.Highest {
			s.Highest = newY
		}
	}

	s.Rocks = positions
}

var shapes = []Shape{
	// ####
	{
		Rocks: map[int][]int{
			0: {0, 1, 2, 3},
		},
	},
	//.#.
	//###
	//.#.
	{
		Rocks: map[int][]int{
			0: {1},
			1: {0, 1, 2},
			2: {1},
		},
	},
	//..#
	//..#
	//###
	{
		Rocks: map[int][]int{
			0: {0, 1, 2},
			1: {2},
			2: {2},
		},
	},
	//#
	//#
	//#
	//#
	{
		Rocks: map[int][]int{
			0: {0},
			1: {0},
			2: {0},
			3: {0},
		},
	},
	//##
	//##
	{
		Rocks: map[int][]int{
			0: {0, 1},
			1: {0, 1},
		},
	},
}

func parseInput(line string) []movement {
	var movements []movement
	for _, dir := range line {
		if dir == '>' {
			movements = append(movements, right)
		} else {
			movements = append(movements, left)
		}
	}

	return movements
}

func getMove(move movement) (int, int) {
	if move == left {
		return 0, -1
	} else if move == right {
		return 0, 1
	} else { //down
		return -1, 0
	}
}

func print(rocks map[int][]int, highest int, lowest int) {
	fmt.Println()
	for y := highest; y >= lowest; y -= 1 {
		for x := 0; x < 7; x += 1 {
			_, hasX := rocks[y]
			if hasX && Contains(rocks[y], x) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func partOne(movents []movement) int {
	allRocks := make(map[int][]int)
	highest := 0
	movementIndex := 0

	for i := 0; i < 2022; i += 1 {
		shape := shapes[i%5].Clone()

		if i == 0 {
			shape.SetStartingPosition(highest + 3)
		} else {
			shape.SetStartingPosition(highest + 4)
		}

		fmt.Println("processing shape:", i+1)
		for {
			_, dx := getMove(movents[movementIndex])

			//move left or right
			if shape.CanMove(0, dx, allRocks) {
				shape.UpdatePosition(0, dx)
			}

			movementIndex += 1
			if movementIndex == len(movents) {
				fmt.Println("looping wind")
				movementIndex = 0
			}

			dy, _ := getMove(down)
			// move down
			if !shape.CanMove(dy, 0, allRocks) {
				shape.AddToMap(allRocks)

				if shape.Highest > highest {
					highest = shape.Highest
				}
				break
			}

			shape.UpdatePosition(dy, 0)
		}
	}

	return highest + 1
}

func partTwo(movents []movement) int {
	numShapes := 2022

	//numShapes := 1_000_000_000_000
	numMovements := len(movents)

	temp := numMovements * 5 // 5 different shapes * number of movements

	baseIterations, remainder := numShapes/5, numShapes%5

	fmt.Println("temp", temp, "base", baseIterations, "rem", remainder)

	calcHeight := func(numShapes int, allRocks map[int][]int) (int, map[int][]int) {
		movementIndex := 0
		highest := 0
		startingHeight := 0
		if allRocks == nil {
			allRocks = make(map[int][]int)
		} else {
			highest = len(allRocks) - 1
			startingHeight = highest
		}

		for i := 0; i < numShapes; i += 1 {
			shape := shapes[i%5].Clone()

			if len(allRocks) == 0 {
				shape.SetStartingPosition(highest + 3)
			} else {
				shape.SetStartingPosition(highest + 4)
			}

			for {
				_, dx := getMove(movents[movementIndex])

				//move left or right
				if shape.CanMove(0, dx, allRocks) {
					shape.UpdatePosition(0, dx)
				}

				movementIndex += 1

				if movementIndex == len(movents) {
					if i%5 == 0 {
						fmt.Println("Repeat on ", i)
					}
					movementIndex = 0
				}

				dy, _ := getMove(down)
				// move down
				if !shape.CanMove(dy, 0, allRocks) {
					shape.AddToMap(allRocks)

					if shape.Highest > highest {
						highest = shape.Highest
					}
					break
				}

				shape.UpdatePosition(dy, 0)
			}
		}

		coveredX := []bool{false, false, false, false, false, false, false}
		allCovered := func(xs []bool) bool {
			for _, x := range xs {
				if !x {
					return false
				}
			}

			return true
		}

		minRocks := make(map[int][]int)
		lowest := 0
		for y := highest; y >= 0; y -= 1 {
			if xs, ok := allRocks[y]; ok {
				for _, x := range xs {
					coveredX[x] = true
					_, hasX := minRocks[y]
					if !hasX {
						minRocks[y] = make([]int, 0)
					}

					minRocks[y] = append(minRocks[y], x)
				}
			}

			if allCovered(coveredX) {
				lowest = y
				break
			}
		}

		print(minRocks, highest, lowest)

		return highest - startingHeight, minRocks
	}

	base, board := calcHeight(5, nil)
	others, board := calcHeight(5, board)
	rem, _ := calcHeight(remainder, board)

	fmt.Println("base", base, "others", others, "rem", rem)
	fmt.Println(base + (others * (baseIterations - 1)) + rem)
	return 0
}

func main() {
	lines := ParseLinesWithoutEndNewLine("./sample.txt")

	movements := parseInput(lines[0])
	// Time(func() {
	// 	fmt.Println("Part 1:", partOne(movements))
	// })

	Time(func() {
		fmt.Println("Part 2:", partTwo(movements))
	})
}
