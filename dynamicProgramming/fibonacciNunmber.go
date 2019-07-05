package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "github.com/k0kubun/pp"
	"../procon"
)
var fibonacciNumberStock []int

func main() {
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}
	fibonacciNumberStock = append(fibonacciNumberStock, 1)
	fibonacciNumberStock = append(fibonacciNumberStock, 1)
	fmt.Println(fibonacciNumber(n))
}

func fibonacciNumber(n int) int {
	if len(fibonacciNumberStock) >= n {
		return fibonacciNumberStock[n - 1]
	} else if n <= 0 {
		return 0
	} else {
		fibonacciNumberStock = append(fibonacciNumberStock, fibonacciNumber(n - 1) + fibonacciNumber(n - 2))
		return fibonacciNumberStock[n - 1]
	}
}
