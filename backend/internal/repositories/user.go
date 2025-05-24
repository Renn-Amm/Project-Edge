package repositories

import (
	"github.com/Renn-Amm/Project-Edge/internal/database"
	"github.com/Renn-Amm/Project-Edge/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    FindByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepository {
	return &userRepo{db: database.DB} 
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}