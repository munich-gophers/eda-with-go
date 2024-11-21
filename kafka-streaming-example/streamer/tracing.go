package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTracing() func() {
	log.Println("Initializing tracing...")
	exporter, err := otlptracehttp.New(context.Background())
	if err != nil {
		log.Fatalf("Failed to create OpenTelemetry exporter: %v", err)
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)

	log.Println("OpenTelemetry tracing initialized")
	return func() {
		_ = tp.Shutdown(context.Background())
	}
}
