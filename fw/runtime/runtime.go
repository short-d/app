package runtime

type Caller struct {
	FullFilename string
	LineNumber   int
}

type Runtime interface {
	LockOSThread()
	Caller(numLevelsUp int) (Caller, error)
}
