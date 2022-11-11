package babybites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphabetSpam(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		x   int
		y   []any
		res string
	}

	testCases := []testCase{
		{x: 5, y: []any{1, 2, 3, "mumble", 5}, res: "makes sense"},
		{x: 8, y: []any{1, 2, 3, "mumble", "mumble", 7, "mumble", 8}, res: "something is fishy"},
		{x: 3, y: []any{"mumble", "mumble", "mumble"}, res: "makes sense"},
	}

	for _, test := range testCases {
		assert.Equal(
			test.res,
			babyBites(test.x, test.y),
		)
	}
}
