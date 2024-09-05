package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	day := "y15d07"
	prompt := getInput(day)

	part1(prompt)
	// assumption - instructions executed in order and all values start at 0
	// WRONGGGG
	// next assumption - all orders executed simultaniously because we cba to
	// figure out the first wire. Instead we store a graph of references
	// Oh no I need structs don't I.
	// part2(prompt[0])
}

type wire struct {
	name     string
	instr    string //ASS/NOT has either value or lparen | SHIFTS always have lparen and val | AND/OR always have an rparen and either val or lparen
	solved   bool
	solution uint16
	valA     uint16
	valB     uint16
	refA     string
	refB     string
}

func part1(p []string) {
	mem := createMemory(p)
	fmt.Println(mem)
	fmt.Println()
	p1, _ := recursiveSolve(mem, "a", make(map[string]bool))
	fmt.Println("Part 1: ", p1)

	mem = createMemory(p)
	mem["b"] = wire{name: "b", instr: "ASS", solved: true, solution: p1}
	fmt.Println(mem)

	p2, _ := recursiveSolve(mem, "a", make(map[string]bool))
	fmt.Println("Part 2: ", p2)
	// // 16076 too high |

}

func recursiveSolve(mem map[string]wire, wire_name string, chain map[string]bool) (uint16, map[string]wire) {
	node := mem[wire_name]
	// fmt.Println(node)
	// fmt.Scan()

	if node.solved {
		return node.solution, mem
	}

	// if chain[wire_name] {
	// 	fmt.Println(wire_name)
	// 	fmt.Println(chain)

	// 	// panic("infinite loop detected with wire_name")
	// }
	// chain[wire_name] = true

	valA := node.valA
	if node.refA != "" {
		valA, mem = recursiveSolve(mem, node.refA, chain)
		node.valA = valA
		node.refA = ""
	}
	valB := node.valB
	if node.refB != "" {
		valB, mem = recursiveSolve(mem, node.refB, chain)
		node.valB = valB
		node.refB = ""
	}

	switch node.instr {
	case "ASS":
		node.solution = valA
	case "NOT":
		node.solution = ^valA
	case "AND":
		node.solution = valA & valB
	case "OR":
		node.solution = valA | valB
	case "LSHIFT":
		node.solution = valA << valB
	case "RSHIFT":
		node.solution = valA >> valB
	default:
		fmt.Println(wire_name, node, valA, valB)
		fmt.Println(chain)
		panic("unkown instruction while solving")
	}

	node.solved = true
	mem[node.name] = node
	return node.solution, mem

}

func createMemory(p []string) map[string]wire {
	var memory map[string]wire = make(map[string]wire)

	for _, line := range p {
		instr := strings.Split(line, " ")

		switch len(instr) {
		case 3:
			memory[instr[2]] = createWire(instr[2], "ASS", instr[0], "")
		case 4:
			memory[instr[3]] = createWire(instr[3], "NOT", instr[1], "")
		case 5:
			memory[instr[4]] = createWire(instr[4], instr[1], instr[0], instr[2])
		default:
			panic("unknown command")
		}
	}

	return memory
}

func createWire(name string, instr string, strA string, strB string) wire {

	valA, e := strconv.Atoi(strA)
	refA := ""
	if e != nil {
		refA = strA
	}
	valB, e := strconv.Atoi(strB)
	refB := ""
	if e != nil {
		refB = strB
	}

	if refA == "" { // if only value is already an integer, return solved wire
		switch instr {
		case "ASS":
			return wire{name: name, instr: instr, solved: true, solution: uint16(valA)}
		case "NOT":
			return wire{name: name, instr: instr, solved: true, solution: uint16(valA)}
		}
	}
	// otherwise return normal connected wire
	return wire{name: name, instr: instr, valA: uint16(valA), refA: refA, valB: uint16(valB), refB: refB}
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
