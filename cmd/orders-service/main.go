package main

import (
	"github.com/mmaxim2710/orders-service/internal/pkg/utils"
	"github.com/mmaxim2710/orders-service/internal/service"
	"log"
)

func main() {
	if err := utils.InitLogger(); err != nil {
		log.Fatal(err)
	}
	if err := service.Start(); err != nil {
		log.Fatal(err)
	}
}
