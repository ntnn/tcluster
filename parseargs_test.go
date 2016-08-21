package main

/*
The MIT License (MIT)
Copyright (c) 2016 Nelo-T. wallus <nelo@wallus.de>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"reflect"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	cases := map[string]struct {
		input    []string
		expected [][]string
	}{
		"One then": {
			input: []string{"first", "then", "second"},
			expected: [][]string{
				[]string{"first"},
				[]string{"second"},
			},
		},
		"Successive thens": {
			input: []string{"first", "then", "then", "second"},
			expected: [][]string{
				[]string{"first"},
				[]string{"second"},
			},
		},
		"Thens at the start": {
			input: []string{"then", "then", "first", "second"},
			expected: [][]string{
				[]string{"first", "second"},
			},
		},
		"Thens at the end": {
			input: []string{"first", "second", "then", "then"},
			expected: [][]string{
				[]string{"first", "second"},
			},
		},
		"No then": {
			input: []string{"first", "second"},
			expected: [][]string{
				[]string{"first", "second"},
			},
		},
	}

	for title, test := range cases {
		result := splitArgs(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("%s: Expected %q, got %q", title, test.expected, result)
		}
	}
}
