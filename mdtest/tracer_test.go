package mdtest

import "testing"

func TestTracerFake(t *testing.T) {
	tracer := NewTracerFake()
	segment := tracer.BeginTrace("a")
	segment = segment.Next("b")
	segment = segment.Next("c")
	segment.End()

	exp := []string{
		"a->b->c",
	}

	SameElements(t, exp, tracer.Traces)
}
