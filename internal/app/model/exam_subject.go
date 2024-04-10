package model

import validation "github.com/go-ozzo/ozzo-validation"

// ExamSubject ...
type ExamSubject struct {
	ID          int    `json:"id"`
	SubjectName string `json:"subject_name"`
}

// Validate ...
func (s *ExamSubject) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.SubjectName, validation.Required),
	)
}
