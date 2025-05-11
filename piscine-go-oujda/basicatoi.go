package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, digit := range s {

		numericValue := int(digit - '0')

		result = result*10 + numericValue
	}

	return result
}

// func main() {
// 	fmt.Println(BasicAtoi("12345"))
// 	fmt.Println(BasicAtoi("0000000012345"))
// 	fmt.Println(BasicAtoi("000000"))
// }
