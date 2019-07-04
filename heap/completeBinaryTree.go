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
	_, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	completeBinaryTree, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	print(completeBinaryTree)
}

func print(tree []int) {
	for i := 0; i < len(tree); i++ {
		fmt.Printf("Node: %d key = %d, parentKey = %d, left key = %d right key %d\n", i + 1, tree[i], (i+1)/2 , i * 2, i * 2 + 1)
	}
}
