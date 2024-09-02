package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	day := "y15d04"
	fmt.Println(day)
	prompt := "iwrupvqb"

	part1(prompt)
	// part2(prompt[0])
}

func part1(key string) {
	answer := 1
	for {
		input := []byte(strings.Join([]string{key, strconv.Itoa(answer)}, ""))
		hash := md5.Sum(input)

		if hex.EncodeToString(hash[:3])[:6] == "000000" {
			fmt.Println(answer)
			break
		}
		answer++

	}

}
