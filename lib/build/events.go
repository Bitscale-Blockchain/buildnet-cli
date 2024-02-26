package build

import (
	"bitscale/buildnet/lib/config"
	"bitscale/buildnet/lib/event"
	"fmt"
	"log"
)

const (
	PipelineStartedEvent   = "PipelineStartedEvent"
	PipelineCompletedEvent = "PipelineCompletedEvent"
	PipelineErrorEvent     = "PipelineErrorEvent"
	StageStartedEvent      = "StageStartedEvent"
	StageCompletedEvent    = "StageCompletedEvent"
	StageErrorEvent        = "StageErrorEvent"
	TaskStartedEvent       = "TaskStartedEvent"
	TaskCompletedEvent     = "TaskCompletedEvent"
	TaskErrorEvent         = "TaskErrorEvent"
)

const (
	PipelineIdleState    = "PipelineIdleState"
	PipelineRunningState = "PipelineRunningState"
	PipelineErrorState   = "PipelineErrorState"
	StageRunningState    = "StageRunningState"
	StageErrorState      = "StageErrorState"
	TaskRunningState     = "TaskRunningState"
	TaskErrorState       = "TaskErrorState"
)

type PipelineStartedEventHandler struct{}
type PipelineCompletedEventHandler struct{}
type PipelineErrorEventHandler struct{}
type StageStartedEventHandler struct{}
type StageCompletedEventHandler struct{}
type StageErrorEventHandler struct{}
type TaskStartedEventHandler struct{}
type TaskCompletedEventHandler struct{}
type TaskErrorEventHandler struct{}

// InitEventHandlers initializes event handlers for package 1.
func InitEventHandlers(eventBus *event.EventBus) {
	// Initialize and subscribe event handlers for package 1
	eventBus.Subscribe(PipelineStartedEvent, &PipelineStartedEventHandler{})
	eventBus.Subscribe(PipelineCompletedEvent, &PipelineCompletedEventHandler{})
	eventBus.Subscribe(PipelineErrorEvent, &PipelineErrorEventHandler{})
	eventBus.Subscribe(StageStartedEvent, &StageStartedEventHandler{})
	eventBus.Subscribe(StageCompletedEvent, &StageCompletedEventHandler{})
	eventBus.Subscribe(StageErrorEvent, &StageErrorEventHandler{})
	eventBus.Subscribe(TaskStartedEvent, &TaskStartedEventHandler{})
	eventBus.Subscribe(TaskCompletedEvent, &TaskCompletedEventHandler{})
	eventBus.Subscribe(TaskErrorEvent, &TaskErrorEventHandler{})
}

// HandleEvent handles a StartApplicationBuildEvent.
func (h *PipelineStartedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *PipelineCompletedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *PipelineErrorEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a StageStartedEventData.
func (h *StageStartedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *StageCompletedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *StageErrorEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *TaskStartedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *TaskCompletedEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}

// HandleEvent handles a CompletedApplicationBuildEvent.
func (h *TaskErrorEventHandler) HandleEvent(eventObj event.Event, context event.EventContext) error {
	// Ensure received valid event type
	if eventObj.Type != PipelineStartedEvent {
		return fmt.Errorf("unexpected event type: %s", eventObj.Type)
	}

	// Type switch to check the type of eventObj.Data
	data, ok := eventObj.Data.(config.Configuration)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected event data type: %T", eventObj.Data)
	}

	log.Printf("Handling event: %s and artifact: %T", eventObj.Type, data)
	return nil
}
