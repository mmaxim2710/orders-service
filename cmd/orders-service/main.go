package main

import (
	ordersApp "github.com/mmaxim2710/orders-service/internal/pkg/app"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"log"
)

func main() {
	if err := utils.InitLogger(); err != nil {
		log.Fatal(err)
	}

	app, err := ordersApp.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
