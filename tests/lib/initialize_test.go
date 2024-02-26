package lib

import (
	"testing"

	"bitscale/buildnet/lib"
)

func TestGetEventBus(t *testing.T) {
	// Create a new test event bus
	testEventBus := lib.GetEventBus()

	// Get the event bus
	actualEventBus := lib.GetEventBus()

	// Check if the actual event bus is the same as the test event bus
	if actualEventBus != testEventBus {
		t.Errorf("GetEventBus() = %v, want %v", actualEventBus, testEventBus)
	}
}
