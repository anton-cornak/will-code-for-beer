package alphabetspam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphabetSpam(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		x   string
		res []float64
	}

	testCases := []testCase{
		{x: "Welcome_NWERC_participants!", res: []float64{0.0740740740740741, 0.666666666666667, 0.222222222222222, 0.0370370370370370}},
		{x: `\/\/in_US$100000_in_our_Ca$h_Lo||ery!!!`, res: []float64{0.128205128205128, 0.333333333333333, 0.102564102564103, 0.435897435897436}},
	}

	for _, test := range testCases {
		assert.Equal(
			test.res,
			alphabetSpam(test.x),
		)
	}
}
