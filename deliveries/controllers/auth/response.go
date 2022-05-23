package auth

import I "instructors-api/entities/instructor"

type ResponseLogin struct {
	ID             uint   `json:"id"`
	InstructorName string `json:"instructor_name"`
	Token          string `json:"token"`
}

func ToResponseLogin(instructor I.Instructors, token string) ResponseLogin {
	return ResponseLogin{
		ID:             instructor.ID,
		InstructorName: instructor.InstructorName,
		Token:          token,
	}
}
