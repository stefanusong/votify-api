package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(Category models.Category) (uuid.UUID, error)
	GetCategories() ([]models.Category, error)
	GetCategoryByID(ID string) (*models.Category, error)
	UpdateCategoryByID(ID string, Category models.Category) (uuid.UUID, error)
	DeleteCategoryByID(ID string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) InsertCategory(Category models.Category) (uuid.UUID, error) {
	err := cr.db.Create(&Category).Error
	return Category.ID, err
}

func (cr *categoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := cr.db.Find(&categories).Error
	return categories, err
}

func (cr *categoryRepository) GetCategoryByID(ID string) (*models.Category, error) {
	var category *models.Category
	res := cr.db.First(&category, "ID = ?", ID)

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return category, res.Error
}

func (cr *categoryRepository) UpdateCategoryByID(ID string, Category models.Category) (uuid.UUID, error) {
	res := cr.db.Model(&models.Category{}).Where("ID = ?", ID).Updates(Category)

	if res.RowsAffected == 0 {
		return uuid.UUID{}, nil
	}

	categoryID, _ := uuid.FromString(ID)
	return categoryID, res.Error
}

func (cr *categoryRepository) DeleteCategoryByID(ID string) error {
	err := cr.db.Delete(&models.Category{}, "ID = ?", ID).Error
	return err
}
