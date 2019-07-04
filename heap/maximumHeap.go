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
	_, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	randomArray, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buildMaxHeap(randomArray))
}

func buildMaxHeap(A []int) []int{
	for i := (len(A) - 1)/2 - 1; i >= 0; i-- {
		A = maxHeapify(A, i)
	}
	return A
}

func maxHeapify(A []int, i int) []int {
	l := 2 * i
	r := 2 * i + 1

	var largest int
	if l <= len(A) - 1 && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= len(A) - 1 && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		maxHeapify(A, largest)
	}
	return A
}
