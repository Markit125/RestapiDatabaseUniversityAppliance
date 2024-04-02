package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestStudent(t *testing.T) *Student {
	return &Student{
		FirstName:    "Mark",
		MiddleName:   "Ivanovich",
		LastName:     "Tarabukin",
		BirthDate:    "11-05-2003",
		Achievements: 0,
		Passport:     "5451535825",
	}
}
