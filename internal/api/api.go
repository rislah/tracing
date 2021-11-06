package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	clients "toy/internal"
)

type Service struct {
	userClient          clients.UserClient
	authenticatorClient clients.AuthenticatorClient
}

func NewService(userClient clients.UserClient, authClient clients.AuthenticatorClient) Service {
	return Service{userClient: userClient, authenticatorClient: authClient}
}

type Server struct {
	svc Service
}

func NewServer(svc Service) Server {
	return Server{svc}
}

type AuthenticatorPasswordReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticatorPasswordResponse struct {
	Token string `json:"token"`
}

func (s *Server) AuthenticatePassword(w http.ResponseWriter, r *http.Request) {
	areq := AuthenticatorPasswordReq{}
	if err := json.NewDecoder(r.Body).Decode(&areq); err != nil {
		log.Fatal(err)
	}

	ctx := r.Context()
	resp, err := s.svc.authenticatorClient.AuthenticatePassword(ctx, areq.Username, areq.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	writeJSON(w, &AuthenticatorPasswordResponse{Token: resp.Token})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	b, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
