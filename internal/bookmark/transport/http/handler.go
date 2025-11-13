package transport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/bookmark/app"
	"github.com/ajay/portfolio-backend/internal/bookmark/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	h := handler{svc: svc}
	r.Route("/bookmarks", func(r chi.Router) {
		r.Post("/{blogId}", h.save)
		r.Get("/", h.list)
		r.Delete("/{id}", h.delete)
	})
}

type handler struct {
	svc *app.Service
}

func (h handler) save(w http.ResponseWriter, r *http.Request) {
	var payload domain.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil && !errors.Is(err, io.EOF) {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	payload.BlogID = chi.URLParam(r, "blogId")
	bookmark, err := h.svc.Save(r.Context(), payload)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, bookmark)
}

func (h handler) list(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	bookmarks, err := h.svc.List(r.Context(), userID)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, bookmarks)
}

func (h handler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, map[string]string{"id": id})
}
