package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	. "github.com/bblasbergjc/aoc-2022/util"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Parent   *Directory
	Children map[string]*Directory
	Name     string
	Files    []*File
	Size     int
}

func NewDirectory() *Directory {
	return &Directory{
		Children: make(map[string]*Directory),
		Files:    make([]*File, 0),
	}
}

func (d *Directory) Up() *Directory {
	return d.Parent
}

func (d *Directory) Mkdir(name string) {
	d.Children[name] = &Directory{
		Name:     name,
		Parent:   d,
		Children: make(map[string]*Directory),
		Files:    make([]*File, 0),
	}
}

func (d *Directory) Cd(name string) *Directory {
	return d.Children[name]
}

func (d *Directory) CalculateSize() int {
	size := 0
	for i := range d.Files {
		size += d.Files[i].Size
	}

	for name := range d.Children {
		if d.Children[name].Size == 0 {
			d.Children[name].CalculateSize()
		}

		size += d.Children[name].Size
	}

	d.Size = size
	return size
}

func determineType(line string) string {
	if strings.HasPrefix(line, "$") {
		return "COMMAND"
	} else if strings.HasPrefix(line, "dir") {
		return "DIRECTORY"
	} else { //assume file otherwise
		return "FILE"
	}
}

// ls - read directories and files until we see a command
// cd if .. call the up command, otherwise use directory.Cd(name)
func partOne(head *Directory) int {
	// iterate and find directories with size <= 100000
	pwd := head
	totalSize := 0

	var dfs func(*Directory)
	dfs = func(root *Directory) {
		for i := range root.Children {
			dir := root.Children[i]

			if dir.Size <= 100000 {
				totalSize += dir.Size
			}

			dfs(dir)
		}
	}
	dfs(pwd)

	return totalSize
}

func partTwo(head *Directory) int {
	pwd := head.Cd("/")
	const totalSpace = 70000000
	const freeSpaceNeeded = 30000000
	unusedSpace := totalSpace - pwd.Size
	spaceToDelete := freeSpaceNeeded - unusedSpace

	smallestSize := math.MaxInt
	var dfs func(*Directory)
	dfs = func(root *Directory) {
		for i := range root.Children {
			dir := root.Children[i]

			if dir.Size >= spaceToDelete && dir.Size < smallestSize {
				smallestSize = dir.Size
			}

			dfs(dir)
		}
	}
	dfs(pwd)

	return smallestSize
}

func createFileSys(lines []string) *Directory {
	head := NewDirectory()
	head.Mkdir("/")
	pwd := head
	for i := 0; i < len(lines)-1; i += 1 {
		line := lines[i]
		lineType := determineType(line)

		if lineType == "COMMAND" {
			tokens := strings.Split(line, " ")
			cmd := tokens[1]

			if cmd == "cd" {
				dir := tokens[2]
				if dir == ".." {
					pwd = pwd.Up()
				} else {
					pwd = pwd.Cd(dir)
				}
			}
			// do nothing for "ls" command, we will read its output in next lines
		} else if lineType == "DIRECTORY" {
			items := strings.Split(line, " ")
			pwd.Mkdir(items[1])
		} else if lineType == "FILE" {
			items := strings.Split(line, " ")
			size, err := strconv.Atoi(items[0])
			CheckErr(err)

			pwd.Files = append(pwd.Files, &File{
				Name: items[1],
				Size: size,
			})
		} else {
			panic("Unknown type: " + lineType)
		}
	}

	// calulate all sizes
	head.CalculateSize()

	return head
}

func main() {
	lines := ParseLines("./day7.txt")
	filesys := createFileSys(lines)

	fmt.Println("Part 1: ", partOne(filesys))
	fmt.Println("Part 2: ", partTwo(filesys))
}
