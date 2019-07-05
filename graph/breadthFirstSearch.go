package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "github.com/k0kubun/pp"
	"../procon"
)

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		nodeInfo, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		sides := make([]int, n)
		for i := 2; i < len(nodeInfo); i++ {
			sides[nodeInfo[i] - 1] = 1
		}
		graph[nodeInfo[0] - 1] = sides;
	}
	distances := make([]int, len(graph))
	fmt.Println(minimumDistance(0, graph, distances))
}

func minimumDistance(index int, graph [][]int, distances []int) []int {
	for i := 0; i < len(graph); i++ {
		if graph[index][i] == 1 && distances[i] == 0{
			distances[i] += distances[index] + 1
		}
	}
	for i := 0; i < len(graph); i++ {
		if graph[index][i] == 1 {
			newDistance := minimumDistance(i, graph, distances)
			for i := 0; i < len(graph); i++ {
				distances[i] = newDistance[i]
			}
		}
	}
	return distances
}
