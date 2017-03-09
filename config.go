package main

import (
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
		log.Fatalf("Error reading config file: %q", err)
	}
}
