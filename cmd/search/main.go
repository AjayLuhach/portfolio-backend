package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/search/module"
)

func main() {
	if err := bootstrap.Run("search", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
