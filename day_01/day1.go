package main

import (
	"container/heap"
	"fmt"
	"strconv"

	. "github.com/bblasbergjc/aoc-2022/util"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func partOne(lines []string) int {
	maxCalories := 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			maxCalories = max(maxCalories, currentCalories)
			currentCalories = 0
		} else {
			cals, err := strconv.ParseInt(line, 10, 64)
			CheckErr(err)

			currentCalories += int(cals)
		}
	}

	return maxCalories
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func partTwo(lines []string) int {
	h := &IntHeap{0, 0, 0}
	heap.Init(h)
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			if currentCalories > (*h)[0] { // if this value is less than the smallest in our heap
				heap.Pop(h)
				heap.Push(h, currentCalories)
			}
			currentCalories = 0
		} else {
			cals, err := strconv.ParseInt(line, 10, 64)
			CheckErr(err)

			currentCalories += int(cals)
		}
	}

	total := 0
	for h.Len() > 0 {
		total += heap.Pop(h).(int)
	}

	return total

}

func main() {
	lines := ParseLines("./day1.txt")

	fmt.Printf("Top three calories sum: %d\n", partTwo(lines))
}
