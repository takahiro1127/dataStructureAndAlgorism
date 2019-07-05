package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "github.com/k0kubun/pp"
	"../procon"
)

type Node struct{
	sides []int
	visitedSides []int
	come, finish int
}

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	graph := make([]Node, n)
	for i := 0; i < n; i++ {
		nodeInfo, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		sides := make([]int, n)
		setVisitedSides := make([]int, n)
		for i := 2; i < len(nodeInfo); i++ {
			sides[nodeInfo[i] - 1] = 1
			setVisitedSides[nodeInfo[i] - 1] = 1
		}
		node := Node {
			sides: sides,
			visitedSides : setVisitedSides,
		}
		graph[nodeInfo[0] - 1] = node
	}
	var time int
	for i := 0; i < len(graph); i++ {
		time++
		time = graph[i].depthFirstSearch(graph, time)
		fmt.Printf("%d %d %d\n", i, graph[i].come, graph[i].finish)
	}
}

func (node *Node)depthFirstSearch( graph []Node, time int) int {
	if node.come == 0 {
		node.come = time
	}
	for i := 0; i < len(node.sides) && graph[node.sides[i]].come == 0 &&  node.visitedSides[i] == 1; i++ {
		time++
		time = graph[node.sides[i]].depthFirstSearch(graph, time)
		node.visitedSides[i] = 0
	}
	node.finish = time
	return time
}
