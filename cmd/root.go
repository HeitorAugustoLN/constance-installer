/*
Copyright © 2023 Heitor Augusto <heitoraln@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "constance",
	Short: "Cross-platform declarative package installer",
	Long: `Constance is a cross-platform declarative package installer.
It is a tool that allows you to install packages in a declarative way,
so you can install the same packages in different machines with the same
configuration file.
	
It works by reading a configuration file and installing the packages using the installed package manager in the system.
	
For more information, visit: github.com/HeitorAugustoLN/constance-installer/wiki`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.constance-installer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
