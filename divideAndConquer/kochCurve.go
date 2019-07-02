package main

import(
	// "bufio"
	// "fmt"
	// "os"
	"strconv"
	"../procon"
	"github.com/k0kubun/pp"
	"math"
)

type Vector struct {
	x, y float64
}

func kochCurve(headAndTail []Vector) []Vector {
	s, t := headAndTail[0], headAndTail[1]
	vectors := make([]Vector, 5)
	vectors[1].x = (2 * s.x + t.x)/3
	vectors[1].y = (2 * s.y + t.y)/3
	vectors[2].x = (s.x + 2 * t.x)/3
	vectors[2].y = (s.y + 2 * t.y)/3
	vectors[3].x = (t.x - s.x)/2-(t.y - s.y)/math.Pow(3, 1/2)+s.x
	vectors[3].y = (t.x - s.x)/math.Pow(3, 1/2)+(t.y - s.y)/2+s.y
	vectors[0], vectors[4] = s, t
	return vectors
}

func makeKochCurve(n int, vectors []Vector) {
	if n == 0 {
		pp.Print(vectors)
		return
	}
	kochVectors := kochCurve(vectors)
	for i := 0; i < len(kochVectors) - 1; i++ {
		headAndTail := []Vector {kochVectors[i], kochVectors[i+1]}
		makeKochCurve(n - 1, headAndTail)
	}
}

func main() {
	stringInput := procon.StrStdin()
	n, err := strconv.Atoi(stringInput)
	if err != nil {
		panic(err)
	}
	vectors := make([]Vector, 2)
	vectors[0].x = 1
	vectors[0].y = 0
	vectors[1].x = 0
	vectors[1].y = 0
	makeKochCurve(n, vectors)
}
