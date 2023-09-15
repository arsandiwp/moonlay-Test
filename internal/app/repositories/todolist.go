package repositories

import (
	"moonlay-test/internal/app/models"

	"gorm.io/gorm"
)

type TodoListRepository interface {
	FindTodoList() ([]models.TodoList, error)
	GetAllTodoLists(page, pageSize int, search string) ([]models.TodoList, int, error)
	GetTodoList(ID int) (models.TodoList, error)
	CreateTodoList(TodoList models.TodoList) (models.TodoList, error)
	UpdateTodoList(TodoList models.TodoList) (models.TodoList, error)
	DeleteTodoList(TodoList models.TodoList) (models.TodoList, error)
}

func RepositoryTodoList(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindTodoList() ([]models.TodoList, error) {
	var todolists []models.TodoList
	err := r.db.Find(&todolists).Error

	return todolists, err
}

func (r *repository) GetAllTodoLists(page, pageSize int, search string) ([]models.TodoList, int, error) {
	var todolists []models.TodoList
	query := r.db.Model(&models.TodoList{})

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalTodoList int64
	err := query.Count(&totalTodoList).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&todolists).Error
	if err != nil {
		return nil, 0, err
	}

	return todolists, int(totalTodoList), nil
}

func (r *repository) GetTodoList(ID int) (models.TodoList, error) {
	var todolist models.TodoList
	err := r.db.First(&todolist, ID).Error

	return todolist, err
}

func (r *repository) CreateTodoList(TodoList models.TodoList) (models.TodoList, error) {
	err := r.db.Create(&TodoList).Error

	return TodoList, err
}

func (r *repository) UpdateTodoList(TodoList models.TodoList) (models.TodoList, error) {
	err := r.db.Save(&TodoList).Error

	return TodoList, err
}

func (r *repository) DeleteTodoList(TodoList models.TodoList) (models.TodoList, error) {
	err := r.db.Delete(&TodoList).Scan(&TodoList).Error

	return TodoList, err
}
