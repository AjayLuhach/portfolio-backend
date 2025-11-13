package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/analytics/app"
	"github.com/ajay/portfolio-backend/internal/analytics/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/analytics", func(r chi.Router) {
		r.Post("/", h.record)
		r.Get("/blog/{id}", h.fetch)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) record(w http.ResponseWriter, r *http.Request) {
	var payload domain.Metrics
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	metrics, err := h.svc.Record(r.Context(), payload)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, metrics)
}

func (h handler) fetch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	metrics, err := h.svc.Fetch(r.Context(), id)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, metrics)
}
