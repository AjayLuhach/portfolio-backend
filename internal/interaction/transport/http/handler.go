package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/interaction/app"
	"github.com/ajay/portfolio-backend/internal/interaction/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/blogs/{id}", func(r chi.Router) {
		r.Post("/comments", h.createComment)
		r.Get("/comments", h.listComments)
		r.Post("/reactions", h.createReaction)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) createComment(w http.ResponseWriter, r *http.Request) {
	var payload domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	payload.BlogID = chi.URLParam(r, "id")
	comment, err := h.svc.CreateComment(r.Context(), payload)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, comment)
}

func (h handler) createReaction(w http.ResponseWriter, r *http.Request) {
	var payload domain.Reaction
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	payload.BlogID = chi.URLParam(r, "id")
	reaction, err := h.svc.CreateReaction(r.Context(), payload)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, reaction)
}

func (h handler) listComments(w http.ResponseWriter, r *http.Request) {
	blogID := chi.URLParam(r, "id")
	comments, err := h.svc.ListComments(r.Context(), blogID)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, comments)
}
