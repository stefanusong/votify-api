package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/services"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	GetAllCategories(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	UpdateCategoryByID(c *gin.Context)
	DeleteCategoryByID(c *gin.Context)
}

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

func (handler *categoryHandler) CreateCategory(c *gin.Context) {
	// Bind Category
	var newCategory request.CreateCategory
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		resp := response.New(false, "Failed to create new category", nil, err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Create category
	data, err := handler.categoryService.CreateCategory(newCategory)
	if err != nil {
		resp := response.New(false, "Failed to create new category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Category created", data, nil)
	c.JSON(http.StatusCreated, resp)
}

func (handler *categoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := handler.categoryService.GetAllCategories()

	if err != nil {
		resp := response.New(false, "Failed to get categories", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Success", categories, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *categoryHandler) GetCategoryByID(c *gin.Context) {
	categoryId := c.Param("id")
	category, err := handler.categoryService.GetCategoryByID(categoryId)

	if category == nil {
		resp := response.New(false, "Failed to get category", nil, "category not found")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	if err != nil {
		resp := response.New(false, "Failed to get category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Success", category, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *categoryHandler) UpdateCategoryByID(c *gin.Context) {
	// Bind Category
	categoryId := c.Param("id")
	var updatedCategory request.UpdateCategory
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		resp := response.New(false, "Failed to update category", nil, err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	category, err := handler.categoryService.UpdateCategoryByID(categoryId, updatedCategory)

	if err != nil {
		resp := response.New(false, "Failed to update category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if category == nil {
		resp := response.New(false, "Failed to update category", nil, "category not found")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := response.New(true, "Success", categoryId, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *categoryHandler) DeleteCategoryByID(c *gin.Context) {
	categoryId := c.Param("id")
	err := handler.categoryService.DeleteCategoryByID(categoryId)

	if err != nil {
		resp := response.New(false, "Failed to delete category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "Success", nil, nil)
	c.JSON(http.StatusOK, resp)
}
