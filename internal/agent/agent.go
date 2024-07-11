package agent

import (
	"github.com/canergulay/gogent/internal/input"
	"github.com/canergulay/gogent/internal/step"
)

type Agent struct {
	Steps []step.Step
}

func NewAgent() Agent {
	return Agent{
		Steps: make([]step.Step, 0),
	}
}

func (a *Agent) AddStep(step step.Step) {
	a.Steps = append(a.Steps, step)
}

func (a *Agent) Execute(input input.Input, stepToPass int) error {
	currentInput := input
	for i, step := range a.Steps {
		if i < stepToPass {
			continue
		}
		currentOutput, err := step.Execute(currentInput)
		if err != nil {
			return err
		}
		currentInput = currentOutput
	}
	return nil
}
