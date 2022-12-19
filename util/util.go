package util

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseLines(file string) []string {
	data, err := os.ReadFile(file)
	CheckErr(err)
	lines := strings.Split(string(data), "\n")
	return lines
}

func ParseLinesWithoutEndNewLine(file string) []string {
	lines := ParseLines(file)
	return lines[:len(lines)-1]
}

func Time(f func()) {
	start := time.Now()
	f()
	fmt.Println("Took:", (time.Since(start)), "ms")
}
