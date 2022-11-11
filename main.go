package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	fmt.Println(len(text))
	res := AlphabetSpam(text)
	fmt.Printf("%f\n%f\n%f\n%f\n", res[0], res[1], res[2], res[3])
}
