package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/auth/app"
	"github.com/ajay/portfolio-backend/internal/auth/domain"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

// Register wires auth routes into the router.
func Register(r chi.Router, svc *app.Service) {
	h := handler{service: svc}
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", h.signup)
		r.Post("/login", h.login)
		r.Post("/refresh", h.refresh)
		r.Get("/me", h.me)
	})
}

type handler struct {
	service *app.Service
}

func (h handler) signup(w http.ResponseWriter, r *http.Request) {
	var req domain.Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	tokens, err := h.service.Signup(r.Context(), req)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpx.JSON(w, http.StatusCreated, map[string]any{"tokens": tokens})
}

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	var req domain.Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	tokens, err := h.service.Login(r.Context(), req)
	if err != nil {
		httpx.Error(w, http.StatusUnauthorized, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, map[string]any{"tokens": tokens})
}

func (h handler) refresh(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid payload")
		return
	}
	tokens, err := h.service.Refresh(r.Context(), payload.RefreshToken)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	httpx.JSON(w, http.StatusOK, map[string]any{"tokens": tokens})
}

func (h handler) me(w http.ResponseWriter, r *http.Request) {
	httpx.JSON(w, http.StatusOK, map[string]any{
		"user": map[string]string{
			"id":    "demo-user",
			"email": "demo@example.com",
		},
	})
}
