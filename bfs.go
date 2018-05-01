package main

import "fmt"

// implementing a recursive bfs because golang doesn't have tuples
// and I'm too lazy to implement bfs without tuples
func bfs2(graph map[string][]string, visited map[string]bool, current string, target string) int {
	if visited[current] {
		return -1
	}

	if current == target {
		return 0
	}

	visited[current] = true
	result := -1
	for _, n := range graph[current] {
		nresult := bfs2(graph, visited, n, target)
		if nresult > -1 && (result == -1 || nresult < result) {
			result = nresult
		}
	}
	visited[current] = false

	if result == -1 {
		return -1
	}

	return result + 1
}

func bfs(graph map[string][]string, from, to string) int {
	return bfs2(graph, make(map[string]bool), from, to)
}

func main() {
	graph := make(map[string][]string)

	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"C"}
	graph["C"] = []string{"C", "D"}

	fmt.Println(bfs(graph, "A", "D"))
	fmt.Println(bfs(graph, "B", "D"))
}
