package config

import (
	"bitscale/buildnet/lib/event"
	"bitscale/buildnet/lib/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	StartConfigurationLoadingEvent = "StartConfigurationLoadingEvent"
	ConfigurationLoadedEvent       = "ConfigurationLoadedEvent"
	ConfigurationSavedEvent        = "ConfigurationSavedEvent"
	ConfigurationValidatedEvent    = "ConfigurationValidatedEvent"
	ConfigurationErrorEvent        = "ConfigurationErrorEvent"
	ConfigurationUpdatedEvent      = "ConfigurationUpdatedEvent"
)

type StartConfigurationLoadingEventHandler struct{}

// InitEventHandlers initializes event handlers for package 1.
func InitEventHandlers(eventBus *event.EventBus) {
	// Initialize and subscribe event handlers for package 1
	eventBus.Subscribe(StartConfigurationLoadingEvent, &StartConfigurationLoadingEventHandler{})
}

// HandleEvent handles a StartConfigurationLoadingEvent.
func (h *StartConfigurationLoadingEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != StartConfigurationLoadingEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	configFile, ok := eventObj.Data.(string)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	// Check if the file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("configuration file '%s' not found", configFile)
	}

	// Read the JSON file using the ReadFile function from the utils package
	data, err := utils.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error reading file:%v", err)
	}

	// Unmarshal the JSON data into the Application struct
	var app *Configuration
	if err := json.Unmarshal(data, &app); err != nil {
		return fmt.Errorf("error unmarshalling JSON data: %v", err)
	}

	log.Printf("Configuration loaded from file: %s", configFile)

	// Publish a ConfigurationLoadedEvent with the loaded application data
	context.EventBus.Publish(
		ConfigurationLoadedEvent, event.Event{
			Type: ConfigurationLoadedEvent,
			Data: app,
		})
	return nil
}
