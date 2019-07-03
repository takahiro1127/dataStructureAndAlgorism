package main

import(
	// "bufio"
	// "fmt"
	// "os"
	"strconv"
	"github.com/k0kubun/pp"
	"../procon"
)

type RootedTree struct{
	parent *RootedTree
	children [3]*RootedTree
	node int
	order int
	depth int
	class string
}

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	rootedTrees := make([]RootedTree, n, n)
	for i := 0; i < n; i++ {
		randomArray, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		if i == 0{
			rootedTrees[0] = createRoot(randomArray)
			continue
		}
		rootedTrees[i] = createTree(randomArray, rootedTrees)
	}
	pp.Println(rootedTrees)
}

func createTree(A []int, rootedTrees []RootedTree) RootedTree {
	created, rootedTree := searchCreated(rootedTrees[0], A[0])
	if !created {
		rootedTree.class = "root"
	}	else {
		rootedTree.class = "internal node"
	}
	rootedTree.node = A[0]
	for i := 0; i < A[1]; i++ {
		var childTree RootedTree
		childTree.node = A[i + 2]
		childTree.class = "leaf"
		rootedTree.children[i] = &childTree
	}
	return *rootedTree
}

func searchCreated(rootedTree RootedTree, node int) (bool, *RootedTree) {
	if rootedTree.node == node { return true, &rootedTree}
	for i := 0; i < 3; i++ {
		if rootedTree.children[i].node == node {return true, rootedTree.children[i]}
		children := *rootedTree.children[i]
		result, node :=  searchCreated(children, node)
		if result {
			return true, node
		}
	}
	var root RootedTree
	return false, &root
}

func createRoot(A []int) RootedTree {
	var rootedTree RootedTree
	rootedTree.node = A[0]
	rootedTree.class = "root"
	for i := 0; i < A[1]; i++ {
		var childTree RootedTree
		childTree.node = A[i + 2]
		childTree.class = "leaf"
		rootedTree.children[i] = &childTree
	}
	return rootedTree
}
