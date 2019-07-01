package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	stringSplited := SplitStrStdin(" ")
	fmt.Println(reversePolishNotation(stringSplited))
	slice := []int{0, 1}
	_, slice = pop(slice)
	fmt.Println(slice)
	last ,slice := pop(slice)
	fmt.Println(slice)
	fmt.Println(last)
}

func reversePolishNotation(array []string) int {
	var stack []int
	for _, num := range array {
		if num == "+" {
			last, stack := pop(stack)
			secondlast, stack := pop(stack)
			newlast := last + secondlast
			fmt.Println("今から足すよ")
			fmt.Println(stack)
			fmt.Println(newlast)
			stack[len(stack) - 1] = newlast
		} else if num == "-" {
			last, stack := pop(stack)
			secondlast, stack := pop(stack)
			newlast := last - secondlast
			stack = append(stack, newlast)
		} else if num == "*" {
			last, stack := pop(stack)
			secondlast, stack := pop(stack)
			newlast := last * secondlast
			stack = append(stack, newlast)
		} else {
			integer, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			stack = append(stack, integer)
		}
		fmt.Println(stack)
	}
	result, stack := pop(stack)
	return result
}

func pop(stack []int) (int, []int) {
	if len(stack) == 1 {
		return stack[len(stack)-1], []int{}
	}
	return stack[len(stack)-1], stack[:len(stack)-1]
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
