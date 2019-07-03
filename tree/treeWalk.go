package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"github.com/k0kubun/pp"
	"../procon"
)

type Node struct{
	parent, left, right, value int
}

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	binaryTree := make([]Node, n, n)
	for i := 0; i < n; i++ {
		binaryTree[i].parent = -1
	}
	for i := 0; i < n; i++ {
		nodeInfo, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		binaryTree[nodeInfo[0]].value = nodeInfo[0]
		if nodeInfo[1] != -1 {
			binaryTree[nodeInfo[0]].left = nodeInfo[1]
			binaryTree[nodeInfo[1]].parent = nodeInfo[0]
		} else {
			binaryTree[nodeInfo[0]].left = nodeInfo[1]
		}
		if nodeInfo[2] != -1 {
			binaryTree[nodeInfo[0]].right = nodeInfo[2]
			binaryTree[nodeInfo[2]].parent = nodeInfo[0]
		} else {
			binaryTree[nodeInfo[0]].right = nodeInfo[2]
		}
	}
	pp.Println(binaryTree)
	fmt.Println(postParse(binaryTree[0], binaryTree))
}

func preParse(node Node, binaryTree []Node) []int {
	var order []int
	order = append(order, node.value)
	if node.left != -1 {
		order = append(order, preParse(binaryTree[node.left], binaryTree)...)
	}
	if node.right != -1{
		order = append(order, preParse(binaryTree[node.right], binaryTree)...)
	}
	return order
}

func inParse(node Node, binaryTree []Node) []int {
	var order []int
	if node.left == -1 {
		order = append(order, node.value)
		return order
	} else {
		order = append(order, inParse(binaryTree[node.left], binaryTree)...)
		order = append(order, node.value)
		order = append(order, inParse(binaryTree[node.right], binaryTree)...)
		return order
	}
}

func postParse(node Node, binaryTree []Node) []int {
	var order []int
	if node.left == -1 {
		order = append(order, node.value)
		return order
	} else {
		order = append(order, inParse(binaryTree[node.left], binaryTree)...)
		order = append(order, inParse(binaryTree[node.right], binaryTree)...)
		order = append(order, node.value)
		return order
	}
}




func height(binaryTree []Node) []int {
	heights := make([]int, len(binaryTree), len(binaryTree))
	for i:= 0; i < len(binaryTree); i++ {
		heights[i] = nodeHeight(binaryTree[i], binaryTree)
	}
	return heights
}

func nodeHeight(node Node, binaryTree []Node) int {
	if node.parent == -1 { return 1}
	height := 1 + nodeHeight(binaryTree[node.parent], binaryTree)
	return height
}


