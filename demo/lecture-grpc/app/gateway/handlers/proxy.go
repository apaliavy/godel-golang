package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	pbUsers "github.com/apaliavy/godel-golang/demo/lecture-grpc/app/users/api"
)

type UsersProxy struct {
	uc     pbUsers.UsersClient
	logger *logrus.Entry
}

func NewUsersProxy(uc pbUsers.UsersClient) *UsersProxy {
	return &UsersProxy{
		uc:     uc,
		logger: logrus.NewEntry(logrus.New()).WithField("handler", "users_proxy"),
	}
}

func (p *UsersProxy) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithField("method", "CreateUser")

	req := &pbUsers.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.WithError(err).Error("failed to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := p.uc.CreateUser(r.Context(), req); err != nil {
		logger.WithError(err).Error("received an error from users service")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *UsersProxy) ListUsers(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithField("method", "ListUsers")

	if _, err := p.uc.ListUsers(r.Context(), &pbUsers.ListUsersRequest{}); err != nil {
		logger.WithError(err).Error("received an error from users service")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *UsersProxy) GetUser(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithField("method", "GetUser")

	if _, err := p.uc.GetUser(r.Context(), &pbUsers.GetUserRequest{}); err != nil {
		logger.WithError(err).Error("received an error from users service")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
