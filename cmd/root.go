package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "tcluster",
	Short: "tcluster is clusterssh for tmux (and more)",
	Long:  "clusterssh replacement with additions to act on multiple servers at once",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config file (defaults to $HOME/.tcluster.yaml)")
}
