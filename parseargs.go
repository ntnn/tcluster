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
// a 2d splice containing  inbetween
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
