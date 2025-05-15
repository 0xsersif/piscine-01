package piscine

func ToLower(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	for i, v := range r {
		if v >= 'A' && v <= 'Z' {
			r[i] = v + 32
		}
	}
	return string(r)
}
