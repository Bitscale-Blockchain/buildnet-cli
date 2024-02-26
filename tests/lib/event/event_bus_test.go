package event

import (
	"bitscale/buildnet/lib"
	"bitscale/buildnet/lib/event"
	"testing"
)

type mockEvent struct {
	eventType string
}

func (m *mockEvent) EventType() string {
	return m.eventType
}

type mockEventHandler struct {
	handled bool
}

func (m *mockEventHandler) HandleEvent(event event.Event) error {
	m.handled = true
	return nil
}

func TestEventBus(t *testing.T) {
	// Create a new EventBus
	eb := lib.GetEventBus()

	// Create a mock event and handler
	mockEvent := &mockEvent{eventType: "test"}
	mockHandler := &mockEventHandler{}

	// Subscribe the handler to the event type
	err := eb.Subscribe(mockEvent.EventType(), mockHandler)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Publish the event
	eb.Publish(mockEvent.EventType(), mockEvent)

	// Verify that the handler was called
	if !mockHandler.handled {
		t.Errorf("expected handler to be called")
	}
}
