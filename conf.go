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
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	// ErrNoValidConfigPath is returned by confPath when none of the
	// default locations are acessible.
	ErrNoValidConfigPath = errors.New("No configuration path found.")
)

type config struct {
	Layout  string
	Layouts map[string]string
	Hosts   []string
}

func newConfig(init config) config {
	cfg := config{
		Layout: "5",
		Layouts: map[string]string{
			"1": "even-horizontal",
			"2": "even-vertical",
			"3": "main-horizontal",
			"4": "main-vertical",
			"5": "tiled",
		},
		Hosts: []string{},
	}

	if init.Layout != "" {
		cfg.Layout = init.Layout
	}

	// TODO ugly
	for i := range init.Hosts {
		cfg.Hosts = append(cfg.Hosts, init.Hosts[i])
	}

	for key := range init.Layouts {
		cfg.Layouts[key] = init.Layouts[key]
	}

	return cfg
}

// Parse given filepath into config struct. If the path is a directory
// all yaml files within will be recursively parsed.
func (c *config) parseFile(path string) error {
	log.Debugf("Parsing from path '%s'", path)

	fileinfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileinfo.IsDir() {
		log.Debugf("'%s' is a directory, parsing recursively", path)

		dir, err := os.Open(path)
		if err != nil {
			return err
		}

		files, err := dir.Readdir(0)
		if err != nil {
			return err
		}

		for i := range files {
			err = c.parseFile(path + "/" + files[i].Name())
			if err != nil {
				return err
			}
		}

		return nil
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// TODO expand host with globs or similar
	//   host{01..05}
	// to
	//   host01
	//   host02
	//   host03
	//   host04
	//   host05
	return yaml.Unmarshal(file, c)
}

// Retrieve a valid configuration path:
//	env TCLUSTER_CONF
//	~/.tcluster.d
//	~/.tcluster.yml
//	~/.tcluster.yaml
// The first valid path is returned.
func confPath() (string, error) {
	// TODO: allow multiple colon-delimited directories
	path := os.ExpandEnv(os.Getenv("TCLUSTER_CONF"))
	if path != "" {
		_, err := os.Stat(path)
		if err == nil {
			log.Debugf("Configuration path: %s", path)
			return path, nil
		}
	}

	paths := []string{
		"$HOME/.tcluster.d",
		"$HOME/.tcluster.yml",
		"$HOME/.tcluster.yaml",
	}

	for i := range paths {
		path := os.ExpandEnv(paths[i])
		log.Debugf("Trying path '%s'", path)

		if path != "" {
			_, err := os.Stat(path)
			if err == nil {
				log.Debugf("Configuration path: %s", path)
				return path, nil
			}
		}
	}

	log.Debug("No configuration file found")
	return "", ErrNoValidConfigPath
}
