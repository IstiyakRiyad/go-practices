package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func readFromEnv() {
	viper.SetConfigType("env")				// Type of the configuration file
	viper.SetConfigName("prod")				// Configuration file name
	viper.AddConfigPath(".")				// Look for the file in the current directory
	viper.AddConfigPath("$HOME/.appname")	// If file not find in current directory then check the $HOME/.appname directory

	viper.AutomaticEnv()		// viper will check for an environment variable any time a viper.Get() request is made

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file is not found")
			os.Exit(1)
		} else {
			panic(fmt.Errorf("fetal error config file: %w", err))
		}
	}
}

func readFromYaml() {
	viper.SetConfigType("yaml")				// Type of the configuration file
	viper.SetConfigName("config")			// Configuration file name
	viper.AddConfigPath(".")				// Look for the file in the current directory
	viper.AddConfigPath("$HOME/.appname")	// If file not find in current directory then check the $HOME/.appname directory

	viper.AutomaticEnv()		// viper will check for an environment variable any time a viper.Get() request is made

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file is not found")
			os.Exit(1)
		} else {
			panic(fmt.Errorf("fetal error config file: %w", err))
		}
	}
}


func main() {
	// readFromEnv()
	// readFromYaml()
	Execute()

	fmt.Println(viper.Get("SUPER_ADMIN_NAME"))
	fmt.Println(viper.Get("SUPER_ADMIN_PASSWORD"))

	fmt.Println(viper.Get("path"))
}





