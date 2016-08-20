package main

/*
 Application to start multiple ssh or alike connections to servers in
 a separate tmux window.

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
	"os/exec"
)

var tmuxLayouts = []string{
	"even-horizontal",
	"even-vertical",
	"main-horizontal",
	"main-vertical",
	"tiled",
}

func cmd(cmds []string) {
	c := exec.Command("tmux", cmds...)
	c.Run()
}

func window(s string) {
	cmds := []string{"new-window"}
	if s != "" {
		cmds = append(cmds, "-n", s)
	}
	cmd(cmds)
}

func ssh(host string) {
	log.Debug("Opening connection to", host)
	cmd([]string{
		"send-keys",
		"ssh ",
		host,
		"\n",
	})
}

func split() {
	cmd([]string{"split-window"})
	cmd([]string{"select-layout", tmuxLayouts[conf.Layout-1]})
}
