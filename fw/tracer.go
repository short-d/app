package fw

type Segment interface {
	End()
	Next(name string) Segment
}

type Tracer interface {
	BeginTrace(name string) Segment
}
