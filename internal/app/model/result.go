package model

import validation "github.com/go-ozzo/ozzo-validation"

// Result ...
type Result struct {
	ID            int `json:"id"`
	ResultScore   int `json:"score"`
	StudentID     int `json:"student_id"`
	ExamSubjectID int `json:"exam_subject_id"`
}

// Validate ...
func (s *Result) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.ResultScore, validation.Required),
		validation.Field(&s.StudentID, validation.Required),
		validation.Field(&s.ExamSubjectID, validation.Required),
	)
}
