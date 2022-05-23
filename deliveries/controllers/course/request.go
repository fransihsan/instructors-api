package course

import (
	C "instructors-api/entities/course"

	"gorm.io/gorm"
)

type RequestCreate struct {
	CourseTitle   string `json:"course_title" form:"course_title"`
	CourseBenefit string `json:"course_benefit" form:"course_benefit"`
	CoursePrice   uint   `json:"course_price" form:"course_price"`
	CourseDesc    string `json:"course_desc" form:"course_desc"`
}

func (Req RequestCreate) ToEntityService(instructorID uint) C.Courses {
	return C.Courses{
		CourseTitle:   Req.CourseTitle,
		CourseBenefit: Req.CourseBenefit,
		CoursePrice:   Req.CoursePrice,
		CourseDesc:    Req.CourseDesc,
		IstructorID:   instructorID,
	}
}

type RequestUpdate struct {
	CourseTitle   string `json:"course_title" form:"course_title"`
	CourseBenefit string `json:"course_benefit" form:"course_benefit"`
	CoursePrice   uint   `json:"course_price" form:"course_price"`
	CourseDesc    string `json:"course_desc" form:"course_desc"`
}

func (Req RequestUpdate) ToEntityService(InstructorID, ID uint) C.Courses {
	return C.Courses{
		Model:         gorm.Model{ID: ID},
		CourseTitle:   Req.CourseTitle,
		CourseBenefit: Req.CourseBenefit,
		CoursePrice:   Req.CoursePrice,
		CourseDesc:    Req.CourseDesc,
		IstructorID: InstructorID,
	}
}
