package services

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/models"
	"github.com/stefanusong/votify-api/repositories"
)

type CategoryService interface {
	CreateCategory(categoryReq request.CreateCategory) (any, error)
	GetAllCategories() (any, error)
	GetCategoryByID(ID string) (any, error)
	UpdateCategoryByID(ID string, categoryReq request.UpdateCategory) (any, error)
	DeleteCategoryByID(ID string) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (svc *categoryService) CreateCategory(categoryReq request.CreateCategory) (any, error) {
	category := models.NewCategory(categoryReq.Name)

	categoryId, err := svc.categoryRepo.InsertCategory(category)
	if err != nil {
		return nil, err
	}

	return map[string]string{"category_id": categoryId.String()}, nil
}

func (svc *categoryService) GetAllCategories() (any, error) {
	categories, err := svc.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	return map[string][]models.Category{"categories": categories}, nil
}

func (svc *categoryService) GetCategoryByID(ID string) (any, error) {
	category, err := svc.categoryRepo.GetCategoryByID(ID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, nil
	}

	return map[string]*models.Category{"category": category}, nil
}

func (svc *categoryService) UpdateCategoryByID(ID string, categoryReq request.UpdateCategory) (any, error) {
	category := models.NewCategory(categoryReq.Name)

	categoryId, err := svc.categoryRepo.UpdateCategoryByID(ID, category)
	if err != nil {
		return nil, err
	}

	if (categoryId == uuid.UUID{}) {
		return nil, nil
	}

	return map[string]string{"category_id": categoryId.String()}, nil
}

func (svc *categoryService) DeleteCategoryByID(ID string) error {
	err := svc.categoryRepo.DeleteCategoryByID(ID)
	return err
}
