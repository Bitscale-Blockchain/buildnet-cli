package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use: "scaffold [chain-type]",
	Short: `Scaffold a new project based on a default configuration file 
	for a specific type of chain (e.g., defi, gaming, social, dex)`,
	Long: `Scaffold a new project based on a default configuration file for a specific type of chain. Example: alstra-cli scaffold defi`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		chainType := args[0]

		fmt.Printf("Scaffolding project for chain type '%s'\n", chainType)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

	// Customize usage and error messages
	scaffoldCmd.SetUsageTemplate(`Usage: buildnet-cli scaffold [chain-type]

Scaffold a new project based on a default configuration file for a specific type of
chain (e.g., defi, gaming, social, dex)

Example:
  buildnet-cli scaffold defi
`)
}
