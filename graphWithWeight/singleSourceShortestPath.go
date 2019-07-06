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
	d := make([]int, n)
	reloadD := make([]int, n)
	p := make([]int, n)
	S := make([]bool, n)
	for i := 0; i < n; i++ {
		d[i] = 3000
		reloadD[i] = 3000
		S[i] = false
		p[i] = -1
		connectionInfo, err := procon.SplitIntStdin(" ")
		// fmt.Println(min(connectionInfo))
		if err != nil {
			panic(err)
		}
		m := connectionInfo[0]
		k := connectionInfo[1]
		connection := make([]int, n)
		for j := 1; j <= k; j++ {
			connection[connectionInfo[2 * j]] = connectionInfo[2 * j + 1]
			// fmt.Printf("%d %d\n", connectionInfo[2 * j], connectionInfo[2 * j + 1])
		}
		graph[m] = connection
	}

	S[0] = true
	d[0] = 0
	reloadD[0] = 3000
	for i:= 1; i < n; i++ {
		if d[i] > graph[0][i] && graph[0][i] != 0 {
			d[i] = graph[0][i]
			p[i] = 0
			reloadD[i] =  graph[0][i]
		}
	}
	for i:= 1; i < n; i++ {
		_, newS := min(reloadD)
		S[newS] = true;
		for j := 1; j < n; j++ {
			if d[j] > graph[newS][j] + d[newS] && graph[newS][j] != 0 {
				d[j] = graph[newS][j] + d[newS]
				reloadD[j] = d[j]
				p[j] = newS
			}
		}
		reloadD[i] = 3000
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i, d[i])
	}
	// fmt.Println(min(graph[4]))
}

func min(A []int) (int, int) {
	var index int
	var change = false
	for i := len(A) - 1; i > 0; i-- {
		if A[i] < A[i - 1] {
			if !change {
				index = i
				change = true
			}
			A[i], A[i - 1] = A[i - 1], A[i]
		} else {
			index = i - 1
			change = false
		}
	}
	return A[0], index
}
