package instructor

import (
	"instructors-api/deliveries/helpers/hash"
	I "instructors-api/entities/instructor"
	"strings"

	"gorm.io/gorm"
)

type RequestCreateInstructor struct {
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	InstructorName       string `json:"instructor_name" form:"instructor_name"`
	InstructorProfession string `json:"instructor_profession" form:"instructor_profession"`
	AboutInstructor      string `json:"about_instructor" form:"about_instructor"`
}

func (Req RequestCreateInstructor) ToEntityInstructor() I.Instructors {
	hashedPassword, _ := hash.HashPassword(Req.Password)

	return I.Instructors{
		Email:                Req.Email,
		Password:             hashedPassword,
		InstructorName:       Req.InstructorName,
		InstructorProfession: Req.InstructorProfession,
		AboutInstructor:      Req.AboutInstructor,
	}
}

type RequestUpdateInstructor struct {
	ID                   uint
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	InstructorName       string `json:"instructor_name" form:"instructor_name"`
	InstructorProfession string `json:"instructor_profession" form:"instructor_profession"`
	AboutInstructor      string `json:"about_instructor" form:"about_instructor"`
}

func (Req RequestUpdateInstructor) ToEntityUser(InstructorID uint) I.Instructors {
	if strings.TrimSpace(Req.Password) != "" {
		Req.Password, _ = hash.HashPassword(Req.Password)
	}

	return I.Instructors{
		Model:                gorm.Model{ID: InstructorID},
		Email:                Req.Email,
		Password:             Req.Password,
		InstructorName:       Req.InstructorName,
		InstructorProfession: Req.InstructorProfession,
		AboutInstructor:      Req.AboutInstructor,
	}
}
