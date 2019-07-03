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
	fmt.Println(mergeSort(randomArray))
	// fmt.Println(divide(randomArray))
	// fmt.Println(conquer(divide(randomArray)))
}

func mergeSort(randomArray []int) []int {
	leftDivided, rightDivided := divide(randomArray)
	if len(leftDivided) > 1 {
		leftDivided = mergeSort(leftDivided)
	}
	if len(rightDivided) > 1 {
		rightDivided = mergeSort(rightDivided)
	}
	return conquer(leftDivided, rightDivided)
}

func conquer(leftArray, rightArray []int) []int {
	var sortedArray []int
	for i, j := 0, 0; ; {
		min := procon.Min(leftArray[i], rightArray[j])
		left := (min == leftArray[i])
		if len(leftArray) == i + 1 && left {
			plusArray := rightArray[i:]
			sortedArray  = append(append(sortedArray, min), plusArray...)
			break
		} else if len(rightArray) == j + 1 && !left {
			plusArray := leftArray[i:]
			sortedArray  = append(append(sortedArray, min), plusArray...)
			break
		}
		sortedArray  = append(sortedArray, min)
		if left {
			i++
		} else {
			j++
		}
	}
	return sortedArray
}

func divide(randomArray []int) ([]int, []int) {
	divideLine := len(randomArray)/2
	leftArray := randomArray[:divideLine]
	rightArray := randomArray[divideLine:]
	return leftArray, rightArray
}
// divideは二つの配列に分けるだけ


