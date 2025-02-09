package service

import (
	"climate/internal/business/gateway"
	"climate/internal/business/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const climateURL = "http://localhost:8080/climate"

func NewClimateService(client HTTPClient, tracer trace.Tracer) gateway.ClimateService {
	return &climateService{
		client: client,
		tracer: tracer,
	}
}

type climateService struct {
	client HTTPClient
	tracer trace.Tracer
}

func (s *climateService) Get(ctx context.Context, zipCodeID model.ZipCodeID) (*model.Climate, error) {
	ctxSpan, span := s.tracer.Start(ctx, "climate-service")
	defer span.End()

	trace.SpanFromContext(ctxSpan).AddEvent("get climate service")

	tracerID := span.SpanContext().TraceID().String()
	fmt.Println("Service Tracer ID: ", tracerID)

	fmt.Println("climate service : " + string(zipCodeID))

	req, err := http.NewRequest("GET", climateURL, nil)
	if err != nil {
		return &model.Climate{}, err
	}

	req = req.WithContext(ctxSpan)
	req.Header.Set("Zip_Code", string(zipCodeID))
	req.Header.Set("X-Request-ID", "12d6b0d3-015e-4268-ab6a-eafe1f3a4006")

	PrintHeadersFromContext(ctxSpan)

	otel.GetTextMapPropagator().Inject(ctxSpan, propagation.HeaderCarrier(req.Header))

	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Println("error : " + err.Error())
		return &model.Climate{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &model.Climate{}, errors.New("failed to fetch data from climate service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.Climate{}, err
	}

	var climate model.Climate
	if err := json.Unmarshal(body, &climate); err != nil {
		return &model.Climate{}, err
	}

	if climate == (model.Climate{}) {
		return &model.Climate{}, errors.New(model.ErrZipCodeNotFound)
	}

	return &climate, nil
}

func PrintHeadersFromContext(ctx context.Context) {
	headers, ok := ctx.Value("headers").(http.Header)
	if !ok {
		fmt.Println("No headers found in context")
		return
	}

	for key, values := range headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
