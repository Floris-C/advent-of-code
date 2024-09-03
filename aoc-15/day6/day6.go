package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	day := "y15d06"
	prompt := getInput(day)

	part1(prompt) // 743952 too high (probaply blame on toggle) // 377891 correct
	// part2(prompt[0])
}

func part1(p []string) {
	var m map[int]map[int]int = make(map[int]map[int]int)

	for _, line := range p {
		instr, lx, ly, rx, ry := decodeInstruction(line)

		for x := lx; x <= rx; x++ {
			if _, ok := m[x]; !ok {
				m[x] = make(map[int]int)
			}

			for y := ly; y <= ry; y++ {
				switch instr {
				case "on":
					m[x][y] += 1
				case "off":
					if m[x][y] > 0 {
						m[x][y] -= 1
					}
				case "toggle":
					m[x][y] += 2
				}

			}

		}
	}

	fmt.Println("Part 1: ", countTrues(m))

}

func countTrues(m map[int]map[int]int) int {
	counter := 0
	for _, row := range m {
		for _, val := range row {
			counter += val
		}
	}
	return counter
}

func decodeInstruction(l string) (string, int, int, int, int) {
	l_sliced := strings.Split(l, " ")

	var instr string
	var coordI int
	if l_sliced[0] == "toggle" {
		instr = l_sliced[0]
		coordI = 1
	} else {
		instr = l_sliced[1]
		coordI = 2
	}

	l_coords := strings.Split(l_sliced[coordI], ",")
	lx, lxe := strconv.Atoi(l_coords[0])
	ly, lye := strconv.Atoi(l_coords[1])

	r_coords := strings.Split(l_sliced[coordI+2], ",")
	rx, rxe := strconv.Atoi(r_coords[0])
	ry, rye := strconv.Atoi(r_coords[1])

	for _, e := range []error{lxe, lye, rxe, rye} {
		if e != nil {
			panic(e)
		}
	}

	return instr, lx, ly, rx, ry
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
