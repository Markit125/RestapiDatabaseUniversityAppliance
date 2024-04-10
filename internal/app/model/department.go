package model

import validation "github.com/go-ozzo/ozzo-validation"

// Department ...
type Department struct {
	ID             int    `json:"id"`
	DepartmentName string `json:"department_name"`
}

// Validate ...
func (s *Department) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.DepartmentName, validation.Required),
	)
}
