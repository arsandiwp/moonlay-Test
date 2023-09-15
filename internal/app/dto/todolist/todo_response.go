package todolistsdto

type TodoListResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
}
