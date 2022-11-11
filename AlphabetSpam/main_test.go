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
		{x: "Welcome_NWERC_participants!", res: []float64{0.07407407407407407, 0.6666666666666666, 0.2222222222222222, 0.037037037037037035}},
		{x: `\/\/in_US$100000_in_our_Ca$h_Lo||ery!!!`, res: []float64{0.1282051282051282, 0.3333333333333333, 0.10256410256410256, 0.4358974358974359}},
	}

	for _, test := range testCases {
		assert.Equal(
			test.res,
			AlphabetSpam(test.x),
		)
	}
}
