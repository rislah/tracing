package user

import (
	"context"
	"errors"
	"sync"
	"toy/schema/usergrpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	usergrpc.UnimplementedUserServiceServer
	Svc service
}

func (s *Server) GetUserByName(ctx context.Context, req *usergrpc.GetUserByNameReq) (*usergrpc.GetUserByNameResponse, error) {
	usr, err := s.Svc.GetUserByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &usergrpc.GetUserByNameResponse{User: &usergrpc.User{
		Id:           usr.Id,
		Username:     usr.Username,
		PasswordHash: usr.PasswordHash,
	}}, nil
}

type service struct {
	store  *userStore
	tracer trace.Tracer
}

func NewService(store *userStore) service {
	return service{
		store:  store,
		tracer: otel.GetTracerProvider().Tracer("user-service"),
	}
}

func (s service) GetUserByName(ctx context.Context, username string) (User, error) {
	ctxSpan, span := s.tracer.Start(ctx, "GetUserByName")
	span.SetAttributes(attribute.String("username", username))
	defer span.End()
	return s.store.GetByName(ctxSpan, username)
}

type User struct {
	Id           string
	Username     string
	PasswordHash string
}

type Operation string

const (
	OperationAdd       Operation = "add"
	OperationGetByName Operation = "get_by_name"
)

type userStore struct {
	users  map[string]User
	mu     sync.RWMutex
	tracer trace.Tracer
}

func NewStore() *userStore {
	tracer := otel.Tracer("user-store")
	return &userStore{
		users:  make(map[string]User),
		tracer: tracer,
	}
}

func (us *userStore) newSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	ctx, span := us.tracer.Start(ctx, name)
	span.SetAttributes(
		semconv.DBSystemPostgreSQL,
	)
	return ctx, span
}

func (us *userStore) Add(ctx context.Context, user User) error {
	_, span := us.newSpan(ctx, string(OperationAdd))
	span.SetAttributes(attribute.String("user_id", user.Id), attribute.String("username", user.Username))
	defer span.End()

	us.mu.Lock()
	defer us.mu.Unlock()
	us.users[user.Username] = user
	return nil
}

func (us *userStore) GetByName(ctx context.Context, username string) (User, error) {
	_, span := us.newSpan(ctx, string(OperationGetByName))
	span.SetAttributes(attribute.String("username", username))
	span.SetAttributes(attribute.Bool("found", true))
	defer span.End()

	us.mu.RLock()
	defer us.mu.RUnlock()
	usr, ok := us.users[username]
	if !ok {
		span.SetAttributes(attribute.Bool("found", false))
		return User{}, errors.New("not found")
	}

	return usr, nil
}
