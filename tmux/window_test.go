package tmux

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewWindow(t *testing.T) {
	cases := map[string]struct {
		str    string
		result []string
	}{
		"no title": {
			str:    "",
			result: []string{"new-window"},
		},
	}

	for title, test := range cases {
		cmd := func(in []string) error {
			if !reflect.DeepEqual(in, test.result) {
				return errors.New(fmt.Sprintf("Input not equal to expected: input %q, output %q", in, test.result))
			}
			return nil
		}
		_ = cmd([]string{test.str})

		err := NewWindow(test.str)
		if err != nil {
			t.Errorf("%s: Error: %q", title, err)
		}
	}
}
