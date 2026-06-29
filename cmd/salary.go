/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"finkit/internal/bootstrap"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// salaryCmd represents the salary command
var salaryCmd = &cobra.Command{
	Use:   "salary <amount> <country>",
	Short: "Calculate tax and net salary for a given gross salary",
	Long: `Calculate the estimated taxes and net salary based on a gross salary amount and country.

The command will display:
- Gross salary (input amount)
- Estimated taxes based on country tax brackets
- Net salary (gross salary minus taxes)`,
	Example: `
finkit tax salary 50000 ES
finkit tax salary 75000 ES
finkit tax salary 120000 ES`,
	RunE: func(cmd *cobra.Command, args []string) error {

		app := cmd.Context().Value("app").(*bootstrap.App)

		salary, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return errors.New("salary must be a number")
		}

		country := args[1]

		taxAmount := app.Tax.CalculateTaxSalary(salary, country)

		fmt.Printf("Gross salary: %.2f\n", salary)
		fmt.Printf("Estimated taxes: %.2f\n", taxAmount)
		fmt.Printf("Net salary: %.2f\n", salary-taxAmount)
		return nil
	},
}

func init() {
	taxCmd.AddCommand(salaryCmd)
}
