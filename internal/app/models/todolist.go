package models

type TodoList struct {
	ID          int               `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string            `json:"title" gorm:"type: varchar(100)"`
	Description string            `json:"description" gorm:"type: text"`
	Files       string            `json:"files" gorm:"type: varchar(255)"`
	SubList     []SubListResponse `json:"sublist" gorm:"foreignKey:TodoID"`
}

type TodoListResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
}

func (TodoListResponse) TableName() string {
	return "todo_lists"
}
