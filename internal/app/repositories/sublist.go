package repositories

import (
	"moonlay-test/internal/app/models"

	"gorm.io/gorm"
)

type SubListRepository interface {
	GetAllSubLists(page, pageSize int, search string) ([]models.SubList, int, error)
	GetSubList(ID int) (models.SubList, error)
	CreateSubList(SubList models.SubList) (models.SubList, error)
	UpdateSubList(SubList models.SubList) (models.SubList, error)
	DeleteSubList(SubList models.SubList) (models.SubList, error)
}

func RepositorySubList(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllSubLists(page, pageSize int, search string) ([]models.SubList, int, error) {
	var sublists []models.SubList
	query := r.db.Model(&models.SubList{})

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalSubLists int64
	err := query.Count(&totalSubLists).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&totalSubLists).Error
	if err != nil {
		return nil, 0, err
	}

	return sublists, int(totalSubLists), nil
}

func (r *repository) GetSubList(ID int) (models.SubList, error) {
	var sublist models.SubList
	err := r.db.Preload("Todo").First(&sublist, ID).Error

	return sublist, err
}

func (r *repository) CreateSubList(SubList models.SubList) (models.SubList, error) {
	err := r.db.Create(&SubList).Error

	return SubList, err
}

func (r *repository) UpdateSubList(SubList models.SubList) (models.SubList, error) {
	err := r.db.Save(&SubList).Error

	return SubList, err
}

func (r *repository) DeleteSubList(SubList models.SubList) (models.SubList, error) {
	err := r.db.Delete(&SubList).Scan(&SubList).Error

	return SubList, err
}
