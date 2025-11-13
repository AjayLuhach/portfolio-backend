package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/blog/app"
	"github.com/ajay/portfolio-backend/internal/blog/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

// Register wires blog endpoints.
func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/blogs", func(r chi.Router) {
		r.Post("/", h.createDraft)
		r.Put("/{id}/publish", h.publish)
		r.Get("/trending", h.trending)
		r.Get("/{id}/history", h.history)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) createDraft(w http.ResponseWriter, r *http.Request) {
	var draft domain.Post
	if err := json.NewDecoder(r.Body).Decode(&draft); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	post, err := h.svc.CreateDraft(r.Context(), draft)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, post)
}

func (h handler) publish(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := h.svc.Publish(r.Context(), id)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, post)
}

func (h handler) trending(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.Trending(r.Context())
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, posts)
}

func (h handler) history(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	history, err := h.svc.History(r.Context(), id)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, history)
}
