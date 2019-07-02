package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"../procon"
)


func main() {
	stringInput := procon.StrStdin()
	_, err := strconv.Atoi(stringInput)
	if err != nil {
		panic(err)
	}
	parts, err := procon.SplitIntStdin(" ")
	parts = procon.InsertionSort(parts, len(parts))
	if err != nil {
		panic(err)
	}
	stringInput = procon.StrStdin()
	targetCount, err := strconv.Atoi(stringInput)
	if err != nil {
		panic(err)
	}
	targets, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	for i := 0; i < targetCount; i++ {
		fmt.Println(exhausiveSearch(targets[i], parts))
	}
}

func exhausiveSearch(target int, parts []int) bool {
	// fmt.Println(target)
	// fmt.Println(parts)
	for i := 0; i < len(parts); i++ {
		if parts[i] == 0 { continue } else if parts[i] > target { return false }
		remaining := target - parts[i]
		if remaining > 0 {
			parts[i] = 0
			if exhausiveSearch(remaining, parts) { return true }
			parts[i] = target - remaining
		} else if remaining == 0 {
			return true
		}
	}
	return false
}
