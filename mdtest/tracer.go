package mdtest

import (
	"strings"

	"github.com/short-d/app/fw"
)

var _ fw.Tracer = (*TracerFake)(nil)

type TracerFake struct {
	Traces []string
}

func (t *TracerFake) BeginTrace(name string) fw.Segment {
	return &TraceFake{
		name:   name,
		tracer: t,
		trace:  make([]string, 0),
	}
}

func NewTracerFake() TracerFake {
	return TracerFake{
		Traces: make([]string, 0),
	}
}

var _ fw.Segment = (*TraceFake)(nil)

type TraceFake struct {
	tracer *TracerFake
	name   string
	trace  []string
}

func (t *TraceFake) End() {
	t.trace = append(t.trace, t.name)
	t.tracer.Traces = append(t.tracer.Traces, strings.Join(t.trace, "->"))
}

func (t TraceFake) Next(name string) fw.Segment {
	t.trace = append(t.trace, t.name)
	return &TraceFake{
		tracer: t.tracer,
		name:   name,
		trace:  t.trace,
	}
}
