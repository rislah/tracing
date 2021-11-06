package main

import (
	"context"
	"flag"
	"log"
	"net"
	user "toy/internal/user"
	"toy/schema/usergrpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8081", "")

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
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("user"))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	defer tp.Shutdown(context.Background())

	store := user.NewStore()
	store.Add(context.Background(), user.User{
		Id:           "1",
		Username:     "kasutaja",
		PasswordHash: "$2a$10$cooR9Q2.ycvu6HEttewRi.cRK6DRR7gYS0POD.u89kzh8AgD0GG7.",
	})

	svc := user.NewService(store)
	userServer := &user.Server{
		Svc: svc,
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	usergrpc.RegisterUserServiceServer(srv, userServer)

	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
