package model

import validation "github.com/go-ozzo/ozzo-validation"

// Appliance ...
type Appliance struct {
	ID             int `json:"id"`
	StudentID      int `json:"student_id"`
	DepartmentID   int `json:"department_id"`
	ExamSubject1ID int `json:"exam_subject_1_id"`
	ExamSubject2ID int `json:"exam_subject_2_id"`
	ExamSubject3ID int `json:"exam_subject_3_id"`
}

// Validate ...
func (s *Appliance) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.StudentID, validation.Required),
		validation.Field(&s.DepartmentID, validation.Required),
		validation.Field(&s.ExamSubject1ID, validation.Required),
		validation.Field(&s.ExamSubject2ID, validation.Required),
		validation.Field(&s.ExamSubject3ID, validation.Required),
	)
}
