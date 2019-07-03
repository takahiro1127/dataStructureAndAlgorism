package main

import(
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	"../procon"
)

func main() {
	randomArray, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	fmt.Println(partition(randomArray))
}

func partition(A []int, p int, r int) int {
	x = A[r]
	i = p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[r], A[i+1] = A[i+1], A[r]
	return i+1
}
