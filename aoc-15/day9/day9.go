package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	day := "y15d09"
	prompt := getInput(day)

	// part1(prompt) // ! TOO LOW: 168, 183
	part1again(prompt)
	part2(prompt)
}

func part1(p []string) {
	// Create a graph based on the inputs
	g := createGraph(p)
	visualiseGraph(g, "full")
	// Derive the "minimum spanning tree" using ```func MinimumSpanningTree```
	mst, e := graph.MinimumSpanningTree(g)
	checkPanic(e)
	visualiseGraph(mst, "mst")
	edges, e := mst.Edges()
	checkPanic(e)
	weightSum := 0
	for _, edge := range edges {
		weightSum += edge.Properties.Weight
	}
	fmt.Println(weightSum)
}

func part1again(p []string) {

	city_list, city_dist_map := parse_input(p)

	fmt.Println((city_list))
	fmt.Println((city_dist_map))

	shortest_path := 9999999
	for i, starting_city := range city_list {
		fmt.Println("attempting", starting_city, "cur_cities", city_list)
		to_visit := make([]string, len(city_list))
		_ = copy(to_visit, city_list)
		path_dist := recursive_city_dist(starting_city, append(to_visit[:i], to_visit[i+1:]...), 0, city_dist_map)
		shortest_path = min(path_dist, shortest_path)
	}
	fmt.Println("Solution part 1: ", shortest_path)
}

func parse_input(p []string) ([]string, map[string]map[string]int) {
	city_list := []string{}
	city_dist_map := make(map[string]map[string]int)

	for _, l := range p {
		split_line := strings.Split(l, " ")
		city_A, city_B, dist := split_line[0], split_line[2], split_line[4]

		if !slices.Contains(city_list, city_A) {
			city_list = append(city_list, city_A)
		}
		if !slices.Contains(city_list, city_B) {
			city_list = append(city_list, city_B)
		}

		dist_int, e := strconv.Atoi(dist)
		checkPanic(e)

		if _, ok := city_dist_map[city_A]; !ok {
			city_dist_map[city_A] = make(map[string]int)
		}
		city_dist_map[city_A][city_B] = dist_int
		if _, ok := city_dist_map[city_B]; !ok {
			city_dist_map[city_B] = make(map[string]int)
		}
		city_dist_map[city_B][city_A] = dist_int

	}

	return city_list, city_dist_map
}

func part2(p []string) {
	city_list, city_dist_map := parse_input(p)

	fmt.Println((city_list))
	fmt.Println((city_dist_map))

	longest_path := 0
	for i, starting_city := range city_list {
		fmt.Println("attempting", starting_city, "cur_cities", city_list)
		to_visit := make([]string, len(city_list))
		_ = copy(to_visit, city_list)
		path_dist := recursive_city_dist_maximalist(starting_city, append(to_visit[:i], to_visit[i+1:]...), 0, city_dist_map)
		longest_path = max(path_dist, longest_path)
	}
	fmt.Println("Solution part 2: ", longest_path)
}

func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func recursive_city_dist(city string, to_visit []string, dist_travelled int, city_dist_map map[string]map[string]int) int {
	// fmt.Println(to_visit, dist_travelled)
	if len(to_visit) == 0 {
		return dist_travelled
	}
	shortest_dist := 99999
	for i, next_city := range to_visit {
		next_dist := city_dist_map[city][next_city]
		next_visits := make([]string, len(to_visit))
		_ = copy(next_visits, to_visit)

		new_dist := recursive_city_dist(next_city, append(next_visits[:i], next_visits[i+1:]...), dist_travelled+next_dist, city_dist_map)
		// fmt.Println("retained?", to_visit)
		shortest_dist = min(shortest_dist, new_dist)
	}
	return shortest_dist
}

func recursive_city_dist_maximalist(city string, to_visit []string, dist_travelled int, city_dist_map map[string]map[string]int) int {
	// fmt.Println(to_visit, dist_travelled)
	if len(to_visit) == 0 {
		return dist_travelled
	}
	longest_dist := 0
	for i, next_city := range to_visit {
		next_dist := city_dist_map[city][next_city]
		next_visits := make([]string, len(to_visit))
		_ = copy(next_visits, to_visit)

		new_dist := recursive_city_dist_maximalist(next_city, append(next_visits[:i], next_visits[i+1:]...), dist_travelled+next_dist, city_dist_map)
		// fmt.Println("retained?", to_visit)
		longest_dist = max(longest_dist, new_dist)
	}
	return longest_dist
}

func createGraph(p []string) graph.Graph[string, string] {
	g := graph.New(graph.StringHash)
	for _, l := range p {
		split_line := strings.Split(l, " ")
		city_a, city_b, distance := split_line[0], split_line[2], split_line[4]
		// add cities as a vertex (vertexes cannot be duplicated so its fine)
		g.AddVertex(city_a)
		g.AddVertex(city_b)
		// add a weighted edge between cities
		dist_int, err := strconv.Atoi(distance)
		if err != nil {
			panic("error parsing distance")
		}
		g.AddEdge(city_a, city_b, graph.EdgeWeight(dist_int))
	}
	return g
}

func visualiseGraph(g graph.Graph[string, string], n string) {
	file_name := strings.Join([]string{n, "_graph.gv"}, "")
	file, _ := os.Create(file_name)
	_ = draw.DOT(g, file)
}

func getInput(day string) []string {
	f := strings.Join([]string{day, ".txt"}, "")
	p := filepath.Join(".\\.\\inputs", f)
	fmt.Println("importing: ", p)

	dat, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\r\n")
	return lines
}
