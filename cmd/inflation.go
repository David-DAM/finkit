/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finkit/internal/bootstrap"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inflationAmount float64
	inflationYears  int
	inflationRate   float64
)

// inflationCmd represents the inflation command
var inflationCmd = &cobra.Command{
	Use:   "inflation",
	Short: "Calculate the future value of an amount considering inflation",
	Long: `Calculate the future value of an amount after applying inflation over a specified number of years.

This command helps you estimate how much an amount will be worth in the future 
considering the effects of inflation at a given annual rate.`,
	Example: `
finkit inflation --initial 10000 --years 10 --rate 3
finkit inflation --initial 5000 --years 5 --rate 2.5
finkit inflation --initial 50000 --years 20 --rate 4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app := cmd.Context().Value("app").(*bootstrap.App)

		purchasingPower := app.Inflation.Do(inflationAmount, inflationYears, inflationRate)

		moneyLost := inflationAmount - purchasingPower
		fmt.Printf("%.2f today\n\n", inflationAmount)
		fmt.Printf("%.2f purchasing power in %d years (%.2f%% annual inflation)\n", purchasingPower, inflationYears, inflationRate)
		fmt.Printf("%.2f money lost\n", moneyLost)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inflationCmd)

	inflationCmd.Flags().Float64Var(&inflationAmount, "initial", 0, "Initial investment amount")
	inflationCmd.Flags().IntVar(&inflationYears, "years", 0, "Number of years for investment")
	inflationCmd.Flags().Float64Var(&inflationRate, "rate", 0, "Annual interest rate (as percentage)")
}
