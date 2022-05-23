package auth

import I "instructors-api/entities/instructor"

type Auth interface {
	Login(email, password string) (I.Instructors, error)
}
