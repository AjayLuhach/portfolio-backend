package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/interaction/module"
)

func main() {
	if err := bootstrap.Run("interaction", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
