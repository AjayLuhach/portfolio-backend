package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/worker/module"
)

func main() {
	if err := bootstrap.Run("worker", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
