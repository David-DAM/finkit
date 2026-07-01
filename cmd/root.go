/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"finkit/internal/bootstrap"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finkit",
	Short: "FinKit is a powerful and lightweight Command Line Interface (CLI) tool built in Go for financial operations",
	Long:  "FinKit is a powerful and lightweight Command Line Interface (CLI) tool built in Go for financial operations",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		app := bootstrap.BuildApp(verbose)

		ctx := context.WithValue(cmd.Context(), "app", app)
		cmd.SetContext(ctx)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var verbose bool

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logs")
}
