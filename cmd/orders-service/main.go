package main

import (
	"github.com/mmaxim2710/orders-service/config"
	"github.com/mmaxim2710/orders-service/internal/app"
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"log"
)

func main() {
	if err := utils.InitLogger(); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	app.Run(cfg)
}
