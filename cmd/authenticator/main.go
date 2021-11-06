package main

import (
	"context"
	"flag"
	"log"
	"net"
	clients "toy/internal"
	"toy/internal/authenticator"
	"toy/internal/jwt"
	"toy/schema/authenticatorgrpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8082", "")
var userAddr = flag.String("userAddr", ":8081", "")

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
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("authenticator"))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	defer tp.Shutdown(context.Background())

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	svc := authenticator.NewAuthenticator(clients.NewUserClient(*userAddr), jwt.NewHS256Wrapper("secret"))
	authServer := &authenticator.Server{A: svc}
	authenticatorgrpc.RegisterAuthenticatorServer(srv, authServer)

	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
