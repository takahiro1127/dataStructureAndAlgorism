package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	// "github.com/k0kubun/pp"
	"../procon"
)

func main() {
	friendsInfo, err := procon.SplitIntStdin(" ")
	if err != nil {
		panic(err)
	}
	numberOfPeople := friendsInfo[0]
	friendLists := make([][]int, numberOfPeople)
	for i := 0; i < numberOfPeople; i++ {
		friendList := make([]int, numberOfPeople)
		friendLists[i] = friendList
	}
	for i := 0; i < friendsInfo[1]; i++ {
		relationship, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		friendLists[relationship[0]][relationship[1]], friendLists[relationship[1]][relationship[0]] = 1, 1
	}
	n, err := strconv.Atoi(procon.StrStdin())
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		relationship, err := procon.SplitIntStdin(" ")
		if err != nil {
			panic(err)
		}
		fmt.Println(connectedFriends(friendLists, relationship[0], relationship[1]))
	}
}

func connectedFriends(friendLists [][]int, start int, goal int) bool {
	find := false
	if friendLists[start][goal] == 1 {
		find = true
	} else {
		for i := 0; i < len(friendLists); i++ {
			friendLists[i][start] = 0
			if friendLists[start][i] == 1 {
				find = find || connectedFriends(friendLists, i, goal)
			}
		}
	}
	return find
}
