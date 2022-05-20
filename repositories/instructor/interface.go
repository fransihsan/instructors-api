package instructor

import I "instructors-api/entities/instructor"

type Instructor interface {
	Create(newInstructor I.Instructors) (I.Instructors, error)
	Update(userUpdate I.Instructors) (I.Instructors, error)
}