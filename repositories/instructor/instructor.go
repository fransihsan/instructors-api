package instructor

import (
	"errors"
	I "instructors-api/entities/instructor"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type InstructorRepository struct {
	db *gorm.DB
}

func NewInstructorRepository(db *gorm.DB) *InstructorRepository {
	return &InstructorRepository{
		db: db,
	}
}

func (repo *InstructorRepository) Create(newInstructor I.Instructors) (I.Instructors, error) {
	if err := repo.db.Create(&newInstructor).Error; err != nil {
		log.Warn(err)
		return I.Instructors{}, errors.New("gagal membuat instruktur baru")
	}
	return newInstructor, nil
}

func (repo *InstructorRepository) Update(instructorUpdate I.Instructors) (I.Instructors, error) {
	if rowsAffected := repo.db.Model(&instructorUpdate).Updates(instructorUpdate).RowsAffected; rowsAffected == 0 {
		return I.Instructors{}, errors.New("tidak ada perubahan pada data instructor")
	}
	repo.db.First(&instructorUpdate)
	return instructorUpdate, nil
}