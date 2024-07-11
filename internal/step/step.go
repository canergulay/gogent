package step

import (
	"github.com/canergulay/gogent/internal/input"
	"github.com/canergulay/gogent/internal/output"
)

type Step interface {
	Execute(input input.Input) (output.OutPut, error)
}
