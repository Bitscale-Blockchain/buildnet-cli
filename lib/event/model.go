package event

import (
	"sync"
)

type Event struct {
	Type string
	Data interface{}
}

// Core event state
type EventContext struct {
	EventBus *EventBus
}

// EventHandler is an interface for event handlers.
type EventHandler interface {
	HandleEvent(Event, EventContext) error
}

// EventTypeRegistry is a registry of event types and their corresponding handlers.
type EventTypeRegistry struct {
	Registry map[string][]EventHandler
	mu       sync.RWMutex
	eventBus *EventBus // Reference to the EventBus
}

// EventBus is a central component that manages the distribution of events to subscribers.
type EventBus struct {
	Context  *EventContext
	Registry *EventTypeRegistry
}
