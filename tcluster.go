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
	"errors"
	"flag"
	"os"

	"github.com/divoxx/llog"
)

var (
	conf = newConfig(config{})
	log  = llog.New(os.Stdout, llog.INFO)
)

var (
	ErrEmptyHostList = errors.New("Empty host list passed")
)

func openHosts(title string, hosts []string) error {
	if len(hosts) == 0 {
		return ErrEmptyHostList
	}

	window(title)

	for i := range hosts {
		ssh(hosts[i])
		if i < len(hosts)-1 {
			split(conf.Layouts[conf.Layout])
		}
	}

	return nil
}

func main() {
	var (
		confpath string
		err      error
	)

	flag.StringVar(&confpath, "config", "", "Specify configuration file")
	debug := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()

	if *debug {
		log.SetLevel(llog.DEBUG)
	}

	if confpath == "" {
		confpath, err = confPath()
		if err != nil {
			log.Errorf("Error getting confpath, got '%s': %v", confpath, err)
		}
	}

	err = conf.parseFile(confpath)
	if err != nil {
		log.Errorf("Failed to parse configuration file(s): %v", err)
	}

	args := splitArgs(flag.Args())
	for i := range args {
		blockconf := newConfig(conf)
		expressions, err := blockconf.parseArgs(args[i])
		if err != nil {
			log.Errorf("Got error parsing arguments for a block, not continuing this block: %v", err)
			continue
		}

		err = openHosts(blockconf.Title, blockconf.matchHosts(expressions))
		if err != nil {
			log.Errorf("Got error opening connections to hosts: %v", err)
		}
	}
}
