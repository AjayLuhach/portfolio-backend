package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/gateway/app"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, map[string]any{
			"services": svc.Routes(),
		})
	})
}
