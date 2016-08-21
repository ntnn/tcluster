package main

import "strings"

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

// findXInArray returns the positions of word in splice.
func findXInArray(word string, splice []string) []int {
	positions := []int{}

	for i := range splice {
		if word == splice[i] {
			positions = append(positions, i)
		}
	}

	return positions
}

// splitArgs splits the passed splice on the keyword "then" and returns
// a 2d splice containing  in between
func splitArgs(args []string) [][]string {
	thens := findXInArray("then", args)
	if len(thens) == 0 {
		return [][]string{args}
	}

	splits := [][]string{}

	left := 0
	for _, right := range thens {
		temp := args[left:right]

		if len(findXInArray("then", temp)) != len(temp) {
			splits = append(splits, temp)
		}

		left = right
	}

	// grab the last part of the passed arguments, but omit if
	// 'then' is the last entry
	lastthen := thens[len(thens)-1] + 1
	if lastthen < len(args) {
		splits = append(splits, args[lastthen:])
	}

	return splits
}

// parseArgs takes a list of strings (a block of expressions) and parses
// config modifiers into the configuration
// Everything with a colon in it is considered to be a config modifier,
// unrecognized modifiers are logged, but the method continues.
// Everything not considered a config modifier is considered an
// expression and returned.
// TODO title modifier to set window name
func (c *config) parseArgs(args []string) ([]string, error) {
	var expressions []string

	for i := range args {
		split := strings.Split(args[i], ":")
		if len(split) > 1 {
			flag, arg := split[0], split[1]

			switch flag {
			case "l", "layout":
				c.Layout = arg
			default:
				log.Errorf("Parsed unknown config modified %q with value %q", flag, arg)
			}
		} else {
			expressions = append(expressions, args[i])
		}
	}

	return expressions, nil
}
