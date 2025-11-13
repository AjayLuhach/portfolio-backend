package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/auth/module"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
)

func main() {
	if err := bootstrap.Run("auth", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
