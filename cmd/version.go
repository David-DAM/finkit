/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finkit/internal/version"
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of FinKit",
	Long: `Display the current version information of FinKit CLI.

This command shows the version number of the FinKit application you are currently running.`,
	Example: `finkit version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
