package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"../procon"
)

var sc = bufio.NewScanner(os.Stdin)
func main() {
	sc.Scan()
	n, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	fmt.Println(n)
	i, err := procon.SplitIntStdin(" ")
	if err != nil {
			fmt.Println(err)
	}
	bubbleSort(i, len(i))
}

func bubbleSort(randomArray []int, n int) []int {
	flag := true
	v := 1
	for ; flag || v == n; v++ {
		flag = false
		for j := n - 1; j >= 1; j = j - 1{
			if randomArray[j] < randomArray[j - 1] {
				randomArray[j], randomArray[j - 1] = randomArray[j - 1], randomArray[j]
				flag = true
			}
		}
	}
	fmt.Println(randomArray)
	fmt.Println(v)
	return randomArray
}

