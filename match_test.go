package main

import "testing"

func TestExpressionMatchesHost(t *testing.T) {
	cases := map[string]struct {
		expr, host string
		expect     bool
	}{
		"x matches x": {
			expr:   "this-is-a-host",
			host:   "this-is-a-host",
			expect: true,
		},
		"x matches xyz": {
			expr:   "this-is",
			host:   "this-is-a-host",
			expect: true,
		},
		"y does not match x": {
			expr:   "this-is-a-host",
			host:   "this-is-not-a-host",
			expect: false,
		},
		"x matches x.domain": {
			expr:   "some-host",
			host:   "some-host.domain",
			expect: true,
		},
	}

	for title, test := range cases {
		result := expressionMatchesHost(test.expr, test.host)
		if result != test.expect {
			t.Errorf("%s: %q does not match %q", title, result, test.expect)
		}
	}
}

func TestHostsMatchingExpression(t *testing.T) {
	cases := map[string]struct {
		expr          string
		hosts, expect []string
	}{
		"all match": {
			expr: "common",
			hosts: []string{
				"node-common-01",
				"node-uncommon-02.domain",
				"another-common-node",
				"common-node-03",
				"node-not-so-common",
				"node-not-so-common.domain",
			},
		},
	}

	// for title, test := range cases {
	// 	result :=
	// }
}
