package main

import "fmt"

var graph = map[string]map[string]bool{}

func main() {
	addEdge("Kirov", "Moscow")
	addEdge("Kirov", "Saint-Petersburg")
	addEdge("Minsk", "London")
	fmt.Println(hasEdge("Minsk", "London"))
	fmt.Println(hasEdge("Minsk", "Kirov"))
	fmt.Println(hasEdge("Samara", "Kirov"))
}

func addEdge(from, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
		graph[from] = edge
	}
	edge[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
