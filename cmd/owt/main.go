package main

import (
	"log"

	"contacts_api/modules/config"
	"contacts_api/modules/database"
	"contacts_api/modules/rest"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init config
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// init db
	db, err := database.Init(conf.Database)
	if err != nil {
		panic(err)
	}

	// init server
	app := fiber.New()

	// setup handlers
	rest.Setup(app, db)

	log.Fatal(app.Listen(":" + conf.ServerPort))
}
