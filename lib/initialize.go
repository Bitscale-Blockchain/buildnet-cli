package lib

import (
	"bitscale/buildnet/lib/build"
	"bitscale/buildnet/lib/config"
	"bitscale/buildnet/lib/event"
	"sync"
)

var (
	onceEventBus      sync.Once
	onceEventRegistry sync.Once
	eventBus          *event.EventBus
	eventRegistry     *event.EventTypeRegistry
)

func GetEventBus() *event.EventBus {
	onceEventBus.Do(func() {
		eventBus = &event.EventBus{
			Registry: GetEventRegistry(),
			Context:  &event.EventContext{},
		}
		eventBus.Context.EventBus = eventBus
		config.InitEventHandlers(eventBus)
		build.InitEventHandlers(eventBus)
	})
	return eventBus
}

// GetEventRegistry returns the singleton instance of EventTypeRegistry.
func GetEventRegistry() *event.EventTypeRegistry {
	onceEventRegistry.Do(func() {
		eventRegistry = &event.EventTypeRegistry{
			Registry: make(map[string][]event.EventHandler),
		}
	})
	return eventRegistry
}
