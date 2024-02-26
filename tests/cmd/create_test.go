package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"bitscale/buildnet/lib"
	"bitscale/buildnet/lib/config"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommand(t *testing.T) {
	// Set up a temporary directory for the test
	tmpDir := t.TempDir()

	// Create a test configuration file in the temporary directory
	configFile := fmt.Sprintf("%s/config.json", tmpDir)
	err := os.WriteFile(configFile, []byte(`{"key": "value"}`), 0644)
	assert.NoError(t, err)

	// Create a buffer to capture command output
	var buf bytes.Buffer

	// Mock the event bus
	eventBus := lib.GetEventBus()
	eventBus.Subscribe("ConfigurationLoadedEvent", &config.LoadConfiguratonEvent{})

	// Set up the create command with the mocked event bus
	createCmd := &cobra.Command{
		Use:   "create [config-file]",
		Short: "Create a new project based on the provided configuration file",
		Long:  `Create a new project based on the provided configuration file. Example: alstra-cli create config.json`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configFile := args[0]

			// Check if the file exists
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				fmt.Fprintf(&buf, "Error: Configuration file '%s' not found\n", configFile)
				os.Exit(1)
			}
			eventObj := &config.LoadConfiguratonEvent{
				ConfigFilePath: configFile,
			}
			lib.GetEventBus().Publish(config.LoadConfigEvent, eventObj)

		},
	}
	createCmd.SetOut(&buf)

	// Execute the create command with the test configuration file
	createCmd.SetArgs([]string{configFile})
	err = createCmd.Execute()
	assert.NoError(t, err)

	// Assert that the event was published to the event bus
	assert.Contains(t, buf.String(), "Configuration loaded from file")

	// Assert that the event handler was called
	assert.Contains(t, buf.String(), "Configuration loaded from file: config.json")
}
