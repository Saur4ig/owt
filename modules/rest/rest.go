package rest

import (
	"database/sql"

	"contacts_api/modules/database"
	restapi "contacts_api/modules/rest/api"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

const URLPrefix = "/v1"

func Setup(app *fiber.App, db *sql.DB) {
	devLog, _ := zap.NewDevelopment()
	logger := devLog.Sugar()

	// create handler here
	logger.Debug("handlers starting")
	handler := restapi.New(database.NewUsers(db), database.NewSkills(db), logger)
	logger.Debug("handlers created")

	// base group for v1 api
	group := app.Group(URLPrefix)

	// set all routes
	logger.Debug("setting up routes")
	routes(group, handler, logger)
	logger.Debug("server started")
}
