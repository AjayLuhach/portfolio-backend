package main

import (
	"log"

	"github.com/ajay/portfolio-backend/internal/blog/module"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
)

func main() {
	if err := bootstrap.Run("blog", module.Registrar()); err != nil {
		log.Fatal(err)
	}
}
