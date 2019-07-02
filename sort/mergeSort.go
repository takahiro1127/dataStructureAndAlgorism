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
	// mergeSort(randomArray)
	fmt.Println(divide(randomArray))
}

// func mergeSort(randomArray []int) []int {
// }

// func conquer(rightArray, leftArray []int) []int {

// }

func divide(randomArray []int) ([]int, []int) {
	divideLine := len(randomArray)/2
	leftArray := randomArray[:divideLine]
	rightArray := randomArray[divideLine:]
	return leftArray, rightArray
}
// divideは二つの配列に分けるだけ


