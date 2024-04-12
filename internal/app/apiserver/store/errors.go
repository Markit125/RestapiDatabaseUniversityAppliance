package store

import "errors"

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrRecordAlredyExists = errors.New("record already exists")
)
