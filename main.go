package main

import (
	"moonlay-test/internal/database"
	"moonlay-test/internal/database/migration"
	"moonlay-test/web/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.DatabaseInit()
	migration.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	e.Logger.Fatal(e.Start(":5000"))
}
