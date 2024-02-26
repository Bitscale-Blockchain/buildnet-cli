package cmd

import (
	"bitscale/buildnet/lib"
	"bitscale/buildnet/lib/config"
	"bitscale/buildnet/lib/event"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [config-file]",
	Short: "Create a new project based on the provided configuration file",
	Long:  `Create a new project based on the provided configuration file. Example: alstra-cli create config.json`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]

		lib.GetEventBus().Publish(
			config.StartConfigurationLoadingEvent, event.Event{
				Data: configFile,
				Type: config.StartConfigurationLoadingEvent,
			})
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Customize usage and error messages
	createCmd.SetUsageTemplate(`Usage: buildnet-cli create [config-file]

Create a new project based on the provided configuration file

Example: buildnet-cli create config.json
`)
}
