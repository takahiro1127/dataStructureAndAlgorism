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
	var A []int
	for {
		order := procon.SplitStrStdin(" ")
		method := order[0]
		if len(order) == 2{
			target, err := strconv.Atoi(order[1])
			if err != nil {
				panic(err)
			}
		}
		if method == "insert" {
			A = append(A, target)
			A = buildMaxHeap(A)
		} else if method == "extract" {
			A = extract(A)
		} else {
			break
		}
	}
}

func insert(A []int, i int) []int {
	A = append(A, i)
	for i := len(A) - 1, i >= 1; {
		if A[i] > A[(i - 1) / 2] {
			break
		} else {
			A[i], A[(i - 1) / 2] = A[(i - 1) / 2], A[i]
			i = (i - 1) / 2
		}
	}
}

func extract(A) []int{
	fmt.Println(A[0])
	B := A[0:]
	return buildMaxHeap(B)
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

