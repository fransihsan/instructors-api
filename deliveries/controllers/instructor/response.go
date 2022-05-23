package instructor

import I "instructors-api/entities/instructor"

type ResponseCreateInstructor struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	InstructorName       string `json:"instructor_name"`
	InstructorProfession string `json:"instructor_profession"`
	AboutInstructor      string `json:"about_instructor"`
}

func ToResponseCreateInstructor(Instructor I.Instructors) ResponseCreateInstructor {
	return ResponseCreateInstructor{
		ID:                   Instructor.ID,
		Email:                Instructor.Email,
		InstructorName:       Instructor.InstructorName,
		InstructorProfession: Instructor.InstructorProfession,
		AboutInstructor:      Instructor.AboutInstructor,
	}
}

type ResponseUpdateInstructor struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	InstructorName       string `json:"instructor_name"`
	InstructorProfession string `json:"instructor_profession"`
	AboutInstructor      string `json:"about_instructor"`
}

func ToResponseUpdateInstructor(Instructor I.Instructors) ResponseUpdateInstructor {
	return ResponseUpdateInstructor{
		ID:      Instructor.ID,
		Email:   Instructor.Email,
		InstructorName:       Instructor.InstructorName,
		InstructorProfession: Instructor.InstructorProfession,
		AboutInstructor:      Instructor.AboutInstructor,
	}
}
