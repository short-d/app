package fw

type Caller struct {
	FullFilename string
	LineNumber   int
}

type ProgramRuntime interface {
	LockOSThread()
	Caller(numLevelsUp int) (Caller, error)
}
