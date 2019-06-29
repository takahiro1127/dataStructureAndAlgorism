package procon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(randomArray []int, n int) []int {
	for i, j := 1, 0; i < n; i, j  = i + 1, j + 1{
		v := randomArray[i]
		k := j
		for ;k >= 0 && randomArray[k] > v; k = k - 1 {
			randomArray[k+1] = randomArray[k]
		}
		randomArray[k+1] = v
		fmt.Println(randomArray)
	}
	return randomArray
}

func SplitIntStdin(delim string) (intSlices []int, err error) {
    // 文字列スライスを取得
    stringSplited := SplitStrStdin(" ")

    // 整数スライスに保存
    for i := range stringSplited {
        var iparam int
        iparam, err = strconv.Atoi(stringSplited[i])
        if err != nil {
            return
        }
        intSlices = append(intSlices, iparam)
    }
    return
}

func SplitStrStdin(delim string) (stringReturned []string) {
    // 文字列で1行取得
    stringInput := StrStdin()

    // 空白で分割
    // stringSplited := strings.Split(stringInput, delim)
    stringReturned = SplitWithoutEmpty(stringInput, delim)

    return
}

func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
    stringSplited := strings.Split(stringTargeted, delim)

    for _, str := range stringSplited {
        if str != "" {
            stringReturned = append(stringReturned, str)
        }
    }

    return
}

func StrStdin() (stringInput string) {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    stringInput = scanner.Text()

    stringInput = strings.TrimSpace(stringInput)
    return
}

