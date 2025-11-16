package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/analytics/module"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
)

// main is intentionally tiny: each service passes its registrar into bootstrap.
// bootstrap.Run handles config, logging, router setup, and HTTP server lifecycle,
// leaving this file as the canonical "entry point" Go expects.
func main() {
	if err := bootstrap.Run("analytics", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
