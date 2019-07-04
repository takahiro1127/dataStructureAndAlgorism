package main

import(
	// "bufio"
	// "fmt"
	// "os"
	"strconv"
	"github.com/k0kubun/pp"
	"../procon"
)

type Node struct{
	parent, left, right, value, place int
}

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	binaryTree := make([]Node, n, n)
	for i := 0; i < n; i++ {
		binaryTree[i].parent = -1
		binaryTree[i].value = -1
		binaryTree[i].right = -1
		binaryTree[i].left = -1
	}
	for i := 0; i < n; i++ {
		nodeValue, err := strconv.Atoi(procon.StrStdin())
		if err != nil {
			panic(err)
		}
		binaryTree[i].insert(nodeValue, binaryTree)
	}
	pp.Println(binaryTree)
	// fmt.Println(postParse(binaryTree[0], binaryTree))
}

func (node *Node)insert(nodeValue int, binaryTree []Node) {
	i := empty(binaryTree)
	if node.value == -1 &&  i == 0{
		node.value = nodeValue
		node.place = i
		return
	}
	if node.value < nodeValue {
		if node.right != -1 {
			binaryTree[node.right].insert(nodeValue, binaryTree)
		} else {
			node.right = i
			pp.Print("書き換えてるよ")
			binaryTree[i].value = nodeValue
			binaryTree[i].parent = node.place
			binaryTree[i].place = i
		}
	} else {
		if node.left != -1 {
			binaryTree[node.right].insert(nodeValue, binaryTree)
		} else {
			node.left = i
			pp.Print("書き換えてるよ")
			binaryTree[i].value = nodeValue
			binaryTree[i].parent = node.place
			binaryTree[i].place = i
		}
	}
}

func empty(binaryTree []Node) int {
	for i := 0; i < len(binaryTree); i++ {
		if binaryTree[i].value == -1 && binaryTree[i].parent == -1 {
			return i
		}
	}
	return len(binaryTree)
}
