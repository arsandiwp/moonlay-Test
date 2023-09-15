package routes

import (
	"moonlay-test/internal/app/handlers"
	"moonlay-test/internal/app/repositories"
	"moonlay-test/internal/database"
	"moonlay-test/web/middleware"

	"github.com/labstack/echo/v4"
)

func SubListRoutes(e *echo.Group) {
	sublistRepository := repositories.RepositorySubList(database.DB)
	todolistRepository := repositories.RepositoryTodoList(database.DB)
	h := handlers.HandlerSubList(sublistRepository, todolistRepository)

	e.GET("/sublists", h.GetAllSubLists)
	e.GET("/sublist/:id", h.GetSubList)
	e.POST("/sublist", middleware.UploadFile(h.CreateSubList))
	e.PATCH("/sublist/:id", middleware.UploadFile(h.UpdateSubList))
}
