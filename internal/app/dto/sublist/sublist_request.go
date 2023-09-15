package sublistsdto

import "moonlay-test/internal/app/models"

type CreateSubListRequest struct {
	TodoID      int             `json:"todo_id"`
	Todo        models.TodoList `json:"todo"`
	Title       string          `json:"title" form:"title" validate:"required,max=100"`
	Description string          `json:"description" form:"description" validate:"required,max=1000"`
	Files       string          `json:"files" form:"files"`
}

type UpdateSubListRequest struct {
	TodoID      int             `json:"todo_id"`
	Todo        models.TodoList `json:"todo"`
	Title       string          `json:"title" form:"title"`
	Description string          `json:"description" form:"description"`
	Files       string          `json:"files" form:"files"`
}
