package procon

func Exponentiate(base, index int) int {
	power := 1
	for i := 0; i < index; i++ {
		power = power * base
	}
	return power
}
