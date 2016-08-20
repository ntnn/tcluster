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
	"flag"
	"log"
	"os"
	"regexp"
	"sort"
)

var conf = new(config)

func collectHosts() []string {
	hosts := map[string]bool{}

	// matching input against defined hosts
	for _, host := range conf.Hosts {
		log.Println("Checking host", host)
		for _, arg := range os.Args[1:] {
			log.Println("\tagainst pattern", arg)
			isMatch, _ := regexp.MatchString(arg, host)
			if isMatch {
				log.Println("\t\tmatched")
				hosts[host] = true
			}
		}
	}

	// transfer keys from created map to list
	list := make([]string, len(hosts))
	i := 0
	for key, _ := range hosts {
		list[i] = key
		i++
	}

	// sort it
	sort.Strings(list)

	return list
}

func openHosts() {
	hosts := collectHosts()
	if len(hosts) > 0 {
		window("")

		i := 0
		for j := range hosts {
			ssh(hosts[j])
			if i < len(hosts)-1 {
				split()
			}
			i++
		}
	}
}

func main() {
	var confpath string
	flag.StringVar(&confpath, "config", "", "Specify configuration file")
	flag.Parse()

	// define err here to prevent redefinition of confpath in the
	// inner block by using :=
	var err error
	if confpath == "" {
		confpath, err = confPath()
		if err != nil {
			log.Panic(err)
		}
	}

	err = conf.parseFile(confpath)
	if err != nil {
		log.Panic(err)
	}

	openHosts()
}
