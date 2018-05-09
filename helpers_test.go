package cropper

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	table := []struct {
		Purpose       string
		URL           string
		Width, Height int
		Exp           string
		Err           error
	}{
		{"1.OK", "https://some.com", 500, 700,
			`[chrome --headless --disable-gpu --hide-scrollbars --screenshot --window-size=500,700 https://some.com]`,
			nil,
		},
		{"2.Fail: zero width", "https://some.com", 0, 700,
			`[]`,
			errors.New("width must be greater than zero"),
		},
		{"3.Fail: zero height", "https://some.com", 500, 0,
			`[]`,
			errors.New("height must be greater than zero"),
		},
		{"4.Fail: invalid url", "//1m", 500, 700,
			`[]`,
			errors.New("invalid scheme"),
		},
	}
	for _, x := range table {
		args, err := screenshotArgs("chrome", x.URL, x.Width, x.Height)
		assert.EqualValues(t, x.Err, err, x.Purpose)

		actual := fmt.Sprintf("%v", args)
		assert.EqualValues(t, x.Exp, actual, x.Purpose, actual)
	}
}
