/*
Copyright © 2022 Aditya Agrawal adiag1200@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "drive-encrypt",
	Short: "Tool for file encryption",
	Long: `Tool for file encryption
Can be used using command line or using a web GUI (powered by a http server).`,
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")
}
