package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"../procon"
)

var sc = bufio.NewScanner(os.Stdin)
func main() {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	fmt.Println(n)
	i, err := procon.SplitIntStdin(" ")
	if err != nil {
			fmt.Println(err)
	}
	selectionSort(i, len(i))
}

func selectionSort(A []int, length int){
	for i := 0; i < length; i ++ {
		minj := i
		for j := i; j < length; j++ {
			if A[j] < A[minj] {
			minj = j
			}
		}
		A[i], A[minj] = A[minj], A[i]
	}
	fmt.Println(A)
}
