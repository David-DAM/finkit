/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finkit/internal/bootstrap"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := cmd.Context().Value("app").(*bootstrap.App)

		currencies, err := app.Currency.Currencies()
		if err != nil {
			return err
		}

		fmt.Println("Currency list:")
		fmt.Println()
		fmt.Printf("%-12s %-12s %-40s %-10s %-12s\n", "ISO Code", "ISO Numeric", "Name", "Symbol", "Start Date")
		fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────")
		for _, c := range currencies {
			fmt.Printf("%-12s %-12s %-40s %-10s %-12s\n", c.IsoCode, c.IsoNumeric, c.Name, c.Symbol, c.StartDate)
		}
		return nil
	},
}

func init() {
	currencyCmd.AddCommand(listCmd)
}
