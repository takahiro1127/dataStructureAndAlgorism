package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		return -1
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var max int
	var min int
	difference := 0
	for i := 0; i < int(math.Pow(10, 9)); i++ {
		input := nextInt()
		if i == 0 {
			min = input
		}
		if max < input {
			max = input
		} else if input == -1 {
			if difference < max-min {
				difference = max - min
			}
			break
		} else if min > input {
			if difference < max-min {
				difference = max - min
			}
			min, max = input, 0
		}
	}
	fmt.Println(difference)
}
