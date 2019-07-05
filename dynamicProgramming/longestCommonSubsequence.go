package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "github.com/k0kubun/pp"
	"../procon"
)
var fibonacciNumberStock []int

var lcs [][]int
func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		lcs = nil
		x := SplitStrStdin(" ")
		y := SplitStrStdin(" ")
		fmt.Println(longestCommonSubsequence(x, y))
	}
}

func longestCommonSubsequence(x, y []string) int {
	if x[len(x)] == y[len(y)] {
		newx, newy := x[:len(x) - 1], y[:len(y) - 1]
		lcs[len(x) - 1][len(y) - 1] = longestCommonSubsequence(newx, newy) + 1
	} else {
		newx, newy := x[:len(x) - 1], y[:len(y) - 1]
		reloadx, reloady:= longestCommonSubsequence(y, newx), longestCommonSubsequence(x, newy)
		if reloadx > reloady {
			lcs[len(x) - 1][len(y) - 1] = reloadx
		} else {
			lcs[len(x) - 1][len(y) - 1] = reloady
		}
	}
	return lcs[len(x) - 1][len(y) - 1]
}
