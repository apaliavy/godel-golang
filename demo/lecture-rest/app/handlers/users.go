package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/errors"
	"github.com/apaliavy/godel-golang/demo/lecture-rest/app/users"
)

type Users struct {
	logger *logrus.Logger
	repo   *users.Repository
}

func NewUsersHandler(repo *users.Repository) *Users {
	return &Users{
		logger: logrus.New(),
		repo:   repo,
	}
}

func (uh *Users) Create(w http.ResponseWriter, r *http.Request) {
	payload := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		uh.logger.WithError(err).Error("failed to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.createUser(r.Context(), payload); err != nil {
		uh.logger.WithError(err).Error("failed to create user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *Users) Get(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		uh.logger.WithError(err).Error("incorrect id received in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = uh.repo.Get(r.Context(), uid)
	if err != nil && err == errors.ErrUserNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		uh.logger.WithError(err).Error("received non-expected error from users repository")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// todo: write user response
}

func (uh *Users) Update(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		uh.logger.WithError(err).Error("incorrect id received in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		uh.logger.WithError(err).Error("failed to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure user exists, try to find it in the database
	_, err = uh.repo.Get(r.Context(), uid)
	if err != nil && err != errors.ErrUserNotFound {
		uh.logger.WithError(err).Error("received non-expected error from users repository")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// resource not found - let's create it
	if err != nil && err == errors.ErrUserNotFound {
		if err := uh.createUser(r.Context(), payload); err != nil {
			uh.logger.WithError(err).Error("received non-expected error from users repository")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}

	// user exists - try to update it
	if err = uh.repo.Update(r.Context(), payload); err != nil {
		uh.logger.WithError(err).Error("received non-expected error from users repository")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uh *Users) Delete(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		uh.logger.WithError(err).Error("incorrect id received in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = uh.repo.Delete(r.Context(), uid); err != nil {
		uh.logger.WithError(err).Error("received non-expected error from users repository")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uh *Users) Modify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (uh *Users) createUser(ctx context.Context, payload *users.User) error {
	return uh.repo.Create(ctx, payload)
}
