package AuthService

import (
	"Sistem-Absensi-Backend-Go/Database"
	"Sistem-Absensi-Backend-Go/Dto/AuthDto"
	"Sistem-Absensi-Backend-Go/Models/Entity"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func SignUp(name, email, password, role string) (*AuthDto.UserResponse, error) {
	db := Database.DB

	var existing Entity.User
	if err := db.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, errors.New("user already exists")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &Entity.User{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      role,
		IsActive:  "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	resp := &AuthDto.UserResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return resp, nil
}

func SignIn(email, password string) (*AuthDto.UserResponse, error) {
	db := Database.DB

	var user Entity.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	resp := &AuthDto.UserResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return resp, nil
}
