package event

import (
	"log"
)

// Publish publishes an event to all subscribers.
func (eb *EventBus) Publish(eventType string, event Event) {
	eb.Registry.mu.RLock()
	defer eb.Registry.mu.RUnlock()

	handlers, ok := eb.Registry.Registry[eventType]
	if !ok {
		return
	}

	for _, handler := range handlers {
		if err := handler.HandleEvent(event, *eb.Context); err != nil {
			log.Printf("Error handling event: %v", err)
		}
	}
}

// Subscribe subscribes to an event type with a handler.
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) error {
	eb.Registry.mu.Lock()
	defer eb.Registry.mu.Unlock()

	eb.Registry.Registry[eventType] = append(eb.Registry.Registry[eventType], handler)
	return nil
}
