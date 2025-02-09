package service

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// função que retorna um tracer
func NewTracer() trace.Tracer {
	return otel.Tracer("climate-wrapper")
}
