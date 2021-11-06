package clients

import (
	"context"
	"toy/schema/authenticatorgrpc"
	"toy/schema/usergrpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type AuthenticatorClient struct{ addr string }

func NewAuthenticatorClient(addr string) AuthenticatorClient {
	return AuthenticatorClient{addr}
}

func (a AuthenticatorClient) newClient() (*grpc.ClientConn, authenticatorgrpc.AuthenticatorClient, error) {
	conn, err := grpc.Dial(a.addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, nil, err
	}

	client := authenticatorgrpc.NewAuthenticatorClient(conn)

	return conn, client, nil
}

func (a AuthenticatorClient) AuthenticatePassword(ctx context.Context, username, password string) (*authenticatorgrpc.AuthenticatePasswordResponse, error) {
	conn, client, err := a.newClient()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return client.AuthenticatePassword(ctx, &authenticatorgrpc.AuthenticatePasswordReq{Username: username, Password: password})
}

type UserClient struct {
	addr string
}

func NewUserClient(addr string) UserClient {
	return UserClient{addr: addr}
}

func (u UserClient) newClient(addr string) (*grpc.ClientConn, usergrpc.UserServiceClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, nil, err
	}

	client := usergrpc.NewUserServiceClient(conn)

	return conn, client, nil
}

func (u UserClient) GetUserByName(ctx context.Context, name string) (*usergrpc.GetUserByNameResponse, error) {
	conn, client, err := u.newClient(u.addr)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return client.GetUserByName(ctx, &usergrpc.GetUserByNameReq{Name: name})
}
