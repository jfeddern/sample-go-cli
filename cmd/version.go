/*
Copyright © 2022 Jan Feddern
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version information of the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Currently using CLI-Version: %v \n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
