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
	fmt.Println(countingSort(randomArray))
}

func countingSort(A []int) []int {
	B := make([]int, len(A), len(A))
	var C [2000000]int

	for j := 0; j < len(A); j++ {
		C[A[j]]++
	}

	for i := 1; i <= 2000000; i++ {
		C[i] = C[i - 1] + C[i]
	}

	for j := len(A) - 1; j > 0; j-- {
		B[C[A[j]]] = A[j]
		C[A[j]]--
	}
	return B
}
