package course

import C "instructors-api/entities/course"

type Course interface {
	Create(newInstructor C.Courses) (C.Courses, error)
	Get() ([]FormatGetCourse, error)
	GetDetail(ID uint) ([]FormatGetDetailCourse, error)
	Update(userUpdate C.Courses) (C.Courses, error)
	Delete(InstructorID, ID uint) error
}