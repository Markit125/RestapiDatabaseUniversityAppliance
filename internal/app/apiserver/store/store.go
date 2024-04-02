package store

// Store ...
type Store interface {
	User() UserRepository
	Student() StudentRepository
}
