package app

import (
	"bitscale/buildnet/app/ignite"
	"bitscale/buildnet/lib/app"
	"bitscale/buildnet/lib/build"
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
	BlockchainBuildStartedEvent    = "BlockchainBuildStartedEvent"
	BlockchainBuildCompletedEvent  = "BlockchainBuildCompletedEvent"
	ArtifactBuildStartedEvent      = "ArtifactBuildStartedEvent"
	ArtifactBuildCompletedEvent    = "ArtifactBuildCompletedEvent"
)

type StartConfigurationLoadingEventHandler struct{}
type ConfigurationLoadedEventEventHandler struct{}

// InitEventHandlers initializes event handlers for package 1.
func InitEventHandlers(eventBus *event.EventBus) {
	// Initialize and subscribe event handlers for package 1
	eventBus.Subscribe(StartConfigurationLoadingEvent, &StartConfigurationLoadingEventHandler{})
	eventBus.Subscribe(ConfigurationLoadedEvent, &ConfigurationLoadedEventEventHandler{})
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
	var app *app.Blockchain
	if err := json.Unmarshal(data, &app); err != nil {
		return fmt.Errorf("error unmarshalling JSON data: %v", err)
	}

	log.Printf("Configuration loaded from file: %s", configFile)
	var configuration *build.Configuration = &build.Configuration{
		Data: app,
	}
	// Publish a ConfigurationLoadedEvent with the loaded application data
	context.EventBus.Publish(
		ConfigurationLoadedEvent, event.Event{
			Type: ConfigurationLoadedEvent,
			Data: configuration,
		})
	return nil
}

// HandleEvent handles a ConfigurationLoadedEvent.
func (h *ConfigurationLoadedEventEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != ConfigurationLoadedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	configuration, ok := eventObj.Data.(*build.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}
	// Initialize the build configuration
	pipelineFactory := &ignite.IgniteCliBuilderPipelineFactory{}
	buildConfig := build.NewBuildConfiguration(*configuration, *pipelineFactory.CreatePipeline())

	// Initialize the build environment
	currentWorkingDir, err := utils.GetCurrentDirectory()
	if err != nil {
		return fmt.Errorf("counld not get current woriking director: %s", err)
	}
	// Type switch to check the type of eventObj.Data
	app, ok := configuration.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}

	targetDir := fmt.Sprintf("%s/%s", currentWorkingDir, "target")
	projectDir := fmt.Sprintf("%s/%s", targetDir, app.Name)

	environment := &build.BuildEnvironment{
		WorkingDir: currentWorkingDir,
		ProjectDir: projectDir,
		TargetDir:  targetDir,
	}
	// Initialize the build context
	buildContext := build.NewBuildContext(context.EventBus, buildConfig, environment)
	// Publish a ConfigurationLoadedEvent with the loaded application data
	context.EventBus.Publish(
		build.StartBuildEvent, event.Event{
			Type: build.StartBuildEvent,
			Data: buildContext,
		})
	return nil
}
