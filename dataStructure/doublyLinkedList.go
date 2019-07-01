package main

import (
	// "bufio"
	// "os"
	// "strconv"
	// "strings"
	"github.com/k0kubun/pp"
	// "fmt"
	// "../procon"
)

type LinkedList struct {
	nodes []Node
	head int //nextIndexで指定
	tail int //nextIndexで指定
}

type Node struct {
	key int
	nextIndex int
}

func main() {
	node1 := Node{
		nextIndex: 1,
		key: 2,
	}
	node2 := Node{
		nextIndex: 2,
		key: 4,
	}
	node3 := Node{
		nextIndex: 0,
		key: 8,
	}
	nodes := []Node {node1, node2, node3}
	linkedList := LinkedList{
		nodes: nodes,
		head: 0,
		tail: 2,
	}
	linkedList.insert(16)
	linkedList.insert(32)
	linkedList.delete(8)
	pp.Print(linkedList)
}

func (linkedList LinkedList) insert(key int) {
	node := Node{
		nextIndex: linkedList.head,
		key: key,
	}
	linkedList.head = len(linkedList.nodes)
	linkedList.nodes = append(linkedList.nodes, node)
}

func (linkedList LinkedList) findFirstKey(key int) int int {
	index := linkedList.head
	var beforeIndex int
	for {
		if linkedList.nodes[index].key == key {
			return beforeIndex, index
		} else if index == linkedList.tail {
			panic("key not found")
		} else {
			beforeIndex = index
			index = linkedList.nodes[index].nextIndex
		}
	}
}

func (linkedList LinkedList) delete(key int) {
	beforeIndex, index := linkedList.findFirstKey(key)
	if index == linkedList.head {
		linkedList.head = linkedList.nodes[index].nextIndex
	} else if index == linkedList.tail {
		linkedList.tail = beforeIndex
	}
	linkedList.nodes[beforeIndex].nextIndex = linkedList.nodes[index].nextIndex
}
