package piscine

func TrimAtoi(s string) int {
	x := []rune(s)
	index := []int{}
	t := 1
	var n int
	for i, char := range s {
		if char >= '0' && char <= '9' {
			index = append(index, i)
		}
		if len(index) == 0 && char == '-' {
			t = -1
		}
	}
	if len(index) == 0 {
		return 0
	}
	for i := 0; i < len(index); i++ {
		n = n*10 + int(x[index[i]]-'0')
	}
	return n * t
}
