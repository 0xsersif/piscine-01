package piscine

func ToUpper(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	for i, v := range r {
		if v >= 'a' && v <= 'z' {
			r[i] = v - 32
		}
	}
	return string(r)
}
