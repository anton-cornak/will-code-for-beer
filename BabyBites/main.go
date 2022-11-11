package babybites

import "strings"

func babyBites(nr int, words string) string {
	// split by space
	arr := strings.Fields(words)

	if len(arr) != nr {
		return "something is fishy"
	}

	// loop through
	// for _, nr := range arr {

	// }

	return ""
}
