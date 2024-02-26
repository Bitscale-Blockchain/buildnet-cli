package app

import (
	"bitscale/buildnet/lib/build"
)

var Pipeline = build.NewPipelineBuilder().
	WithName("MyPipeline").
	AddStage("Stage1").
	AddTask("Task1", func() error {
		// Task1 implementation
		return nil
	}).
	AddTask("Task2", func() error {
		// Task2 implementation
		return nil
	}).
	AddStage("Stage2").
	AddTask("Task3", func() error {
		// Task3 implementation
		return nil
	}).
	Build()
