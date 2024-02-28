package ignite

import (
	"bitscale/buildnet/lib/app"
	"bitscale/buildnet/lib/build"
	"log"
)

// MyPipelineFactory is an implementation of PipelineFactory that creates a specific pipeline.
type IgniteCliBuilderPipelineFactory struct{}

func (f *IgniteCliBuilderPipelineFactory) CreatePipeline() *build.Pipeline {
	return build.NewPipelineBuilder().
		WithName("IgniteCosmosPipeline").
		AddStage("IgniteStage").
		AddTask("InitializeTargetDirectoryTask", app.InitializeTargetDirectoryTask).
		AddTask("ScaffoldIgniteProjectTask", ScaffoldIgniteProjectTask).
		AddTask("ScaffoldIgniteModulesTask", ScaffoldIgniteModulesTask).
		AddTask("ScaffoldIgniteTypesTask", ScaffoldIgniteTypesTask).
		AddTask("ScaffoldIgniteMessagesTask", ScaffoldIgniteMessagesTask).
		AddTask("ScaffoldIgniteQueriesTask", ScaffoldIgniteQueriesTask).
		AddStage("Stage2").
		AddTask("Task3", func(context *build.BuildContext) error {
			log.Print("@@@@@@@@@@@@@@@@@Executing task3@@@@@@@@@@@@@@@@@")
			// Task3 implementation
			return nil
		}).
		Build()
}
