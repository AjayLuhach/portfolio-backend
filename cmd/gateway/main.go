package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/gateway/module"
)

func main() {
	if err := bootstrap.Run("gateway", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
