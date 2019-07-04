package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"github.com/k0kubun/pp"
	"../procon"
)

type Node struct {
	key int
	parent, left, right *Node
}

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	var root Node
	for i := 0; i < n; i++ {
		order := procon.SplitStrStdin(" ")
		method := order[0]
		target, err := strconv.Atoi(order[1])
		if err != nil {
			panic(err)
		}

		if method == "insert" {
			if i == 0 {
				root.key = target
				continue
			}
			root.insert(target)
		} else if method == "find" {
			fmt.Println(root.find(target))
		} else if method == "print" {
			root.preParse()
			root.inParse()
		} else if method == "delete" {
			root.delete(target)
		} else {
			fmt.Println("not exist such a order")
		}
	}
}

func (node Node)insert(key int) {
	if node.key < key && node.right == nil {
		newNode := Node {
			key: key,
			parent: &node,
		}
		node.right = &newNode
		pp.Print(node)
		pp.Print(newNode)
	} else if node.key < key {
		node.right.insert(key)
	} else if node.left == nil {
		newNode := Node{
			key: key,
			parent: &node,
		}
		node.left = &newNode
		pp.Print(node)
		pp.Print(newNode)
	} else {
		node.left.insert(key)
	}
}

func (node Node)find(key int) bool {
	if node.key == key {
		return true
	} else if node.key < key {
		if node.right == nil { return false }
		return node.right.find(key)
	} else {
		if node.left == nil { return false }
		return node.left.find(key)
	}
}

func (node Node)preParse() {

}

func (node Node)inParse() {

}

func (node Node)delete(key int) {
	if node.key == key {
		if node.left == nil && node.right != nil {
			node.parent.left = node.right
		} else if node.left == nil && node.right == nil {
			return
		}
		node.key = node.left.key
		node.left.delete(node.left.key)
		pp.Print(node)
	} else if node.key < key {
		if node.right == nil { 
			fmt.Println("ここいる1")
			return }
		node.right.delete(key)
	} else {
		if node.left == nil { 
			fmt.Println("ここいる2")
			pp.Print(node)
			return }
		node.left.delete(key)
	}
}
