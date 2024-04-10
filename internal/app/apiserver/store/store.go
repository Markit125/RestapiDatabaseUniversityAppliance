package store

// Store ...
type Store interface {
	User() UserRepository
	Student() StudentRepository
	// ExamSubject() ExamSubjectRepository
	// ExamResult() ExamResultRepository
	// Department() DepartmentRepository
	// Appliance() ApplianceRepository
}
