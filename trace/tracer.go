package trace

import "io"

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
}

func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}
