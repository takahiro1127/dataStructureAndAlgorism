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
	matrixes := make([][2]int, n)
	for i := 0; i < n; i++ {
		matrix, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		matrixes[i][0], matrixes[i][1] = matrix[0], matrix[1]
	}
	chainMultipleCount(matrixes)
}

func chainMultipleCount([][2]int) {
	
}
