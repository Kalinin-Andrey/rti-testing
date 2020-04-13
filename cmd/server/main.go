package main

import (
	"log"

	"github.com/Kalinin-Andrey/rti-testing/pkg/config"

	commonApp "github.com/Kalinin-Andrey/rti-testing/internal/app"
	"github.com/Kalinin-Andrey/rti-testing/internal/app/api"
)


func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("Can not load the config")
	}
	app := api.New(commonApp.New(*cfg), *cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error while application is running: %s", err.Error())
	}
}

