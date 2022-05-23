package course

import C "instructors-api/entities/course"

type ResponseCreate struct {
	InstructorID  uint   `json:"instructor_id"`
	CourseTitle   string `json:"course_title"`
	CourseBenefit string `json:"course_benefit"`
	CoursePrice   uint   `json:"course_price"`
	CourseDesc    string `json:"course_desc"`
}

func ToResponseCreate(course C.Courses) ResponseCreate {
	return ResponseCreate{
		InstructorID:  course.IstructorID,
		CourseTitle:   course.CourseTitle,
		CourseBenefit: course.CourseBenefit,
		CoursePrice:   course.CoursePrice,
		CourseDesc:    course.CourseDesc,
	}
}

type ResponseUpdate struct {
	ID            uint   `json:"id"`
	InstructorID  uint   `json:"instructor_id"`
	CourseTitle   string `json:"course_title"`
	CourseBenefit string `json:"course_benefit"`
	CoursePrice   uint   `json:"course_price"`
	CourseDesc    string `json:"course_desc"`
}

func ToResponseUpdate(course C.Courses) ResponseUpdate {
	return ResponseUpdate{
		ID:            course.ID,
		InstructorID:  course.IstructorID,
		CourseTitle:   course.CourseTitle,
		CourseBenefit: course.CourseBenefit,
		CoursePrice:   course.CoursePrice,
		CourseDesc:    course.CourseDesc,
	}
}
