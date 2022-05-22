package course

type FormatGetCourse struct {
	ID             uint   `json:"id"`
	CourseTitle    string `json:"course_title"`
	InstructorName string `json:"instructor_name"`
	CoursePrice    uint   `json:"course_price"`
}

type FormatGetDetailCourse struct {
	ID             uint   `json:"id"`
	CourseTitle    string `json:"course_title"`
	CourseBenefit  string `json:"course_benefit"`
	CoursePrice    uint   `json:"course_price"`
	CourseDesc     string `json:"course_desc"`
	InstructorName string `json:"instructor_name"`
}
