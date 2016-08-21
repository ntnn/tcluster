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

func TestParseArgs(t *testing.T) {
	cases := map[string]struct {
		args         []string
		expected     config
		expectedErr  error
		expectedExpr []string
	}{
		"None": {
			args: []string{},
			expected: newConfig(config{
				Layout: "5",
			}),
			expectedErr:  nil,
			expectedExpr: []string{},
		},
		"Layout 1 short": {
			args: []string{"l:1"},
			expected: newConfig(config{
				Layout: "1",
			}),
			expectedErr:  nil,
			expectedExpr: []string{},
		},
		"Layout 1 long": {
			args: []string{"layout:1"},
			expected: newConfig(config{
				Layout: "1",
			}),
			expectedErr:  nil,
			expectedExpr: []string{},
		},
		"One expression": {
			args:         []string{"partial-host"},
			expected:     newConfig(config{}),
			expectedErr:  nil,
			expectedExpr: []string{"partial-host"},
		},
		"Layout 1 long and expression": {
			args: []string{"layout:1", "partial-host"},
			expected: newConfig(config{
				Layout: "1",
			}),
			expectedErr:  nil,
			expectedExpr: []string{"partial-host"},
		},
		"Layout 1 short and long and expressions": {
			args: []string{"layout:1", "l:2", "partial-host", "another-host"},
			expected: newConfig(config{
				Layout: "2",
			}),
			expectedErr:  nil,
			expectedExpr: []string{"partial-host", "another-host"},
		},
	}

	for title, test := range cases {
		testconf := newConfig(config{})

		resultExpr, resultErr := testconf.parseArgs(test.args)
		if resultErr != test.expectedErr {
			t.Errorf("%s: Expected error %q, got %q", title, test.expectedErr, resultErr)
		}

		if !reflect.DeepEqual(test.expected, testconf) {
			t.Errorf("%s: Expected result %q, got %q", title, test.expected, testconf)
		}

		// length has to be checked because DeepEqual checks for
		// deep equality of the slice's members - when there are
		// no members the slices aren't deeply equal
		if (len(test.expectedExpr) > 0 || len(resultExpr) > 0) && !reflect.DeepEqual(test.expectedExpr, resultExpr) {
			t.Errorf("%s: Expected expressions %q, got %q", title, test.expectedExpr, resultExpr)
		}
	}
}
