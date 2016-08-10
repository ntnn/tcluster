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
	"os"
	"errors"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Layout int
	Hosts []string
}

func checkpath(s string) (string, error) {
	_, err := os.Stat(s)
	if err == nil {
		return s, nil
	}
	return "", err
}

func confpath() (string, error) {
	path := os.ExpandEnv(os.Getenv("TCLUSTER_CONF"))
	if path != "" { return checkpath(path) }

	paths := []string {
		os.ExpandEnv("$HOME/.tcluster.yml"),
		os.ExpandEnv("$HOME/.tcluster.yaml"),
	}

	for _, path := range paths {
		if path != "" {
			_, err := os.Stat(path)
			if err == nil { return path, err }
		}
	}

	return "", errors.New("No configuration file found")
}

func (c *config) Parse() *config {
	path, err := confpath()
	if err != nil { log.Panic(err) }
	log.Println(path)

	file, err := ioutil.ReadFile(path)
	if err != nil { log.Panic(err) }

	err = yaml.Unmarshal(file, c)
	if err != nil { log.Panic(err) }

	log.Println(c)
	return c
}
