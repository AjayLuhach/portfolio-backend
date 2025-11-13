package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/search/app"
	"github.com/ajay/portfolio-backend/internal/search/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/search", func(r chi.Router) {
		r.Post("/index", h.index)
		r.Get("/", h.query)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) index(w http.ResponseWriter, r *http.Request) {
	var doc domain.Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	if err := h.svc.Index(r.Context(), doc); err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusAccepted, map[string]string{"id": doc.ID})
}

func (h handler) query(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("q")
	docs, err := h.svc.Query(r.Context(), term)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, docs)
}
