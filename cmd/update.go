/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finkit/internal/bootstrap"
	"finkit/internal/version"
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for available updates",
	Long: `Check if a newer version of FinKit is available.

This command queries the latest release from the GitHub repository
and compares it with your current version. If an update is available,
it displays the version information and download URL.`,
	Example: `finkit update`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := cmd.Context().Value("app").(*bootstrap.App)
		check, err := app.Update.Check(cmd.Context())
		if err != nil {
			return err
		}
		fmt.Println("A new version is available!")
		fmt.Printf("Current: %s\n", version.Version)
		fmt.Printf("Latest: %s\n", check.TagName)
		fmt.Println(check.HtmlUrl)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
