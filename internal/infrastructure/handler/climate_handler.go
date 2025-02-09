package handler

import (
	"climate/internal/business/model"
	"climate/internal/business/usercase"
	"climate/internal/infrastructure/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const ZipCode = "zip_code"

func ClimateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	// Initialize the tracer
	tracer := service.NewTracer()

	// Start a new span
	ctxSpan, span := tracer.Start(ctx, "climate-handler")
	defer span.End()

	trace.SpanFromContext(ctxSpan).AddEvent("climate handler before usercase")

	tracerID := span.SpanContext().TraceID().String()
	fmt.Println("Handler Tracer ID: ", tracerID)

	if r.Method != http.MethodGet {
		handleZipCodeError(w, errors.New(model.ErrMethodNotAllowed))
		return
	}

	zipCode := r.Header.Get(ZipCode)
	if zipCode == "" {
		http.Error(w, "zip_code header is missing", http.StatusBadRequest)
		return
	}

	client := &http.Client{} // Use http.Client which implements HTTPClient

	climateService := service.NewClimateService(client, tracer)
	zipCodeUserCase := usercase.NewClimateUserCase(climateService, tracer)

	climate, err := zipCodeUserCase.Get(ctxSpan, model.ZipCodeID(zipCode))
	if err != nil {
		handleZipCodeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(climate)
	if err != nil {
		return
	}
}
func handleZipCodeError(w http.ResponseWriter, err error) {
	if err.Error() == model.ErrMethodNotAllowed {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	if err.Error() == model.ErrZipCodeIDInvalid {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err.Error() == model.ErrZipCodeNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
