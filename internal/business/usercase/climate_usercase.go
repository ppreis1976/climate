package usercase

import (
	"climate/internal/business/gateway"
	"climate/internal/business/model"
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
)

type (
	ClimateUserCase interface {
		Get(ctx context.Context, zipCodeID model.ZipCodeID) (model.Climate, error)
	}

	climateUserCase struct {
		climateService gateway.ClimateService
		tracer         trace.Tracer
	}
)

func NewClimateUserCase(climateService gateway.ClimateService, trace trace.Tracer) ClimateUserCase {
	return climateUserCase{
		climateService: climateService,
		tracer:         trace,
	}
}

func (u climateUserCase) Get(ctx context.Context, zipCodeID model.ZipCodeID) (model.Climate, error) {
	// Start a new span
	ctxSpan, span := u.tracer.Start(ctx, "climate-usercase")
	defer span.End()

	trace.SpanFromContext(ctxSpan).AddEvent("get climate usercase")

	tracerID := span.SpanContext().TraceID().String()
	fmt.Println("User Case Tracer ID: ", tracerID)

	if err := zipCodeID.Validate(); err != nil {
		return model.Climate{}, err
	}

	climate, err := u.climateService.Get(ctxSpan, zipCodeID)
	if err != nil {
		return model.Climate{}, err
	}

	return *climate, nil
}
