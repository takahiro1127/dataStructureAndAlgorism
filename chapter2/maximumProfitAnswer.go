package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
var sc = bufio.NewScanner(os.Stdin)
func main() {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	var R [200000]int
	for i := 0; i < n; i++ {
		sc.Scan()
		R[i], e = strconv.Atoi(sc.Text())
		if e != nil {
		panic(e)
		}
	}

  minv := R[0]

	for i := 1; i < n; i++ {
		maxv = int(math.Max(float64(maxv), float64(R[i] - minv)))
		minv = int(math.Min(float64(minv), float64(R[i])))
	}

	fmt.Println("%s", maxv)
}
