package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "viper",
	Short: "Configuration file for the server",
	
	Run: func(cmd *cobra.Command, args []string) {
		// Empty for running the initConfig
	},
}

var cfgFile = ""

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Configuration file name")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Name of the super admin")

	viper.BindPFlag("super_admin_name", rootCmd.PersistentFlags().Lookup("username"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)			// Name of the config path
	} else {
		viper.SetConfigType("yaml")				// Type of the configuration file
		viper.SetConfigName("config")			// Configuration file name
		viper.AddConfigPath(".")				// Look for the file in the current directory
		viper.AddConfigPath("$HOME/.appname")	// If file not find in current directory then check the $HOME/.appname directory
	}

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Errorf("Error in cobra rootCmd execution: %u\n", err))
	}
}


