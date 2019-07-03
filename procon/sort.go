package procon

func IsStable(in, out []int) bool {
	n := len(in)
	stability := false
	for i := 0; i < n; i ++ {
		for j := i + 1; j < n; i++ {
			for a := 0; a < n; i++ {
				for b := a + 1; b < n; b++ {
					if in[i] == in[j] && in[i] == out[b] && in[j] == out[a] {
						stability = true
					}
				}
			}
		}
	}
	return stability
}

func SelectionSort(A []int, length int){
	for i := 0; i < length; i ++ {
		minj := i
		for j := i; j < length; j++ {
			if A[j] < A[minj] {
			minj = j
			}
		}
		A[i], A[minj] = A[minj], A[i]
	}
}

func InsertionSort(randomArray []int, n int) []int {
	for i, j := 1, 0; i < n; i, j  = i + 1, j + 1{
		v := randomArray[i]
		k := j
		for ;k >= 0 && randomArray[k] > v; k = k - 1 {
			randomArray[k+1] = randomArray[k]
		}
		randomArray[k+1] = v
	}
	return randomArray
}

func BubbleSort(randomArray []int, n int) []int {
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
	return randomArray
}

func MergeSort(randomArray []int) []int {
	leftDivided, rightDivided := divide(randomArray)
	if len(leftDivided) > 1 {
		leftDivided = MergeSort(leftDivided)
	}
	if len(rightDivided) > 1 {
		rightDivided = MergeSort(rightDivided)
	}
	return conquer(leftDivided, rightDivided)
}

func conquer(leftArray, rightArray []int) []int {
	var sortedArray []int
	for i, j := 0, 0; ; {
		min := Min(leftArray[i], rightArray[j])
		left := (min == leftArray[i])
		if len(leftArray) == i + 1 && left {
			plusArray := rightArray[i:]
			sortedArray  = append(append(sortedArray, min), plusArray...)
			break
		} else if len(rightArray) == j + 1 && !left {
			plusArray := leftArray[i:]
			sortedArray  = append(append(sortedArray, min), plusArray...)
			break
		}
		sortedArray  = append(sortedArray, min)
		if left {
			i++
		} else {
			j++
		}
	}
	return sortedArray
}

func divide(randomArray []int) ([]int, []int) {
	divideLine := len(randomArray)/2
	leftArray := randomArray[:divideLine]
	rightArray := randomArray[divideLine:]
	return leftArray, rightArray
}
