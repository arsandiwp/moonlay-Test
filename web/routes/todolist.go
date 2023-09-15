package routes

import (
	"moonlay-test/internal/app/handlers"
	"moonlay-test/internal/app/repositories"
	"moonlay-test/internal/database"
	"moonlay-test/web/middleware"

	"github.com/labstack/echo/v4"
)

func TodoListRoute(e *echo.Group) {
	todolistRepository := repositories.RepositoryTodoList(database.DB)
	h := handlers.HandlerTodoList(todolistRepository)

	e.GET("/alltodos", h.FindTodoList)
	e.GET("/todolists", h.GetAllTodoLists)
	e.GET("/todolist/:id", h.GetTodoList)
	e.POST("/todolist", middleware.UploadFile(h.CreateTodoList))
	e.PATCH("/todolist/:id", middleware.UploadFile(h.UpdateTodoList))
}
