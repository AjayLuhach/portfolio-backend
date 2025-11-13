package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/bookmark/module"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
)

func main() {
	if err := bootstrap.Run("bookmark", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
