package build

type PipelineBuilder struct {
	pipeline     *Pipeline
	currentStage *Stage
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

func (b *PipelineBuilder) AddTask(name string, fn func() error) *PipelineBuilder {
	if b.currentStage == nil {
		panic("No stage added. Please add a stage first.")
	}
	b.currentStage.Tasks = append(b.currentStage.Tasks, &Task{Name: name, Function: fn})
	return b
}

func (b *PipelineBuilder) Build() *Pipeline {
	return b.pipeline
}
