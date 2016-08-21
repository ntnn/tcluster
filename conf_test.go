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
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cases := map[string]struct {
		input, expected config
	}{
		"No modification": {
			input: config{},
			expected: config{
				Layout: "5",
				Layouts: map[string]string{
					"1": "even-horizontal",
					"2": "even-vertical",
					"3": "main-horizontal",
					"4": "main-vertical",
					"5": "tiled",
				},
				Hosts: make([]string, 0),
			},
		},
		"Layout changed": {
			input: config{
				Layout: "2",
			},
			expected: config{
				Layout: "2",
				Layouts: map[string]string{
					"1": "even-horizontal",
					"2": "even-vertical",
					"3": "main-horizontal",
					"4": "main-vertical",
					"5": "tiled",
				},
				Hosts: make([]string, 0),
			},
		},
	}

	for title, test := range cases {
		result := newConfig(test.input)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("%s: Expected %q, got %q", title, test.expected, result)
		}
	}
}

func TestParseFile(t *testing.T) {
	cases := map[string]struct {
		path        string
		expected    config
		expectedErr error
	}{
		"Simple": {
			path: "test_data/simple.yaml",
			expected: newConfig(config{
				Layout: "1",
				Hosts:  []string{"first.fully.qualified.host", "second.fully.qualified.host"},
			}),
			expectedErr: nil,
		},
		"One in dir": {
			path: "test_data/dir_single",
			expected: newConfig(config{
				Layout: "1",
				Hosts: []string{
					"first.fully.qualified.host",
					"second.fully.qualified.host",
					"host-without-domain",
				},
			}),
			expectedErr: nil,
		},
		"Multiple in dir": {
			path: "test_data/dir_multiple",
			expected: newConfig(config{
				Layout: "1",
				Hosts: []string{
					"first.fully.qualified.host",
					"host-without-domain",
				},
			}),
			expectedErr: nil,
		},
		"Defined layouts, overwriting default": {
			path: "test_data/layouts.yaml",
			expected: newConfig(config{
				Layouts: map[string]string{
					"one": "layout-string",
					"2":   "tiled",
				},
			}),
			expectedErr: nil,
		},
	}

	for title, test := range cases {
		result := newConfig(config{})
		err := result.parseFile(test.path)
		if err != test.expectedErr {
			t.Errorf("%s: Expected error '%v', got '%v'", title, test.expectedErr, err)
		}

		if !reflect.DeepEqual(test.expected, result) {
			t.Error(title, ": Expexted result:", test.expected, "got:", result)
		}
	}
}

type confPathCase struct {
	env         string
	dirs        []string
	files       []string
	expected    string
	extectedErr error
}

func TestConfPath(t *testing.T) {
	tempdir := filepath.Join(os.TempDir(), "tcluster_test")
	temphome := filepath.Join(tempdir, "home")
	os.Setenv("HOME", temphome)

	err := os.MkdirAll(temphome, os.ModeTemporary|0700)
	if err != nil {
		t.Errorf("Error creating temporary test dir: %v", err)
	}

	cases := map[string]confPathCase{
		"Environment": {
			filepath.Join(tempdir, "env_file"),
			[]string{},
			[]string{},
			filepath.Join(tempdir, "env_file"),
			nil,
		},
		"Directory": {
			"",
			[]string{filepath.Join(temphome, ".tcluster.d")},
			[]string{filepath.Join(temphome, ".tcluster.yaml")},
			filepath.Join(temphome, ".tcluster.d"),
			nil,
		},
		"File": {
			"",
			[]string{},
			[]string{filepath.Join(temphome, ".tcluster.yaml")},
			filepath.Join(temphome, ".tcluster.yaml"),
			nil,
		},
		"Error": {
			"",
			[]string{},
			[]string{},
			"",
			ErrNoValidConfigPath,
		},
	}

	for title, test := range cases {
		testConfPathHelper(t, title, test)
	}
}

func testConfPathHelper(t *testing.T, title string, test confPathCase) {
	if test.env != "" {
		os.Setenv("TCLUSTER_CONF", test.env)
		defer os.Setenv("TCLUSTER_CONF", "")

		err := os.MkdirAll(test.env, os.ModeDir)
		defer os.RemoveAll(test.env)

		if err != nil {
			t.Errorf("%s: Error preparing test: %v", title, err)
		}
	}

	for i := range test.dirs {
		os.Mkdir(test.dirs[i], os.ModeTemporary|0700)
		defer os.RemoveAll(test.dirs[i])
	}

	for i := range test.files {
		os.Create(test.files[i])
		defer os.Remove(test.files[i])
	}

	result, resultErr := confPath()
	if result != test.expected {
		t.Errorf("%s: Expected result '%s' does not match result '%s'", title, test.expected, result)
	}

	if resultErr != test.extectedErr {
		t.Errorf("%s: Expected error '%v' does not match error '%s'", title, test.extectedErr, resultErr)
	}
}
