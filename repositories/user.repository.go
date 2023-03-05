package repositories

import (
	"github.com/stefanusong/votify-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) *models.User
	InsertUser(user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUserByUsername(username string) *models.User {
	var user *models.User
	res := ur.db.Where("username = ?", username).First(&user)

	if res.RowsAffected == 0 {
		user = nil
	}

	return user
}

func (ur *userRepository) InsertUser(user models.User) error {
	return ur.db.Create(&user).Error
}
