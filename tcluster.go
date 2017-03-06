package main

import (
	"fmt"
	"os"

	"github.com/ntnn/tcluster/cmd"

	"github.com/spf13/cobra/cobra/cmd"
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
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
