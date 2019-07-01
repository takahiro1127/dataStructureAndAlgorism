package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var sc = bufio.NewScanner(os.Stdin)
func main() {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	i, err := SplitIntStdin(" ")
	if err != nil {
			fmt.Println(err)
	}
	stableSort(i, n)
}

func isStable(in, out []int) bool {
	n := len(in)
	for i := 0; i < n; i ++ {
		for j := i + 1; j < n; i++ {
			for a := 0; a < n; i++ {
				for b := a + 1; b < n; b++ {
					if in[i] == in[j] && in[i] == out[b] && in[j] == out[a] {
						return true
					}
				}
			}
		}
	}
}
