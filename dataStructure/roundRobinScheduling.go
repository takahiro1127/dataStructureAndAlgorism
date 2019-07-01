package main

import (
	// "bufio"
	// "os"
	"strconv"
	// "strings"
	"github.com/k0kubun/pp"
	"fmt"
	"../procon"
)

type Que struct{
	processes []Process
	processCount int
	remainingProcessCount int
	quantum int
}

type Process struct {
	name string
	time int
}

func main() {
	var que Que
	initialInfo, err := procon.SplitIntStdin(" ")
	if err != nil {
			fmt.Println(err)
	}
	que.processCount, que.quantum = initialInfo[0], initialInfo[1]
	for i := 0; i < que.processCount; i++ {
		stringSplited := procon.SplitStrStdin(" ")
		time, err := strconv.Atoi(stringSplited[1])
		if err != nil {
				return
    }
		var process Process = Process{
			name: stringSplited[0],
			time: time,
		}
		que.processes = append(que.processes, process)
	}
	que.excute()
}

func (que Que)excute() int {
	var excutingTime int
	var resultBoard map[string]int
	que.remainingProcessCount = que.processCount
	for que.remainingProcessCount > 0 {
		for i := 0; i < len(que.processes); i++ {
			if que.processes[i].time == 0 { continue }
			remainingTime := que.processes[i].time - que.quantum
			if remainingTime > 0 {
				que.processes[i].time = remainingTime
				excutingTime += que.quantum
			} else if remainingTime <= 0 {
				que.processes[i].time = 0
				excutingTime += (-1 * remainingTime)
				pp.Print(que.processes[i].name)
				pp.Print(excutingTime)
				resultBoard[que.processes[i].name] = excutingTime
				pp.Print("ここまできてる")
				pp.Print(resultBoard)
				resultBoard = map[string]int {}
				que.remainingProcessCount--
			}
			fmt.Println(que.processes[i])
		}
	}
	return excutingTime
}
