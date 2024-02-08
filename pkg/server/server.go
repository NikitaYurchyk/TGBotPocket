package server

import (
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/http"
	"strconv"
)

type AuthServer struct {
	server       *http.Server
	pocketClient *pocket.Client
	repo         repository.TokenRepository
	redirectURL  string
}

func InitAuthServer(pocketClient *pocket.Client, repo repository.TokenRepository, redirect string) *AuthServer {
	return &AuthServer{pocketClient: pocketClient, repo: repo, redirectURL: redirect}
}

func (as *AuthServer) Start() error {
	as.server = &http.Server{
		Handler: as,
		Addr:    ":3232",
	}
	return as.server.ListenAndServe()
}

func (as *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := as.repo.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authResp, err := as.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = as.repo.Save(chatID, authResp.AccessToken, repository.AccessTokens); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Location", as.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)

}
