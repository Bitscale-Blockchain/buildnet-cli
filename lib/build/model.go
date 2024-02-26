package build

import (
	"bitscale/buildnet/lib/event"
)

type Pipeline struct {
	Name   string
	Stages []*Stage
}

type Stage struct {
	Name  string
	Tasks []*Task
}

type Task struct {
	Name     string
	Function func() error
}

type BuildContext struct {
	event.EventContext
	State         string
	Errors        []error
	Pipeline      Pipeline
	Configuration *BuildConfiguration
	Progress      *BuildProgress
	Results       *BuildResults
	Environment   *BuildEnvironment
	Artifacts     []string
	Dependencies  []string
	Metadata      map[string]interface{}
}

// BuildConfiguration represents the configuration settings for the build.
type BuildConfiguration struct{}

// BuildProgress represents the progress of the build.
type BuildProgress struct{}

// BuildResults represents the results of the build.
type BuildResults struct{}

// BuildEnvironment represents the environment in which the build is being executed.
type BuildEnvironment struct{}

// NewBuildState creates a new BuildState with the initial state.
func NewBuildContext(eventBus *event.EventBus) *BuildContext {
	return &BuildContext{
		EventContext: event.EventContext{
			EventBus: eventBus,
		},
		State:         "started",
		Configuration: &BuildConfiguration{},
		Progress:      &BuildProgress{},
		Results:       &BuildResults{},
		Errors:        []error{},
		Environment:   &BuildEnvironment{},
	}
}
