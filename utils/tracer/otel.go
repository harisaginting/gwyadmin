package tracer

import (
	"context"
	"log"

	"github.com/harisaginting/gwyn/utils/helper"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"

	// stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

var span oteltrace.Span
var tracer oteltrace.Tracer

func InitTracer() {
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// jaeger
	exp, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost("localhost"),
			jaeger.WithAgentPort("6831"),
		),
	)
	log.Println("exporter : ", exp)
	// stdout
	// exp, err := stdout.New(stdout.WithPrettyPrint())

	if err != nil {
		log.Fatal(err)
		return
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ginting-srv"),
		)),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	tracer = tp.Tracer("gwyn-app")
}

func Span(ctx context.Context, name string) oteltrace.Span {
	_, span = tracer.Start(ctx, name)
	return span
}

func SetAttributeString(span oteltrace.Span, key string, value interface{}) {
	val := helper.ForceString(value)
	span.SetAttributes(attribute.String(key, val))
}

func SetAttributeInt(span oteltrace.Span, key string, value interface{}) {
	val := helper.ForceInt(value)
	span.SetAttributes(attribute.Int(key, val))
}

func addEvent(span oteltrace.Span, event string) {
	span.AddEvent(event)
}
