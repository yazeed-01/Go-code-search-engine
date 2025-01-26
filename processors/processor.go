package processors

type Processor interface {
	Process(content string) interface{}
}

type ProcessorFactory struct{}

func (f *ProcessorFactory) CreateProcessor(processorType string) Processor {
	switch processorType {
	case "class":
		return &ClassProcessor{}
	case "method":
		return &MethodProcessor{}
	case "variable":
		return &VariableProcessor{}
	case "relationship":
		return &RelationshipProcessor{}
	default:
		return nil
	}
}
