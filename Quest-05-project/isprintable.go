package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if rune(char) < 32 || rune(char) > 127 {
			return false
		}
	}
	return true
}
