package instructor

import (
	"gorm.io/gorm"
)

type Instructors struct {
	gorm.Model
	Email                string `gorm:"type:varchar(255);not null;unique"`
	Password             string `gorm:"type:varchar(255);not null"`
	InstructorName       string `gorm:"type:varchar(255);not null"`
	InstructorProfession string `gorm:"type:varchar(255);not null"`
	AboutInstructor      string `gorm:"type:varchar(255)"`
	Courses  []course.Courses   `gorm:"foreignKey:IstructorID"`
}
