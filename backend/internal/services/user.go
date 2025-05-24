package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Renn-Amm/Project-Edge/internal/form"
	"github.com/Renn-Amm/Project-Edge/internal/models"
	"github.com/Renn-Amm/Project-Edge/internal/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp (user *models.User) error
	Login (request *form.LoginRequest) (string,error)
	GetUserByEmail (email string) (*models.User,error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) SignUp (user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }
	user.Password = string(hashedPassword)
	if err := s.repo.CreateUser(user); err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}

func (s *UserServiceImpl) Login (request *form.LoginRequest) (string,error) {

	user, err := s.repo.FindByEmail(request.Email)

	if err != nil {
        return "", fmt.Errorf("failed to find user: %w", err)
    }

	if user == nil {
        return "", errors.New("invalid email or password")
    }

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	JWTSecret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(JWTSecret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}
func (s *UserServiceImpl) GetUserByEmail (email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}