package event

// RegisterEventType registers an event type with its corresponding handlers.
func (r *EventTypeRegistry) RegisterEventType(eventType string, handlers []EventHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Registry[eventType] = handlers
}

// SubscribeToEvent subscribes the provided event handlers to the given event type.
func (r *EventTypeRegistry) SubscribeToEvent(eventType string) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	handlers, ok := r.Registry[eventType]
	if !ok {
		return
	}
	for _, handler := range handlers {
		// Provide the event handler with the EventBus instance
		if h, ok := handler.(interface{ SetEventBus(*EventBus) }); ok {
			h.SetEventBus(r.eventBus)
		}
	}
}

// SetEventBus sets the EventBus instance for the EventTypeRegistry.
func (r *EventTypeRegistry) SetEventBus(eventBus *EventBus) {
	r.eventBus = eventBus
}
