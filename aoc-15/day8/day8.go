package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := "y15d08"
	prompt := getInput(day)

	p1 := part1(prompt)
	fmt.Println("Part 1: ", p1) // 1235 too low -

	p2 := part2(prompt)
	fmt.Println("Part 2: ", p2) // 2163 too high (forgot continue after parsing \\ (would see second \ as new escape))
}

func part1(p []string) int {
	literals_count := 0
	values_count := 0
	for _, l := range p {
		// fmt.Println(l)
		literals_count += len(l)
		values_count += countStringValues(l)
		// fmt.Println(literals_count, values_count)
	}
	return literals_count - values_count
}

func part2(p []string) int {
	literals_count := 0
	reencoded_count := 0
	for _, l := range p {
		// fmt.Println(l)
		literals_count += len(l)
		reencoded_count += calcReencodedSize(l)
		// fmt.Println(literals_count, reencoded_count)
	}
	return reencoded_count - literals_count
}

func calcReencodedSize(l string) int {
	size := len(l) + 4

	l = l[1 : len(l)-1]
	escaping := false

	for _, c := range l {
		// fmt.Println(c, string(c), character_count)
		// given any unescaped normal character +1 and continue
		if escaping {
			switch c {
			case '\\':
				size += 2
			case '"':
				size += 2
			default:
				size += 1
			}
			escaping = false
			continue
		}

		escaping = c == '\\'

	}

	return size
}

func countStringValues(l string) int {
	l = l[1 : len(l)-1]
	character_count := 0
	character_window := ""

	hex_digits := "0123456789abcdef"
	// octal_digits := "01234567"
	// fmt.Println(l)

	for _, c := range l {
		// fmt.Println(c, string(c), character_count)
		// given any unescaped normal character +1 and continue

		switch character_window {
		case "":
			if c == '\\' {
				character_window = "\\"
				continue
			}
			character_count += 1
		case "\\":
			if c == 'x' {
				character_window = "\\x"
				continue
			}
			character_count += 1
			character_window = ""
		case "\\x":
			if strings.ContainsRune(hex_digits, c) {
				character_window = "\\xx"
				continue
			}
			character_count += 2
			character_window = ""
		case "\\xx":
			if strings.ContainsRune(hex_digits, c) {
				character_count++
				character_window = ""
				continue
			}
			character_count += 3
			character_window = ""
		}

	}
	return character_count
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
