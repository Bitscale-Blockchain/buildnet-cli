package config_test

import (
	"bitscale/buildnet/lib/config"
	"testing"
)

func TestConfigurationEvents(t *testing.T) {
	// Create a mock event
	mockConfigFilePath := "test/config.yaml"
	mockEvent := &config.LoadConfiguratonEvent{ConfigFilePath: mockConfigFilePath}

	// Create a ConfigurationLoadedEventHandler
	handler := &config.ConfigurationLoadedEventHandler{}

	// Call HandleEvent
	err := handler.HandleEvent(mockEvent)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Write additional test cases as needed
}
