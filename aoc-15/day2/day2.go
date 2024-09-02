package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	day := "y15d02"
	prompt := getInput(day)

	part1(prompt)
	part2(prompt)
}

func part1(p []string) {

	total := 0
	for _, line := range p {
		sides := strings.SplitN(line, "x", 3)

		l, errl := strconv.Atoi(sides[0])
		w, errw := strconv.Atoi(sides[1])
		h, errh := strconv.Atoi(sides[2])
		for _, err := range []error{errl, errw, errh} {
			if err != nil {
				fmt.Println("panic", l, w, h)
			}
		}

		side1 := l * w
		side2 := w * h
		side3 := l * h

		size := 2*(side1+side2+side3) + min(side1, side2, side3)
		total = total + size
	}

	fmt.Println("Part 1: ", total)
}

func part2(p []string) {
	total := 0
	for _, line := range p {
		sides := strings.SplitN(line, "x", 3)

		l, errl := strconv.Atoi(sides[0])
		w, errw := strconv.Atoi(sides[1])
		h, errh := strconv.Atoi(sides[2])
		for _, err := range []error{errl, errw, errh} {
			if err != nil {
				fmt.Println("panic", l, w, h)
			}
		}

		orderedSides := []int{l, w, h}
		slices.Sort(orderedSides)
		ribbon := orderedSides[0]*2 + orderedSides[1]*2
		bow := l * w * h

		total += ribbon + bow
	}

	fmt.Println("Part 2: ", total)
}

func getInput(day string) []string {
	f := strings.Join([]string{day, ".txt"}, "")
	p := filepath.Join(".\\.\\inputs", f)
	fmt.Println(p)

	dat, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\r\n")
	return lines
}
