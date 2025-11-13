package transport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/worker/app"
	"github.com/ajay/portfolio-backend/pkg/httpx"
)

func Register(r chi.Router, svc *app.Service) {
	r.Post("/tasks/{job}", func(w http.ResponseWriter, r *http.Request) {
		job := chi.URLParam(r, "job")
		payload := map[string]any{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil && !errors.Is(err, io.EOF) {
			httpx.Error(w, http.StatusBadRequest, "invalid payload")
			return
		}
		if err := svc.Enqueue(r.Context(), job, payload); err != nil {
			httpx.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		httpx.JSON(w, http.StatusAccepted, map[string]string{"job": job})
	})
}
