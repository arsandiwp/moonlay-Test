package sublistsdto

import "moonlay-test/internal/app/models"

type SubListResponse struct {
	ID          int                     `json:"id"`
	TodoID      int                     `json:"todo_id"`
	Todo        models.TodoListResponse `json:"todo"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Files       string                  `json:"files"`
}
