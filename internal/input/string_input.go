package input

type stringInput struct {
	Value string
}

func NewStringInput(value string) stringInput {
	return stringInput{
		Value: value,
	}
}

func (s stringInput) GetAsString() string {
	return s.Value
}
