package piscine

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := []rune(s)
	if result[0] >= 'a' && result[0] <= 'z' {
		result[0] -= 32
	}
	for i := 1; i < len(result); i++ {
		if (result[i-1] < 'a' || result[i-1] > 'z') && (result[i-1] < 'A' || result[i-1] > 'Z') {
			if result[i] >= 'a' && result[i] <= 'z' {
				result[i] -= 32
			}
		} else if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 32
		}
	}
	return string(result)
}
