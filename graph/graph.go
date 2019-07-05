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
		node := make([]int, n)
		for i := 2; i < len(nodeInfo); i++ {
			node[nodeInfo[i] - 1] = 1
		}
		graph[nodeInfo[0] - 1] = node
		fmt.Println(node)
	}
}
