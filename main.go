package main

import (
	"fmt"
)

func AlphabetSpam(text string) (res []float64) {
	var whitespaces_count, lowercase_count, uppercase_count, symbol_count int

	for _, c := range text {
		if c >= 97 && c <= 122 {
			lowercase_count++
			continue
		}
		if c >= 65 && c <= 90 {
			uppercase_count++
			continue
		}
		if c == 95 {
			whitespaces_count++
			continue
		}
		symbol_count++
	}

	var length = float64(len(text))

	res = append(res, float64(whitespaces_count)/length)
	res = append(res, float64(lowercase_count)/length)
	res = append(res, float64(uppercase_count)/length)
	res = append(res, float64(symbol_count)/length)

	return res
}

func main() {
	var text string
	fmt.Scanln(&text)
	res := AlphabetSpam(text)

	fmt.Println(res[0])
	fmt.Println(res[1])
	fmt.Println(res[2])
	fmt.Println(res[3])
}
