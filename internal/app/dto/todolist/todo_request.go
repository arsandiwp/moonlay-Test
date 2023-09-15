package todolistsdto

type CreateTodoListRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=100"`
	Description string `json:"description" form:"description" validate:"required,max=1000"`
	Files       string `json:"files" form:"files"`
}

type UpdateTodoListRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Files       string `json:"files" form:"files"`
}
