package main

import (
	"strings"

	"github.com/ntnn/tcluster/tmux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ssh(host string) error {
	return tmux.SendKeys(
		strings.Join(
			[]string{"ssh", host},
			" ",
		),
	)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "tcluster ssh opens connections to other nodes",
	Run: func(cmd *cobra.Command, args []string) {
		// Grab title and create new window
		title := viper.Get("Title")
		tmux.NewWindow(title)

		// Match args against hosts and open connections

	},
}

func init() {
	RootCmd.AddCommand(sshCmd)
}
