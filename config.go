package main

import (
	"regexp"
	"sort"

	"github.com/spf13/viper"

	"log"
)

func config() {
	viper.SetDefault("Layout", "5")
	viper.SetDefault("Layouts", map[string]string{
		"1": "even-horizontal",
		"2": "even-vertical",
		"3": "main-horizontal",
		"4": "main-vertical",
		"5": "tiled",
	})
	viper.SetDefault("Hosts", []string{})
	viper.SetDefault("Title", "")

	viper.SetConfigName("tcluster")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$HOME/.tcluster")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fatalf("Error reading config file: %q", err)
	}
}

func matchHosts(expressions []string) []string {
	hosts := map[string]bool{}
	cfgHosts := viper.Get("Hosts")

	// matching input against defined hosts
	for _, host := range cfgHosts {
		log.Debug("Checking host", host)

		for _, expr := range expressions {
			log.Debug("\tagainst expression", expr)

			// checking for a leading minus at the beginning
			if expr[0] == '-' {
				isMatch, _ := regexp.MatchString(expr[1:], host)
				if isMatch {
					log.Debugf("Hosts '%s' matched against '%s', removing from host list", host, expr)
					hosts[host] = false
				}
				continue
			}

			isMatch, _ := regexp.MatchString(expr, host)
			if isMatch {
				log.Debug("\t\tmatched")
				hosts[host] = true
			}

		}
	}

	// transfer keys from created map to list
	list := make([]string, 0)
	for key, value := range hosts {
		if value {
			list = append(list, key)
		}
	}

	// sort it
	sort.Strings(list)

	return list
}
