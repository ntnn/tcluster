package main

import (
	"log"
	"regexp"
)

// expressionMatchesHost checks if a given expression matches a given
// hostname
func expressionMatchesHost(expr, host string) bool {
	if expr[0] == '-' {
		expr = expr[1:]
	}

	isMatch, err := regexp.MatchString(expr, host)
	if err != nil {
		log.Printf("Error matching %q to %q: %v", expr, host, err)
	}

	return isMatch
}

// expressionMatchesHosts returns hosts matching the expression as
// a slice
func hostsMatchingExpression(hosts []string, expr string) []string {
	matches := []string{}

	for _, host := range hosts {
		if expressionMatchesHost(expr, host) {
			matches = append(matches, host)
		}
	}

	return matches
}

// matchHosts returns hosts matching a slice of expressions as a slice
func matchHosts(expressions []string, hosts []string) []string {
	matches := []string{}

	for _, expr := range expressions {
		matches = append(
			matches,
			hostsMatchingExpression(hosts, expr)...,
		)
	}

	return matches
}
