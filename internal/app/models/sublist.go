package models

type SubList struct {
	ID          int              `json:"id" gorm:"primaryKey;autoIncrement"`
	TodoID      int              `json:"todo_id"`
	Todo        TodoListResponse `json:"todo"`
	Title       string           `json:"title" gorm:"type: varchar(100)"`
	Description string           `json:"description" gorm:"type: text"`
	Files       string           `json:"files" gorm:"type: varchar(255)"`
}

type SubListResponse struct {
	ID          int    `json:"id"`
	TodoID      int    `json:"tod_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
}

func (SubListResponse) TableName() string {
	return "sub_lists"
}
