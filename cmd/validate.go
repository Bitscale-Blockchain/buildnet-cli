package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate [config-file]",
	Short: "Validate the provided configuration file",
	Long:  `Validate the provided configuration file. Example: alstra-cli validate config.json`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]

		fmt.Printf("Validating configuration file '%s'\n", configFile)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Customize usage and error messages
	validateCmd.SetUsageTemplate(`Usage: buildnet-cli validate [config-file]

Validate the provided configuration file

Example:
  buildnet-cli validate config.json
`)
}
