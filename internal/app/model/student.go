package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Student ...
type Student struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date"`
	Achievements int    `json:"achievements"`
	Passport     string `json:"passport"`
}

// Validate ...
func (s *Student) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.MiddleName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.BirthDate, validation.Required),
		// validation.Field(&s.Achievements, validation.Required),
		validation.Field(&s.Passport, validation.Required, validation.Length(10, 10)),
	)
}
