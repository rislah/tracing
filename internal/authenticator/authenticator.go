package authenticator

import (
	"context"
	"errors"
	clients "toy/internal"
	credentials "toy/internal/credentials"
	"toy/internal/jwt"
	authenticatorgrpc "toy/schema/authenticatorgrpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	authenticatorgrpc.UnimplementedAuthenticatorServer
	A Authenticator
}

func (s *Server) AuthenticatePassword(ctx context.Context, req *authenticatorgrpc.AuthenticatePasswordReq) (*authenticatorgrpc.AuthenticatePasswordResponse, error) {
	jwt, err := s.A.AuthenticatePassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return &authenticatorgrpc.AuthenticatePasswordResponse{Token: jwt}, nil
}

type Authenticator struct {
	userClient clients.UserClient
	jwtWrapper jwt.Wrapper
	tracer     trace.Tracer
}

func NewAuthenticator(userClient clients.UserClient, jwtWrapper jwt.Wrapper) Authenticator {
	return Authenticator{
		userClient: userClient,
		jwtWrapper: jwtWrapper,
		tracer:     otel.GetTracerProvider().Tracer("authenticator-service"),
	}
}

func (a Authenticator) AuthenticatePassword(ctx context.Context, req *authenticatorgrpc.AuthenticatePasswordReq) (string, error) {
	var (
		name     = req.Username
		password = req.Password
	)
	ctxSpan, span := a.tracer.Start(ctx, "AuthenticatePassword")
	defer span.End()

	user, err := a.userClient.GetUserByName(ctxSpan, name)
	if err != nil {
		return "", err
	}

	hashedPassword := user.User.PasswordHash

	_, passwordMatchSpan := a.tracer.Start(ctxSpan, "password_match")
	valid, err := credentials.CompareBCrypt(password, hashedPassword)
	if err != nil {
		return "", err
	}
	passwordMatchSpan.End()

	span.SetAttributes(attribute.Bool("password_match", valid))

	if !valid {
		return "", errors.New("pass mismatch")
	}

	claims := jwt.NewUserClaims(name, "role")

	token, err := a.jwtWrapper.Encode(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
