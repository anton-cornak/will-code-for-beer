package alphabetspam

import (
	"fmt"
	"time"
)

func alphabetSpam(text string) (res []float64) {
	start := time.Now()

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

	end := time.Now()
	fmt.Println("elapsed: ", end.Sub(start))

	return res
}
