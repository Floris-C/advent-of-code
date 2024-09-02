package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := "y15d01"
	prompt := getInput(day)

	part1(prompt)
	part2(prompt)
}

func part1(p string) {
	floor := 0
	for _, c := range p {
		switch c {
		case ')':
			floor--
		case '(':
			floor++
		}
	}
	fmt.Println("Part 1: ", floor)
}

func part2(p string) {
	floor := 0
	target := -1

	for i, c := range p {
		switch c {
		case ')':
			floor--
		case '(':
			floor++
		}
		if floor == target {
			fmt.Println("Part 2: ", i+1)
			break
		}

	}

}

func getInput(day string) string {
	f := strings.Join([]string{day, ".txt"}, "")
	p := filepath.Join(".\\inputs", f)
	fmt.Println(p)

	dat, err := os.ReadFile(p)

	if err != nil {
		panic(err)
	}

	return string(dat)
}
