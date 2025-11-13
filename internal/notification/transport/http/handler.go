package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/notification/app"
	"github.com/ajay/portfolio-backend/internal/notification/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/notifications", func(r chi.Router) {
		r.Post("/", h.create)
		r.Get("/", h.list)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) create(w http.ResponseWriter, r *http.Request) {
	var payload domain.Notification
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	n, err := h.svc.Dispatch(r.Context(), payload)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, n)
}

func (h handler) list(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	items, err := h.svc.List(r.Context(), userID)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, items)
}
