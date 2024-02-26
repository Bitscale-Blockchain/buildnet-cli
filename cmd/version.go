package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the application",
	Long: `
Prints the version of the application. This command can be 
used to check the current version of the application. It is useful 
for identifying the version of the application when reporting issues or seeking support.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Customize usage and error messages
	versionCmd.SetUsageTemplate(`Usage: buildnet-cli version

A brief description of your command

Example:
  buildnet-cli version
`)
}
