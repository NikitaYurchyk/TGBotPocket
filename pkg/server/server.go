package server

import (
	"context"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/http"
	"strconv"
)

type AuthServer struct {
	server  *http.Server
	storage repository.TokenRepo
	client  *pocket.Client

	redirectUrl string
}

func NewAuthServer(redirectUrl string, storage repository.TokenRepo, client *pocket.Client) *AuthServer {
	return &AuthServer{
		redirectUrl: redirectUrl,
		storage:     storage,
		client:      client,
	}
}

func (s *AuthServer) Start() error {
	s.server = &http.Server{
		Handler: s,
		Addr:    ":80",
	}

	return s.server.ListenAndServe()
}

func (s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	chatIDQuery := r.URL.Query().Get("chat_id")
	if chatIDQuery == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDQuery, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	requestToken, err := s.storage.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to get request token"))
		return
	}

	authResp, err := s.client.Authorize(context.Background(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to authorize at Pocket"))
		return
	}

	if err := s.storage.Save(chatID, authResp.AccessToken, repository.AccessTokens); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to save access token to storage"))
		return
	}

	w.Header().Set("Location", s.redirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)
}
