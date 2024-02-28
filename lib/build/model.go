package build

import (
	"bitscale/buildnet/lib/event"
)

type Configuration struct {
	Data interface{}
}

type PipelineBuilder struct {
	pipeline     *Pipeline
	currentStage *Stage
}

// PipelineFactory is the interface for creating pipelines.
type PipelineFactory interface {
	CreatePipeline() *Pipeline
}

type Pipeline struct {
	Name   string
	Stages []*Stage
}

type Stage struct {
	Name  string
	Tasks []*Task
}

type Task struct {
	Name    string
	Execute func(context *BuildContext) error
}

type BuildContext struct {
	event.EventContext
	State         string
	Errors        []error
	Configuration *BuildConfiguration
	Progress      *BuildProgress
	Results       *BuildResults
	Environment   *BuildEnvironment
	Artifacts     []string
	Dependencies  []string
	Metadata      map[string]interface{}
}

// BuildConfiguration represents the configuration settings for the build.
type BuildConfiguration struct {
	Configuration
	Pipeline Pipeline
}

// BuildProgress represents the progress of the build.
type BuildProgress struct{}

// BuildResults represents the results of the build.
type BuildResults struct{}

// BuildEnvironment represents the environment in which the build is being executed.
type BuildEnvironment struct {
	WorkingDir string
	ProjectDir string
	TargetDir  string
}
