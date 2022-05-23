package auth

import (
	"errors"
	I "instructors-api/entities/instructor"
	"instructors-api/repositories/hash"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) Login(email, password string) (I.Instructors, error) {
	var instructor I.Instructors

	repo.db.Model(&instructor).Where("email = ?", email).First(&instructor)

	isMatched := hash.CheckPasswordHash(instructor.Password, password)
	if !isMatched {
		return I.Instructors{}, errors.New("email dan password belum terdaftar")
	}

	return instructor, nil
}
