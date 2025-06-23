package auth

import (
	"errors"
	"studentApi/models"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Signup(email, password string) error {
	//* Steps :
	// check if user already exist
	// If exist - throw error
	// else hash user Password and create a user
	// Will generate the token at login time

	var existingUser models.User
	err := s.DB.Where("email = ?", email).First(&existingUser).Error
	if err == nil { // means user already exist
		return errors.New("user already exists")
	}

	hashedPas, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: hashedPas,
	}

	// create new user
	return s.DB.Create(&user).Error
}

func (s *AuthService) Login(email, password string) (string, error) {
	//* Steps :
	// Check if user exists or not using email
	// If not then throw error
	// If exists match password if matches then generate token
	var user models.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if !CheckPasswordHash(user.Password, password) {
		return "", errors.New("invalid credentials your password doesn't matched")
	}

	// Generate token with UserId and email
	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
