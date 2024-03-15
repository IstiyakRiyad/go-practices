package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "password-gen", // Name of your cli
	Short: "Info about the cli",
	Long: `This is a tool for generating the password.`,

	// Here not give the Run because without the gen key it will show long message. 
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Hello world")
	// },
}

var cfgFile string;

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err);
		os.Exit(1)
	}
}

func init() {
	// This `initConfig` function will be called when command execute will be called.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Configuration file name")
}
func initConfig() {
	// Code for loading config file using viper lib
}





