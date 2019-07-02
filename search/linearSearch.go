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
		if linearSearch(search, targetInteger) {
			found = append(found, targetInteger)
		}
	}
	fmt.Println(len(found))
}

func linearSearch(A []int, n int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] == n {
			return true
		}
	}
	return false
}
