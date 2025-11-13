package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/analytics/module"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
)

func main() {
	if err := bootstrap.Run("analytics", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
