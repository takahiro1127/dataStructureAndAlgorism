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
		connection, err := procon.SplitIntStdin(" ")
		for j := 0; j < len(connection); j++ {
			if connection[j] == -1 {connection[j] = 3000}
		}
		graph[i] = connection
	}
	fmt.Println(minimumSpan(graph))
}

func minimumSpan(graph [][]int) int {
	int u, minv
	n := len(graph)
	d := make([]int, n)
	color := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = 3000
		p[i] = -1
		color[i] = 0
	}

	d[0] = 0

	for i := 0; ; i++ {
		minv = 3000
		u = -1
		for i := 0; i < n; i++ {
			if minv > d[i] && color[i] != 1 {
				u = i
				minv = d[i] //一番近い隣接点を選んでいる
			}
		}
		if (u == -1) { break }
		color[u] = 1 //一番近い隣接点は全域木に入る

		for i := 0; i < n; i++ {
			if color[i] != 1 && M[u][v] <= 2000 {
				if d[i] > graph[u][i] { //全域木から近い点を抽出(最新の隣接点についてのみでいい)
					d[i] = graph[u][i]
					p[i] = u
					color[i] = GRAY
				}
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		if p[i] != -1 { sum += M[i][p[i]]}
	}

}
