package request

type CreateCategory struct {
	Name string `json:"name" binding:"required,min=3"`
}

type UpdateCategory struct {
	Name string `json:"name" binding:"required,min=3"`
}
