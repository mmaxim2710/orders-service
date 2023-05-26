package main

import (
	"github.com/mmaxim2710/orders-service/config"
	"github.com/mmaxim2710/orders-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
