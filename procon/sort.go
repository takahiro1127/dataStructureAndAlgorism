package procon

func isStable(in, out []int) bool {
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

func selectionSort(A []int, length int){
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

func insertionSort(randomArray []int, n int) []int {
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
	return randomArray
}
