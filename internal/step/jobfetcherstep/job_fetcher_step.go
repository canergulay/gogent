package jobfetcherstep

import (
	"github.com/canergulay/gogent/internal/input"
	"github.com/canergulay/gogent/internal/output"
	"github.com/canergulay/gogent/pkg/client/http"
)

type fetchJobListingsStep struct {
	httpClient http.HttpClient
}

func NewFetchJobListingsStep() fetchJobListingsStep {
	return fetchJobListingsStep{}
}

func (s *fetchJobListingsStep) Execute(input input.Input) (output.OutPut, error) {
	return nil, nil
}
