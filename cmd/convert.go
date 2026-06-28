package cmd

import (
	"errors"
	"finkit/internal/bootstrap"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert <amount> <from> <to>",
	Short: "Convert between currencies",
	Long:  `Convert an amount from one currency to another using current exchange rates.`,
	Example: `
finkit currency convert 100 EUR USD
finkit currency convert 50 GBP JPY
finkit currency convert 1000 USD CHF
`,
	Args: cobra.ExactArgs(3),
	ValidArgsFunction: func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {

		app := cmd.Context().Value("FinKit").(*bootstrap.App)

		currencies, err := app.Convert.SupportedCurrencies()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		switch len(args) {

		case 0:
			return nil, cobra.ShellCompDirectiveDefault

		case 1:
			return currencies, cobra.ShellCompDirectiveNoFileComp

		case 2:
			from := args[1]

			var result []string

			for _, currency := range currencies {
				if currency != from {
					result = append(result, currency)
				}
			}

			return result, cobra.ShellCompDirectiveNoFileComp

		default:
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		app := cmd.Context().Value("FinKit").(*bootstrap.App)

		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return errors.New("amount must be a number")
		}

		from := args[1]

		to := args[2]

		result, err := app.Convert.Convert(amount, from, to)
		if err != nil {
			return err
		}
		fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
		return nil
	},
}

func init() {
	currencyCmd.AddCommand(convertCmd)
}
