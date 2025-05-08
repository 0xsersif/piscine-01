package piscine

func StrLen(str string) int {
	if str == "" {
		return 0
	}
	count := 0
	for range str {
		count++
	}
	return count
}
