package fw

type Caller struct {
	FullFilename string
	LineNumber   int
}

type ProgramRuntime interface {
	Caller(numLevelsUp int) (Caller, error)
}
