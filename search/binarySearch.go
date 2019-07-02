package main

import (
	// "bufio"
	// "os"
	"strconv"
	// "strings"
	// "github.com/k0kubun/pp"
	"fmt"
	"../procon"
)

func main() {
	stringInput := procon.StrStdin()
	_, err := strconv.Atoi(stringInput)
	if err != nil {
		panic(err)
	}
	search, err := procon.SplitIntStdin(" ")
	search = procon.InsertionSort(search, len(search))
	if err != nil {
		panic(err)
	}
	stringInput = procon.StrStdin()
	targetLength, err := strconv.Atoi(stringInput)
	if err != nil {
		panic(err)
	}
	target, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	var found []int
	for i := 0; i < targetLength; i++ {
		targetInteger := target[i]
		if binarySearch(search, targetInteger) {
			found = append(found, targetInteger)
		}
	}
	fmt.Println(found)
}

func binarySearch(A []int, n int) bool {
	searchRange := len(A)/2
	i := len(A)/2
	for {
		if A[i] > n {
			i = i - searchRange
		} else if A[i] < n {
			i = i + searchRange
		} else {
			return true
		}
		if (searchRange % 2) == 0 {
			searchRange = searchRange/2
		} else if searchRange != 1 {
			searchRange = searchRange/2 + 1
		} else if searchRange == 1 {
			return false
		}
	}
}
