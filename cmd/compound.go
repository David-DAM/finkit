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
	initialInvestment   float64
	monthlyContribution float64
	years               int
	rate                float64
)

// compoundCmd represents the compound command
var compoundCmd = &cobra.Command{
	Use:   "compound",
	Short: "Calculate compound interest for investments",
	Long: `Calculate compound interest for an investment with initial capital, 
monthly contributions, interest rate, and time period.

This command helps you estimate the final amount of an investment 
considering compound interest over a specified number of years with 
optional monthly contributions.`,
	Example: `
finkit interest compound --initial 10000 --monthly 500 --years 10 --rate 5
finkit interest compound --initial 5000 --monthly 200 --years 5 --rate 7.5
finkit interest compound --initial 50000 --monthly 1000 --years 20 --rate 8`,
	RunE: func(cmd *cobra.Command, args []string) error {

		app := cmd.Context().Value("app").(*bootstrap.App)

		interest := app.Compound.Do(initialInvestment, monthlyContribution, years, rate)
		profit := interest - (initialInvestment + monthlyContribution*float64(years*12))

		totalInvested := initialInvestment + monthlyContribution*float64(years*12)

		fmt.Println("Compound Interest Calculation:")
		fmt.Println()
		fmt.Printf("%-25s %15.2f\n", "Initial investment:", initialInvestment)
		fmt.Printf("%-25s %15.2f\n", "Monthly contribution:", monthlyContribution)
		fmt.Printf("%-24s %15.2f%%\n", "Annual rate:", rate)
		fmt.Printf("%-19s %15d years\n", "Investment period:", years)
		fmt.Println("────────────────────────────────────────────────")
		fmt.Printf("%-25s %15.2f\n", "Total invested:", totalInvested)
		fmt.Printf("%-25s %15.2f\n", "Interest earned:", profit)
		fmt.Printf("%-25s %15.2f\n", "Final value:", interest)
		return nil
	},
}

func init() {
	interestCmd.AddCommand(compoundCmd)

	compoundCmd.Flags().Float64Var(&initialInvestment, "initial", 0, "Initial investment amount")
	compoundCmd.Flags().Float64Var(&monthlyContribution, "monthly", 0, "Monthly contribution amount")
	compoundCmd.Flags().IntVar(&years, "years", 0, "Number of years for investment")
	compoundCmd.Flags().Float64Var(&rate, "rate", 0, "Annual interest rate (as percentage)")
}
