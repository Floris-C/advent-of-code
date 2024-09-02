package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := "y15d05"
	prompt := getInput(day)

	part1(prompt)
	part2(prompt) // 193 too high - 189 too high - 94 not right - 69 not right - 51 is correct :D

	fmt.Println(hasTwoPair("aaa"))
}

func part1(p []string) {
	cnt := 0
	for _, s := range p {
		if isNice(s) {
			cnt++
		}
	}
	fmt.Println("Part 1: ", cnt)

}

func part2(p []string) {
	cnt := 0
	for _, s := range p {
		if isNiceTwo(s) {
			cnt++
		}
	}
	fmt.Println("Part 2: ", cnt)

}

func isNice(s string) bool {
	sr := []rune(s)
	var cprev rune
	vowels := 0
	hasRepeat := false

	for _, c := range sr {
		if strings.Contains("aeiou", string(c)) {
			vowels++
		}

		if cprev == c {
			hasRepeat = true
		}

		if (cprev == 'a' && c == 'b') || (cprev == 'c' && c == 'd') || (cprev == 'p' && c == 'q') || (cprev == 'x' && c == 'y') {
			return false
		}

		cprev = c
	}
	// fmt.Print(hasRepeat)
	return ((vowels >= 3) && hasRepeat)
}

func isNiceTwo(s string) bool {
	return hasTwoPair(s) && hasSkipRepeat(s)
}

func hasTwoPair(s string) bool {
	sr := []rune(s)
	a := sr[0]
	for i, b := range sr[1:] {
		fmt.Println(i)
		x := sr[i+2]
		if i+3 >= len(sr) {
			return false
		}
		for _, y := range sr[i+3:] {
			if (a == x) && (b == y) {
				return true
			}
			x = y
		}
		a = b
	}

	return false
}

func hasSkipRepeat(s string) bool {
	sr := []rune(s)
	var cpp rune
	var cp rune
	for _, c := range sr {
		if cpp == c {
			return true
		}
		cpp = cp
		cp = c
	}
	return false
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
