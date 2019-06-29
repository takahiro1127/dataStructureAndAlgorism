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
			min, max = input, input
		} else if max < input {
			max = input
		} else if input == -1 {
			if difference < max-min && difference > 0 {
				difference = max - min
			}
			break
		} else if min > input {
			if difference < max-min {
				difference = max - min
			} else if difference < 0 && difference < input - max {
				difference = input - max
			}
			min, max = input, input
		} else if difference < 0 &&  difference < input - max {
			difference = input - max
		}
		fmt.Println("now difference is %s", difference)
	}
	fmt.Println(difference)
}
