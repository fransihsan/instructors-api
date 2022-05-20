package course

import (
	"gorm.io/gorm"
)

type Courses struct {
	gorm.Model
	CourseTitle   string `gorm:"type:varchar(255);not null;unique"`
	CourseBenefit string `gorm:"type:text;not null"`
	CoursePrice   uint   `gorm:"not null"`
	CourseDesc    string `gorm:"type:text;not null"`
	IstructorID   uint
}
