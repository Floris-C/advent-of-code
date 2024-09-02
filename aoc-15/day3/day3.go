package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := "y15d03"
	prompt := getInput(day)

	part1(prompt[0])
	part2(prompt[0])
}

func part1(p string) {
	var m map[int]map[int]bool = make(map[int]map[int]bool)
	x, y := 0, 0

	for _, r := range p {
		switch r {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		}

		if _, ok := m[x]; !ok {
			m[x] = make(map[int]bool)
		}
		m[x][y] = true
	}

	houses := 0
	for _, r := range m {
		houses += len(r)
	}

	fmt.Println("Part 1: ", houses)
}

func part2(p string) {
	var m map[int]map[int]bool = make(map[int]map[int]bool)
	m[0] = make(map[int]bool)
	m[0][0] = true

	x, y := 0, 0
	rx, ry := 0, 0
	normalSanta := true

	for _, r := range p {
		switch r {
		case '<':
			if normalSanta {
				x--
			} else {
				rx--
			}
		case '>':
			if normalSanta {
				x++
			} else {
				rx++
			}
		case '^':
			if normalSanta {
				y--
			} else {
				ry--
			}
		case 'v':
			if normalSanta {
				y++
			} else {
				ry++
			}
		}

		var nx, ny int
		if normalSanta {
			nx, ny = x, y
		} else {
			nx, ny = rx, ry
		}

		if _, ok := m[nx]; !ok {
			m[nx] = make(map[int]bool)
		}
		m[nx][ny] = true

		normalSanta = !normalSanta
	}

	houses := 0
	for _, r := range m {
		houses += len(r)
	}

	fmt.Println("Part 2: ", houses)
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
