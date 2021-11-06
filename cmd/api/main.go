package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	clients "toy/internal"
	"toy/internal/api"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var (
	addr              = flag.String("addr", ":8080", "")
	authenticatorAddr = flag.String("authenticatorAddr", ":8082", "")
	userAddr          = flag.String("userAddr", ":8081", "")
)

func main() {
	flag.Parse()

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14260/api/traces")))
	if err != nil {
		log.Fatal(err)
	}

	// exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("api-gateway"))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	defer tp.Shutdown(context.Background())

	userClient := clients.NewUserClient(*userAddr)
	authClient := clients.NewAuthenticatorClient(*authenticatorAddr)

	svc := api.NewService(userClient, authClient)
	srv := api.NewServer(svc)

	handler := otelhttp.NewHandler(http.HandlerFunc(srv.AuthenticatePassword), "/auth")
	http.Handle("/auth", handler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
