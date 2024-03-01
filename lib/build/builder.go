package build

import (
	"bitscale/buildnet/lib/event"
	"bitscale/buildnet/lib/utils"
	"fmt"
)

// NewBuildConfiguration creates a new BuildConfiguration with the provided configuration.
// NewBuildEnvironment creates a new BuildEnvironment with the given configuration.
func NewBuildEnvironment(cfg Configuration) (*BuildEnvironment, error) {
	currentWorkingDir, err := utils.GetCurrentDirectory()
	if err != nil {
		return nil, fmt.Errorf("error getting current working directory: %v", err)
	}

	return &BuildEnvironment{
		WorkingDir: currentWorkingDir,
		ProjectDir: currentWorkingDir,
	}, nil
}

// NewBuildConfiguration creates a new BuildConfiguration with the provided configuration.
func NewBuildConfiguration(cfg Configuration, pipeline Pipeline) *BuildConfiguration {
	return &BuildConfiguration{
		Configuration: cfg,
		Pipeline:      pipeline,
	}
}

func NewBuildContext(
	eventBus *event.EventBus,
	configuration *BuildConfiguration,
	environment *BuildEnvironment) *BuildContext {

	return &BuildContext{
		EventContext: event.EventContext{
			EventBus: eventBus,
		},
		State:         "started",
		Configuration: configuration,
		Progress:      &BuildProgress{},
		Results:       &BuildResults{},
		Errors:        []error{},
		Environment:   environment,
	}
}

func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		pipeline: &Pipeline{},
	}
}

func (b *PipelineBuilder) WithName(name string) *PipelineBuilder {
	b.pipeline.Name = name
	return b
}

func (b *PipelineBuilder) AddStage(name string) *PipelineBuilder {
	b.currentStage = &Stage{Name: name}
	b.pipeline.Stages = append(b.pipeline.Stages, b.currentStage)
	return b
}

func (b *PipelineBuilder) AddTask(name string, fn func(context *BuildContext) error) *PipelineBuilder {
	if b.currentStage == nil {
		panic("No stage added. Please add a stage first.")
	}
	b.currentStage.Tasks = append(b.currentStage.Tasks, &Task{Name: name, Execute: fn})
	return b
}

func (b *PipelineBuilder) Build() *Pipeline {
	return b.pipeline
}

// ExecutePipeline executes the build pipeline.
func (context *BuildContext) ExecutePipeline() error {
	// Publish the pipeline started event
	eventBus := context.EventBus
	eventBus.Publish(PipelineStartedEvent, event.Event{Type: PipelineStartedEvent, Data: context})

	buildConfig := context.Configuration
	// Execute the stages of the pipeline
	for _, stage := range buildConfig.Pipeline.Stages {
		// Publish the stage started event
		eventBus.Publish(StageStartedEvent, event.Event{Type: StageStartedEvent, Data: context})

		// Execute the tasks in the stage
		for _, task := range stage.Tasks {
			// Publish the task started event
			eventBus.Publish(TaskStartedEvent, event.Event{Type: TaskStartedEvent, Data: context})

			// Execute the task
			err := task.Execute(context)
			if err != nil {
				// Publish the task error event
				eventBus.Publish(TaskErrorEvent, event.Event{Type: TaskErrorEvent, Data: err})
				return err
			}

			// Publish the task completed event
			eventBus.Publish(TaskCompletedEvent, event.Event{Type: TaskCompletedEvent, Data: context})
		}

		// Publish the stage completed event
		eventBus.Publish(StageCompletedEvent, event.Event{Type: StageCompletedEvent, Data: context})
	}

	// Publish the pipeline completed event
	eventBus.Publish(PipelineCompletedEvent, event.Event{Type: PipelineCompletedEvent, Data: context})

	return nil
}
func (context *BuildContext) GetTargetDir() string {
	return context.Environment.TargetDir
}
func (context *BuildContext) GetProjectDir() string {
	return context.Environment.ProjectDir
}
